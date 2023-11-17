package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
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
	blockMetumRepo   repositories.IBlockMetumRepository
	accountService   IAccountService
	txQ              ITransactionQService
}

func NewTransactionService(
	transactionRepo repositories.ILedgerTransactionRepository,
	journalRep repositories.IJournalEntryRepository,
	accountBlockRepo repositories.IAccountBlockRepository,
	blockMetumRepo repositories.IBlockMetumRepository,
	txQ ITransactionQService,
	accountService IAccountService) *TransactionService {
	return &TransactionService{
		transactionRepo:  transactionRepo,
		journalRepo:      journalRep,
		accountBlockRepo: accountBlockRepo,
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

	dbQueryTx.Commit()

	for _, r := range schedulerResp {
		r.ReleaseLock()
	}

	return nil
}

func (txService *TransactionService) FindTransactionBlockContaining(transactionDate int32) (types.BlockSearchStatus, *types.AccountBlockType) {
	leftBlock, leftErr := txService.accountBlockRepo.FindByDate(transactionDate, "left")
	rightBlock, rightErr := txService.accountBlockRepo.FindByDate(transactionDate, "right")

	if leftErr != nil || rightErr != nil {
		panic("error fetching blocks")
	}

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
		return types.WITHIN_BLOCK, &types.AccountBlockType{
			ID:                rightBlock.ID,
			Status:            types.AccountBlockStatus(rightBlock.Status),
			AccountID:         rightBlock.AccountID,
			BlockSize:         rightBlock.BlockSize,
			TransactionsCount: rightBlock.TransactionsCount,
			CreatedAt:         rightBlock.CreatedAt.String(),
		}
	}

	if rightBlock != nil {
		return types.RIGHT_ONLY, nil
	}

	if leftBlock != nil {
		return types.LEFT_ONLY, nil
	}

	panic("error processing blocks")
}

func (txService *TransactionService) FindTransactionBlocks(accountID string, startDate, endDate int32) ([]types.AccountBlockType, error) {
	listOfMetaInfo := []*model.BlockMetum{}
	accountBlockIds := []string{}
	var startBlock *types.AccountBlockType
	var endBlock *types.AccountBlockType
	accountBlocks := []types.AccountBlockType{}

	if startDate == endDate {
		startBlkStatus, startBlk := txService.FindTransactionBlockContaining(startDate)

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
		_, startBlockV := txService.FindTransactionBlockContaining(startDate)
		_, endBlockV := txService.FindTransactionBlockContaining(endDate)
		middleBlocks, err := txService.blockMetumRepo.FindAllBlockMetaInBetween(startDate, endDate)

		if err != nil {
			return []types.AccountBlockType{}, err
		}

		startBlock = startBlockV
		endBlock = endBlockV

		listOfMetaInfo = append(listOfMetaInfo, middleBlocks...)
	}

	for _, metaInfo := range listOfMetaInfo {
		journals, err := txService.journalRepo.FindAllByTransactionId(metaInfo.TransitionTxID)

		if err != nil {
			return []types.AccountBlockType{}, err
		}

		if len(journals) != 2 {
			return []types.AccountBlockType{}, errors.New("invalid transition transaction journal list")
		}

		if journals[0].Type != string(types.DEBIT) && journals[1].Type != string(types.DEBIT) {
			return []types.AccountBlockType{}, errors.New("transition transaction journal must contain a debit and a credit entry")
		}

		if journals[0].Type == string(types.DEBIT) {
			accountBlockIds = append(accountBlockIds, journals[0].BlockID)
		} else {
			accountBlockIds = append(accountBlockIds, journals[1].BlockID)
		}
	}

	accountBlks, err := txService.accountBlockRepo.FindAllByIDs(accountBlockIds)

	if err != nil {
		return []types.AccountBlockType{}, err
	}

	if startBlock != nil {
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

	if endBlock != nil {
		accountBlocks = append(accountBlocks, *endBlock)
	}

	return accountBlocks, nil
}

func (txService *TransactionService) ComputeBalanceInTransactionBlock(blockId string, stopDate int32) int {
	journals, err := txService.journalRepo.FindAllByBlockId(blockId)

	if err != nil {
		panic(err)
	}

	block, err := txService.accountBlockRepo.FindById(blockId)

	if err != nil {
		panic(err)
	}

	if len(journals) == 0 {
		return 0
	}

	if stopDate < int32(journals[0].CreatedAt.UnixMilli()) {
		return 0
	}

	balance := 0

	if journals[0].Type == string(types.CREDIT) {
		balance = int(journals[0].Amount)
	} else {
		balance -= int(journals[0].Amount)
	}

	startIndex := 1
	n := len(journals) - 1

	if block.Status == string(types.OPEN) {
		startIndex = 0
		n = len(journals)
		balance = 0
	}

	for i := startIndex; i < n; i++ {
		if int32(journals[i].CreatedAt.UnixMilli()) > stopDate {
			return balance
		}

		if journals[i].Type == string(types.DEBIT) {
			balance -= int(journals[i].Amount)
		} else {
			balance = int(journals[i].Amount)
		}
	}

	return balance
}

func (txService *TransactionService) GetOpeningBalance(blockId string) int {
	journals, err := txService.journalRepo.FindAllByBlockId(blockId)

	if err != nil {
		return 0
	}

	if len(journals) == 0 {
		return 0
	}

	return int(journals[0].Amount)
}

func (txService *TransactionService) GetJournalEntries(blockId string) []*model.JournalEntry {
	journals, err := txService.journalRepo.FindAllByBlockId(blockId)

	if err != nil {
		panic(err)
	}

	return journals
}

func (txService *TransactionService) AccountBalanceAsAt(transactionDate int32, accountID string) int {
	transactionBlocks, err := txService.FindTransactionBlocks(accountID, transactionDate, transactionDate)

	if err != nil {
		panic(err)
	}

	if len(transactionBlocks) == 0 {
		panic("no account balance for the specified date!!!")
	}

	balance := txService.ComputeBalanceInTransactionBlock(transactionBlocks[len(transactionBlocks)-1].ID, transactionDate)

	return balance
}

func (txService *TransactionService) AccountBlanceAtEndOfDay(transactioinDate int32, accountID string) int {
	return txService.AccountBalanceAsAt(int32(utils.EndOfDay(int64(transactioinDate))), accountID)
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
		OpeningDate:    oldBlock.CreatedAt.String(),
		ClosingDate:    time.Now().String(),
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
