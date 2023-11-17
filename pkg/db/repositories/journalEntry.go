package repositories

import (
	"context"
	"fmt"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IJournalEntryRepository interface {
	types.IBaseRepository[IJournalEntryRepository]
	Create(input types.CreateJournalEntry) (*model.JournalEntry, error)
	FindById(id string) (*model.JournalEntry, error)
	FindAllByBlockId(blockId string) ([]*model.JournalEntry, error)
	FindAllByTransactionId(transactioinId string) ([]*model.JournalEntry, error)
}

type JournalEntryRepository struct {
	dbQuery *dao.Query
}

func NewJournalEntryRepository() *JournalEntryRepository {
	dbInstance := config.DbInstance().GetDBQuery()
	return &JournalEntryRepository{
		dbQuery: dbInstance,
	}
}

func (journalEntryRepo *JournalEntryRepository) WithTransaction(queryTx types.IDBTransaction) IJournalEntryRepository {
	return &JournalEntryRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (repo *JournalEntryRepository) BeginTransaction() types.IDBTransaction {
	return repo.dbQuery.Begin()
}

func (journalEntryRepo *JournalEntryRepository) Create(input types.CreateJournalEntry) (*model.JournalEntry, error) {
	dbInstance := journalEntryRepo.dbQuery
	journalEntry := dbInstance.JournalEntry.WithContext(context.Background())

	createdJournalEntry := &model.JournalEntry{
		Amount:        float64(input.Amount),
		Type:          string(input.Type),
		BlockID:       input.BlockId,
		TransactionID: input.TransactionId,
		Name:          input.Name,
		Memo:          input.Memo,
		OwnerID:       input.OwnerId,
		// OrganizationID: input.OrganizationId,
		OrganizationID: "d879361c-53a9-4fda-9c23-baefcecb1753",
		AccountNumber:  input.AccountNumber,
	}

	if err := journalEntry.Create(createdJournalEntry); err != nil {
		fmt.Println("INPUT>>>>>>>")
		fmt.Println("BlockID:" + input.BlockId)
		fmt.Println("TransactionID:" + input.TransactionId)
		fmt.Println("OwnerID:" + input.OwnerId)
		fmt.Println("OrganizationID:" + input.OrganizationId)
		fmt.Println(err)
		fmt.Println(">>>>>>>>>>>>")
		return nil, err
	}

	return createdJournalEntry, nil
}

func (journalEntryRepo *JournalEntryRepository) FindById(id string) (*model.JournalEntry, error) {
	dbInstance := journalEntryRepo.dbQuery
	journalEntry := dbInstance.JournalEntry.WithContext(context.Background())

	jEntry, err := journalEntry.Where(dbInstance.JournalEntry.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return jEntry, nil
}

func (journalEntryRepo *JournalEntryRepository) FindAllByBlockId(blockId string) ([]*model.JournalEntry, error) {
	dbInstance := journalEntryRepo.dbQuery
	journalEntry := dbInstance.JournalEntry.WithContext(context.Background())

	jEntries, err := journalEntry.Where(dbInstance.JournalEntry.BlockID.Eq(blockId)).Find()

	if err != nil {
		return nil, err
	}

	return jEntries, nil
}

func (journalEntryRepo *JournalEntryRepository) FindAllByTransactionId(transactioinId string) ([]*model.JournalEntry, error) {
	dbInstance := journalEntryRepo.dbQuery
	journalEntry := dbInstance.JournalEntry.WithContext(context.Background())

	jEntries, err := journalEntry.Where(dbInstance.JournalEntry.TransactionID.Eq(transactioinId)).Find()

	if err != nil {
		return nil, err
	}

	return jEntries, nil
}
