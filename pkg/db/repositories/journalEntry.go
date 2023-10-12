package repositories

import (
	"context"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IJournalEntry interface {
	WithTransaction(queryTx *dao.QueryTx) *JournalEntryRepository
	Create(input types.CreateJournalEntry) (*model.JournalEntry, error)
	FindById(id string) (*model.JournalEntry, error)
	FindAllByBlockId(blockId string) ([]*model.JournalEntry, error)
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

func (journalEntryRepo *JournalEntryRepository) WithTransaction(queryTx *dao.QueryTx) *JournalEntryRepository {
	return &JournalEntryRepository{
		dbQuery: queryTx.Query,
	}
}

func (journalEntryRepo *JournalEntryRepository) Create(input types.CreateJournalEntry) (*model.JournalEntry, error) {
	dbInstance := journalEntryRepo.dbQuery
	journalEntry := dbInstance.JournalEntry.WithContext(context.Background())

	createdJournalEntry := &model.JournalEntry{
		Amount:         float64(input.Amount),
		Type:           string(input.Type),
		BlockID:        input.BlockId,
		TransactionID:  input.TransactionId,
		Name:           input.Name,
		Memo:           input.Memo,
		OwnerID:        input.OwnerId,
		OrganizationID: input.OrganizationId,
		AccountNumber:  input.AccountNumber,
	}

	if err := journalEntry.Create(createdJournalEntry); err != nil {
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
