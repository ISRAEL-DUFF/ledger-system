package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/types/datastructure"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type ITransactionService interface {
	CreateLedgerTransaction(input types.TransactionInput) (types.TransactionResponse, error)
	PostQueuedWalletTransaction(input types.PostTransactionInput) error
	CreateQueuedLedgerTransaction(input types.TransactionInput) (types.TransactionResponse, error)
}

type CreateTransactionOption func(*types.TransactionInput)

type TransactionService struct {
	transactionRepo  repositories.ILedgerTransactionRepository
	journalRepo      repositories.IJournalEntryRepository
	accountBlockRepo repositories.IAccountBlockRepository
	accountRepo      repositories.ILedgerAccount
	blockMetumRepo   repositories.IBlockMetumRepository
	accountService   IAccountService
	txQ              ITransactionQService
}

func NewTransactionService(
	transactionRepo repositories.ILedgerTransactionRepository,
	journalRep repositories.IJournalEntryRepository,
	accountBlockRepo repositories.IAccountBlockRepository,
	accountRepo repositories.ILedgerAccount,
	blockMetumRepo repositories.IBlockMetumRepository,
	txQ ITransactionQService,
	accountService IAccountService) *TransactionService {
	return &TransactionService{
		transactionRepo:  transactionRepo,
		journalRepo:      journalRep,
		accountBlockRepo: accountBlockRepo,
		accountRepo:      accountRepo,
		blockMetumRepo:   blockMetumRepo,
		accountService:   accountService,
		txQ:              txQ,
	}
}

func (txService *TransactionService) CreateLedgerTransaction(input types.TransactionInput) (types.TransactionResponse, error) {
	sumOfCredits := 0
	sumOfDebits := 0

	for _, entry := range input.Entries {
		if entry.Type == types.CREDIT {
			sumOfCredits += entry.Amount
		} else {
			sumOfDebits += entry.Amount
		}
	}

	if sumOfCredits != sumOfDebits {
		return types.TransactionResponse{}, errors.New("invalid transaction amounts")
	}

	// dbQueryTx := dbQuery.Begin()
	dbQueryTx := txService.accountBlockRepo.BeginTransaction()

	txRepo := txService.transactionRepo.WithTransaction(dbQueryTx)

	transaction, err := txRepo.Create(types.CreateLedgerTransaction{
		Status: types.PENDING,
	})

	if err != nil {
		return types.TransactionResponse{}, err
	}

	treatedEntries := make([]types.TransactionEntry, len(input.Entries))

	for index, entry := range input.Entries {
		account := txService.accountService.GetAccount(entry.AccountNumber)

		if entry.Type == types.DEBIT {
			if account.Label == "A1" && account.Balance < entry.Amount {
				// panic("Insufficient fund on account " + entry.AccountNumber)
				return types.TransactionResponse{}, errors.New("Insufficient fund on account " + entry.AccountNumber)
			}
		}

		// TODO: post transaction here
		txService.postTransaction(entry, transaction.ID, *account, dbQueryTx)

		treatedEntries[index] = types.TransactionEntry{
			Amount:        entry.Amount,
			AccountNumber: entry.AccountNumber,
			Type:          entry.Type,
		}
	}

	updateErr := txRepo.UpdateStatus(transaction.ID, types.TRANSACTION_APPROVED)

	if updateErr != nil {
		dbQueryTx.Rollback()
	}

	dbQueryTx.Commit()

	return types.TransactionResponse{
		Status:        types.TRANSACTION_APPROVED,
		TransactionId: transaction.ID,
		Entries:       treatedEntries[:],
	}, nil
}

