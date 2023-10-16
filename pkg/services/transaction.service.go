package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type ITransactionService interface {
	CreateLedgerTransaction(input types.TransactionInput) (types.TransactionResponse, error)
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
	// dbQuery := config.DbInstance().GetDBQuery()

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
				panic("Insufficient fund on account " + entry.AccountNumber)
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

func (txService *TransactionService) CreateQueuedLedgerTransaction(input types.TransactionInput) (types.TransactionResponse, error) {
	r, err := txService.txQ.Schedule(input, func(resInput types.TransactionInput) (types.TransactionResponse, error) {
		resp, err := txService.CreateLedgerTransaction(resInput)

		return resp, err
	})

	return r, err
}

func (txService *TransactionService) postTransactionToBlock(entry types.TransactionInputEntry, blockId string, transactionId string, accountNumber string, dbQueryTx types.IDBTransaction) {
	blockRepo := txService.accountBlockRepo.WithTransaction(dbQueryTx)
	block, err := blockRepo.FindById(blockId)

	if err != nil {
		panic("invalid block")
	}

	journalEntryRepo := txService.journalRepo.WithTransaction(dbQueryTx)
	_, createErr := journalEntryRepo.Create(types.CreateJournalEntry{
		Amount:         entry.Amount,
		Type:           entry.Type,
		BlockId:        block.ID,
		TransactionId:  transactionId,
		AccountNumber:  accountNumber,
		Memo:           entry.Memo,
		OwnerId:        entry.OwnerId,
		OrganizationId: entry.OrganizationId,
	})

	if createErr != nil {
		panic("unable to add journal entry")
	}

	block.TransactionsCount += 1

	if updateErr := blockRepo.Update(block); updateErr != nil {
		panic("unable to update block transanctions count")
	}
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
	}

	txService.postTransactionToBlock(types.TransactionInputEntry{
		TransactionEntry: types.TransactionEntry{
			Amount: account.Balance,
			Type:   journalType1,
		},
		Memo: memoText1,
	}, account.CurrentActiveBlockId, transitionTransaction.ID, account.AccountNumber, queryBdTx)

	txService.postTransactionToBlock(types.TransactionInputEntry{
		TransactionEntry: types.TransactionEntry{
			Amount: account.Balance,
			Type:   journalType2,
		},
		Memo: memoText2,
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
