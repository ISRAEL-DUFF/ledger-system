package repositories

import (
	"context"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type ILedgerTransactionRepository interface {
	types.IBaseRepository[ILedgerTransactionRepository]
	Create(input types.CreateLedgerTransaction) (*model.LedgerTransaction, error)
	FindById(id string) (*model.LedgerTransaction, error)
	UpdateStatus(id string, txStatus types.TransactionStatus) error
}

type LedgerTransactionRepository struct {
	dbQuery *dao.Query
}

func NewLedgerTransactionRepository() *LedgerTransactionRepository {
	dbQuery := config.DbInstance().GetDBQuery()
	return &LedgerTransactionRepository{
		dbQuery: dbQuery,
	}
}

func (tx *LedgerTransactionRepository) WithTransaction(queryTx *dao.QueryTx) *LedgerTransactionRepository {
	return &LedgerTransactionRepository{
		dbQuery: queryTx.Query,
	}
}

func (tx *LedgerTransactionRepository) Create(input types.CreateLedgerTransaction) (*model.LedgerTransaction, error) {
	dbQuery := tx.dbQuery
	ledgerTx := dbQuery.LedgerTransaction.WithContext(context.Background())

	createdTx := &model.LedgerTransaction{
		Status: string(input.Status),
	}

	if err := ledgerTx.Create(createdTx); err != nil {
		return nil, err
	}

	return createdTx, nil
}

func (tx *LedgerTransactionRepository) FindById(id string) (*model.LedgerTransaction, error) {
	dbQuery := tx.dbQuery
	ledgerTx := dbQuery.LedgerTransaction.WithContext(context.Background())

	fetchedTx, err := ledgerTx.Where(dbQuery.LedgerTransaction.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return fetchedTx, nil
}

func (tx *LedgerTransactionRepository) UpdateStatus(id string, txStatus types.TransactionStatus) error {
	dbQuery := tx.dbQuery
	ledgerTx := dbQuery.LedgerTransaction.WithContext(context.Background())

	if _, err := ledgerTx.Update(dbQuery.LedgerTransaction.ID.Eq(id), txStatus); err != nil {
		return err
	}

	return nil
}