func (txService *TransactionService) createSingleLedgerTransaction(input types.TransactionInput, dbQueryTx types.IDBTransaction) (types.TransactionResponse, error) {
	sumOfCredits := 0
	sumOfDebits := 0

	for _, entry := range input.Entries {
		if entry.Type == types.CREDIT {
			sumOfCredits += entry.Amount
		} else {
			sumOfDebits += entry.Amount
		}
	}

	if sumOfCredits != sumOfDebits {
		return types.TransactionResponse{}, errors.New("invalid transaction amounts")
	}

	txRepo := txService.transactionRepo.WithTransaction(dbQueryTx)

	transaction, err := txRepo.Create(types.CreateLedgerTransaction{
		Status: types.PENDING,
	})

	if err != nil {
		return types.TransactionResponse{}, err
	}

	treatedEntries := make([]types.TransactionEntry, len(input.Entries))

	for index, entry := range input.Entries {
		account := txService.accountService.GetAccount(entry.AccountNumber)

		if entry.Type == types.DEBIT {
			if account.Label == "A1" && account.Balance < entry.Amount {
				return types.TransactionResponse{}, errors.New("Insufficient fund on account " + entry.AccountNumber)
			}
		}

		txService.postTransaction(entry, transaction.ID, *account, dbQueryTx)

		treatedEntries[index] = types.TransactionEntry{
			Amount:        entry.Amount,
			AccountNumber: entry.AccountNumber,
			Type:          entry.Type,
		}
	}

	updateErr := txRepo.UpdateStatus(transaction.ID, types.TRANSACTION_APPROVED)

	if updateErr != nil {
		return types.TransactionResponse{}, updateErr
	}

	return types.TransactionResponse{
		Status:        types.TRANSACTION_APPROVED,
		TransactionId: transaction.ID,
		Entries:       treatedEntries[:],
	}, nil
}

func (txService *TransactionService) CreateQueuedLedgerTransaction(input types.TransactionInput) (types.TransactionResponse, error) {
	r, err := txService.txQ.Schedule(input, func(resInput types.TransactionInput) (types.TransactionResponse, error) {
		resp, err := txService.CreateLedgerTransaction(resInput)

		return resp, err
	})

	r.ReleaseLock()

	return r.TxResponse, err
}

func (txService *TransactionService) PostQueuedWalletTransaction(input types.PostTransactionInput) error {

	txEntries, err := txService.accountService.ExtractTransactionEntries(input, [][]types.TransactionInputEntry{})

	if err != nil {
		return err
	}

	dbQueryTx := txService.accountBlockRepo.BeginTransaction()

	fmt.Println(txEntries)

	var schedulerResp []SchedulerResponse = []SchedulerResponse{}

	for _, entry := range txEntries {

		r, err := txService.txQ.Schedule(types.TransactionInput{
			Entries: entry,
		}, func(resInput types.TransactionInput) (types.TransactionResponse, error) {
			resp, er := txService.createSingleLedgerTransaction(resInput, dbQueryTx)

			return resp, er
		})

		if err != nil {
			dbQueryTx.Rollback()
			return err
		}

		schedulerResp = append(schedulerResp, r)
	}

	// txService.mu.Lock()
	dbQueryTx.Commit()
	// txService.mu.Unlock()

	for _, r := range schedulerResp {
		r.ReleaseLock()
	}

	return nil
}

