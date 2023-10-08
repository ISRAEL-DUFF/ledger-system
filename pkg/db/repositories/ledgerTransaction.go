package repositories

import (
	"context"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type ILedgerTransactionRepository interface {
	Create(input types.CreateLedgerTransaction) (*model.LedgerTransaction, error)
	FindById(id string) (*model.LedgerTransaction, error)
}

type LedgerTransactionRepository struct {
}

func (tx *LedgerTransactionRepository) Create(input types.CreateLedgerTransaction) (*model.LedgerTransaction, error) {
	dbQuery := config.DbInstance().GetDBQuery()
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
	dbQuery := config.DbInstance().GetDBQuery()
	ledgerTx := dbQuery.LedgerTransaction.WithContext(context.Background())

	fetchedTx, err := ledgerTx.Where(dbQuery.LedgerTransaction.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return fetchedTx, nil
}