func (txService *TransactionService) FindTransactionBlockContaining(transactionDate int64, accountId string) (types.BlockSearchStatus, *types.AccountBlockType) {
	leftBlock, _ := txService.accountBlockRepo.FindByDate(transactionDate, accountId, "left")
	rightBlock, _ := txService.accountBlockRepo.FindByDate(transactionDate, accountId, "right")

	// if leftErr != nil || rightErr != nil {
	// 	panic("error fetching blocks")
	// }

	if leftBlock != nil && rightBlock != nil && rightBlock.ID == leftBlock.ID {
		fmt.Printf("Found Block %s\n", rightBlock.ID)
		return types.WITHIN_BLOCK, &types.AccountBlockType{
			ID:                rightBlock.ID,
			Status:            types.AccountBlockStatus(rightBlock.Status),
			AccountID:         rightBlock.AccountID,
			BlockSize:         rightBlock.BlockSize,
			TransactionsCount: rightBlock.TransactionsCount,
			CreatedAt:         rightBlock.CreatedAt.String(),
		}
	}

	if leftBlock != nil && rightBlock != nil && rightBlock.ID != leftBlock.ID {
		fmt.Println("No match Found returning RIGHT Block")
		return types.NOT_IN_ANY_BLOCK, &types.AccountBlockType{
			ID:                rightBlock.ID,
			Status:            types.AccountBlockStatus(rightBlock.Status),
			AccountID:         rightBlock.AccountID,
			BlockSize:         rightBlock.BlockSize,
			TransactionsCount: rightBlock.TransactionsCount,
			CreatedAt:         rightBlock.CreatedAt.String(),
		}
	}

	if rightBlock != nil {
		return types.RIGHT_ONLY, &types.AccountBlockType{
			ID:                rightBlock.ID,
			Status:            types.AccountBlockStatus(rightBlock.Status),
			AccountID:         rightBlock.AccountID,
			BlockSize:         rightBlock.BlockSize,
			TransactionsCount: rightBlock.TransactionsCount,
			CreatedAt:         rightBlock.CreatedAt.String(),
		}
	}

	if leftBlock != nil {
		return types.LEFT_ONLY, &types.AccountBlockType{
			ID:                leftBlock.ID,
			Status:            types.AccountBlockStatus(leftBlock.Status),
			AccountID:         leftBlock.AccountID,
			BlockSize:         leftBlock.BlockSize,
			TransactionsCount: leftBlock.TransactionsCount,
			CreatedAt:         leftBlock.CreatedAt.String(),
		}
	}

	panic("error processing blocks")
}

func (txService *TransactionService) FindTransactionBlocks(accountID string, startDate, endDate int64) ([]types.AccountBlockType, error) {
	listOfMetaInfo := []*model.BlockMetum{}
	accountBlockIds := []string{}
	accountBlockIDs := datastructure.NewSet[string]()
	var startBlock *types.AccountBlockType
	var endBlock *types.AccountBlockType
	accountBlocks := []types.AccountBlockType{}
	account, err := txService.accountRepo.FindById(accountID)

	if err != nil {
		return []types.AccountBlockType{}, err
	}

	if startDate == endDate {
		startBlkStatus, startBlk := txService.FindTransactionBlockContaining(startDate, accountID)

		if startBlk != nil {
			return []types.AccountBlockType{
				*startBlk,
			}, nil
		}

		if startBlkStatus == types.LEFT_ONLY {
			fmt.Println("Finding Last Active block")
			startBllk, err := txService.accountBlockRepo.GetCurrentOpenBlock(accountID)

			if err != nil {
				return []types.AccountBlockType{}, err
			}

			if startBllk == nil {
				return []types.AccountBlockType{}, errors.New("invalid block")
			}

			return []types.AccountBlockType{
				{
					ID:                startBllk.ID,
					Status:            types.AccountBlockStatus(startBllk.Status),
					AccountID:         startBllk.AccountID,
					BlockSize:         startBllk.BlockSize,
					TransactionsCount: startBllk.TransactionsCount,
					CreatedAt:         startBllk.CreatedAt.String(),
				},
			}, nil
		}
	} else {
		_, startBlockV := txService.FindTransactionBlockContaining(startDate, accountID)
		_, endBlockV := txService.FindTransactionBlockContaining(endDate, accountID)
		middleBlocks, _ := txService.blockMetumRepo.FindAllBlockMetaInBetween(startDate, endDate, accountID)

		// if err != nil && err.Error() != "record not found" {
		// 	return []types.AccountBlockType{}, err
		// }

		if startBlockV != nil && endBlockV != nil && startBlockV.ID == endBlockV.ID {
			startBlock = startBlockV
			endBlock = nil
		} else {
			startBlock = startBlockV
			endBlock = endBlockV
		}

		listOfMetaInfo = append(listOfMetaInfo, middleBlocks...)
	}

	for _, metaInfo := range listOfMetaInfo {
		journals, _ := txService.journalRepo.FindAllByTransactionId(metaInfo.TransitionTxID)

		// if err != nil {
		// 	return []types.AccountBlockType{}, err
		// }

		if len(journals) != 2 {
			return []types.AccountBlockType{}, errors.New("invalid transition transaction journal list")
		}

		if journals[0].Type != string(types.DEBIT) && journals[1].Type != string(types.DEBIT) {
			return []types.AccountBlockType{}, errors.New("transition transaction journal must contain a debit and a credit entry")
		}

		if journals[0].AccountNumber == account.AccountNumber {
			if !accountBlockIDs.Exists(journals[0].BlockID) {
				accountBlockIds = append(accountBlockIds, journals[0].BlockID)
				accountBlockIDs.Add(journals[0].BlockID)
			}
		}

		if journals[1].AccountNumber == account.AccountNumber {
			if !accountBlockIDs.Exists(journals[1].BlockID) {
				accountBlockIds = append(accountBlockIds, journals[1].BlockID)
				accountBlockIDs.Add(journals[1].BlockID)
			}
		}
	}

	accountBlks, _ := txService.accountBlockRepo.FindAllByIDs(accountBlockIds)

	// if err != nil {
	// 	return []types.AccountBlockType{}, err
	// }

	if startBlock != nil && !accountBlockIDs.Exists(startBlock.ID) {
		accountBlocks = append(accountBlocks, *startBlock)
	}

	for _, acctBlk := range accountBlks {
		accountBlocks = append(accountBlocks, types.AccountBlockType{
			ID:                acctBlk.ID,
			Status:            types.AccountBlockStatus(acctBlk.Status),
			AccountID:         acctBlk.AccountID,
			BlockSize:         acctBlk.BlockSize,
			TransactionsCount: acctBlk.TransactionsCount,
			CreatedAt:         acctBlk.CreatedAt.String(),
		})
	}

	if endBlock != nil && !accountBlockIDs.Exists(endBlock.ID) {
		accountBlocks = append(accountBlocks, *endBlock)
	}

	return accountBlocks, nil
}

func (txService *TransactionService) ComputeBalanceInTransactionBlock(blockId string, stopDate int64) (int, string) {
	journals, err := txService.journalRepo.FindAllByBlockId(blockId)

	if err != nil {
		panic(err)
	}

	block, err := txService.accountBlockRepo.FindById(blockId)

	if err != nil {
		panic(err)
	}

	if len(journals) == 0 {
		return 0, "no_transactions"
	}

	if stopDate < journals[0].CreatedAt.UnixMilli() {
		return 0, "outside_range"
	}

	balance := 0

	if journals[0].Memo == "opening-balance-credit" || journals[0].Memo == "opening-balance-debit" {
		if journals[0].Type == string(types.CREDIT) {
			balance = int(journals[0].Amount)
		} else {
			balance -= int(journals[0].Amount)
		}
	}

	startIndex := 1
	n := len(journals) - 1

	if block.Status == string(types.OPEN) {
		startIndex = 0
		n = len(journals)
		balance = 0
	}

	for i := startIndex; i < n; i++ {
		if journals[i].CreatedAt.UnixMilli() > stopDate {
			return balance, "stop_date_reached"
		}

		if journals[i].Type == string(types.DEBIT) {
			balance -= int(journals[i].Amount)
		} else {
			balance = int(journals[i].Amount)
		}
	}

	return balance, "finished"
}

func (txService *TransactionService) GetOpeningBalance(blockId string) int {
	journals, err := txService.journalRepo.FindAllByBlockId(blockId)

	if err != nil {
		return 0
	}

	if len(journals) == 0 {
		return 0
	}

	if journals[0].Memo == "opening-balance-credit" || journals[0].Memo == "opening-balance-debit" {
		return int(journals[0].Amount)
	}

	if len(journals) > 1 && (journals[1].Memo == "opening-balance-credit" || journals[1].Memo == "opening-balance-debit") {
		return int(journals[1].Amount)
	}

	return 0
}

func (txService *TransactionService) GetJournalEntries(blockId string) []*model.JournalEntry {
	journals, err := txService.journalRepo.FindAllByBlockId(blockId)

	if err != nil {
		panic(err)
	}

	return journals
}

func (txService *TransactionService) AccountBalanceAsAt(transactionDate int64, accountID string) int {
	transactionBlocks, err := txService.FindTransactionBlocks(accountID, transactionDate, transactionDate)

	if err != nil {
		panic(err)
	}

	if len(transactionBlocks) == 0 {
		panic("no account balance for the specified date!!!")
	}

	balance, _ := txService.ComputeBalanceInTransactionBlock(transactionBlocks[len(transactionBlocks)-1].ID, transactionDate)

	return balance
}

func (txService *TransactionService) AccountBlanceAtEndOfDay(transactioinDate int64, accountID string) int {
	return txService.AccountBalanceAsAt(utils.EndOfDay(transactioinDate), accountID)
}

func (txService *TransactionService) ComputeAccountStatementsBetween(startDate, endDate int64, accountID string) ([]types.AccountStatementItem, error) {
	statements := []types.AccountStatementItem{}
	transactionBlocks, err := txService.FindTransactionBlocks(accountID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	if len(transactionBlocks) == 0 {
		return []types.AccountStatementItem{}, nil
	}

	balance, balanceStatus := txService.ComputeBalanceInTransactionBlock(transactionBlocks[0].ID, startDate)
	var openningStatements []types.AccountStatementItem

	m := len(transactionBlocks)

	if balanceStatus == "outside_range" {
		openningBal := txService.GetOpeningBalance(transactionBlocks[0].ID)
		openningStatementss, closingBal := txService.computeAccountStatements(txService.GetJournalEntries(transactionBlocks[0].ID), openningBal, 0)
		balance = closingBal
		openningStatements = openningStatementss
	} else {
		openningStatementss, closingBal := txService.computeAccountStatements(txService.GetJournalEntries(transactionBlocks[0].ID), balance, startDate)
		openningStatements = openningStatementss
		balance = closingBal
	}

	statements = append(statements, openningStatements...)

	for i := 1; i < m-1; i++ {
		journalEntries := txService.GetJournalEntries(transactionBlocks[i].ID)
		newStatements, newBalance := txService.computeAccountStatements(journalEntries, balance, 0)
		balance = newBalance
		statements = append(statements, newStatements...)
	}

	entries := txService.GetJournalEntries(transactionBlocks[m-1].ID)
	newStatements, _ := txService.computeAccountStatements(entries, balance, endDate)
	statements = append(statements, newStatements...)

	return statements, nil
}

func (txService *TransactionService) computeAccountStatements(journalEntries []*model.JournalEntry, openningBalance int, stopDate int64) ([]types.AccountStatementItem, int) {
	statements := []types.AccountStatementItem{}
	balance := openningBalance

	startIndex := 0
	n := len(journalEntries)
	endIndex := n

	if n == 0 {
		return statements, 0
	}

	for i := startIndex; i < endIndex; i++ {
		entry := journalEntries[i]
		if stopDate > 0 && entry.CreatedAt.UnixMilli() > int64(stopDate) {
			break
		}

		if journalEntries[i].Memo == "opening-balance-credit" || journalEntries[i].Memo == "opening-balance-debit" {
			continue
		}

		if journalEntries[i].Memo == "closing-balance-credit" || journalEntries[i].Memo == "closing-balance-debit" {
			continue
		}

		if entry.Type == string(types.DEBIT) {
			balance -= int(entry.Amount)
		} else {
			balance += int(entry.Amount)
		}

		statements = append(statements, types.AccountStatementItem{
			Amount:      int(entry.Amount),
			JournalType: types.JournalType(entry.Type),
			Balance:     balance,
			Date:        entry.CreatedAt.String(),
			Memo:        entry.Memo,
		})
	}

	return statements, balance
}

func (txService *TransactionService) spawnNewAccountBlock(account *types.AccountRepresentation, queryBdTx types.IDBTransaction) *types.AccountRepresentation {
	oldBlockId := account.CurrentActiveBlockId
	newAccount := txService.accountService.InitializeNewBlock(account.AccountNumber, queryBdTx)
	transactionRepo := txService.transactionRepo.WithTransaction(queryBdTx)

	if newAccount.CurrentActiveBlockId == oldBlockId {
		return newAccount
	}

	transitionTransaction, err := transactionRepo.Create(types.CreateLedgerTransaction{
		Status: types.APPROVED,
	})

	if err != nil {
		panic("unable to create transition transaction")
	}

	journalType1 := types.CREDIT
	memoText1 := ""

	journalType2 := types.CREDIT
	memoText2 := ""

	amount := account.Balance

	if account.Balance >= 0 {
		journalType1 = types.CREDIT
		memoText1 = "opening-balance-credit"
		journalType2 = types.DEBIT
		memoText2 = "closing-balance-debit"
	} else {
		journalType1 = types.DEBIT
		memoText1 = "opening-balance-debit"
		journalType2 = types.CREDIT
		memoText2 = "closing-balance-credit"
		amount *= -1
	}

	txService.postTransactionToBlock(types.TransactionInputEntry{
		TransactionEntry: types.TransactionEntry{
			Amount: amount,
			Type:   journalType1,
		},
		Memo:    memoText1,
		OwnerId: newAccount.OwnerId,
	}, newAccount.CurrentActiveBlockId, transitionTransaction.ID, account.AccountNumber, queryBdTx)

	txService.postTransactionToBlock(types.TransactionInputEntry{
		TransactionEntry: types.TransactionEntry{
			Amount: amount,
			Type:   journalType2,
		},
		Memo:    memoText2,
		OwnerId: newAccount.OwnerId,
	}, oldBlockId, transitionTransaction.ID, account.AccountNumber, queryBdTx)

	oldBlock, ftchErr := txService.accountBlockRepo.FindById(oldBlockId)

	if ftchErr != nil {
		panic("unable to retrieve old block")
	}

	fmt.Println("Previous BlockTransactionsCount: " + string(oldBlock.TransactionsCount))

	txBlockMetaRepo := txService.blockMetumRepo.WithTransaction(queryBdTx)
	txBlockMetaRepo.Create(types.CreateBlockMetum{
		AccountId:      account.ID,
		BlockTxLimit:   int(oldBlock.BlockSize),
		TransitionTxId: transitionTransaction.ID,
		OpeningDate:    oldBlock.CreatedAt,
		ClosingDate:    time.Now(),
	})

	return newAccount
}

func (txService *TransactionService) postTransaction(entry types.TransactionInputEntry, transactionId string, account types.AccountRepresentation, dbQueryTx types.IDBTransaction) {
	blockRepo := txService.accountBlockRepo.WithTransaction(dbQueryTx)
	block, err := blockRepo.FindById(account.CurrentActiveBlockId)

	if err != nil {
		panic("invalid block")
	}

	accountData := &account

	if block.TransactionsCount >= block.BlockSize {
		fmt.Println("Transaction block full")
		accountData = txService.spawnNewAccountBlock(&account, dbQueryTx)
	}

	txService.postTransactionToBlock(entry, accountData.CurrentActiveBlockId, transactionId, accountData.AccountNumber, dbQueryTx)
}

func (txService *TransactionService) postTransactionToBlock(entry types.TransactionInputEntry, blockId string, transactionId string, accountNumber string, dbQueryTx types.IDBTransaction) {
	blockRepo := txService.accountBlockRepo.WithTransaction(dbQueryTx)
	block, err := blockRepo.FindById(blockId)

	if err != nil {
		panic("invalid block")
	}

	journalEntryRepo := txService.journalRepo.WithTransaction(dbQueryTx)
	_, createErr := journalEntryRepo.Create(types.CreateJournalEntry{
		Amount:        entry.Amount,
		Type:          entry.Type,
		BlockId:       block.ID,
		TransactionId: transactionId,
		AccountNumber: accountNumber,
		Memo:          entry.Memo,
		OwnerId:       entry.OwnerId,
		// OrganizationId: entry.OrganizationId,
	})

	if createErr != nil {
		panic("unable to add journal entry")
	}

	block.TransactionsCount += 1

	if updateErr := blockRepo.Update(block); updateErr != nil {
		panic("unable to update block transanctions count")
	}
}
