package repositories

import (
	"context"
	"errors"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type ILedgerAccount interface {
	WithTransaction(queryTx *dao.QueryTx) *LedgerAccountRepository
	Create(input types.CreateLedgerAccount) (*model.LedgerAccount, error)
	FindById(id string) (*model.LedgerAccount, error)
	FindByAccountNumber(accountNumber string) (*model.LedgerAccount, error)
	Update(data *model.LedgerAccount) error
}

type LedgerAccountRepository struct {
	dbQuery *dao.Query
}

func NewLedgerAccountRepository() *LedgerAccountRepository {
	var dbInstance = config.DbInstance().GetDBQuery()
	return &LedgerAccountRepository{
		dbQuery: dbInstance,
	}
}

func (ledger *LedgerAccountRepository) WithTransaction(queryTx *dao.QueryTx) *LedgerAccountRepository {
	return &LedgerAccountRepository{
		dbQuery: queryTx.Query,
	}
}

func (ledger *LedgerAccountRepository) Create(input types.CreateLedgerAccount) (*model.LedgerAccount, error) {
	var dbInstance = ledger.dbQuery
	ctx := context.Background()
	ledgerAccount := dbInstance.LedgerAccount.WithContext(ctx)

	createdAccount := &model.LedgerAccount{
		AccountNumber:        input.AccountNumber,
		OwnerID:              input.OwnerId,
		Book:                 string(input.Book),
		CurrentActiveBlockID: input.CurrentActiveBlockId,
		Status:               string(input.Status),
		Label:                input.Label,
		BlockCount:           int32(input.BlockCount),
		Particular:           input.Particular,
	}

	if err := ledgerAccount.Create(createdAccount); err != nil {
		return nil, err
	}

	return createdAccount, nil

}

func (ledger *LedgerAccountRepository) FindById(id string) (*model.LedgerAccount, error) {
	var dbInstance = ledger.dbQuery
	ctx := context.Background()
	ledgerAccount := dbInstance.LedgerAccount.WithContext(ctx)

	account, err := ledgerAccount.Where(dbInstance.LedgerAccount.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (ledger *LedgerAccountRepository) FindByAccountNumber(accountNumber string) (*model.LedgerAccount, error) {
	var dbInstance = ledger.dbQuery
	ctx := context.Background()
	ledgerAccount := dbInstance.LedgerAccount.WithContext(ctx)

	account, err := ledgerAccount.Where(dbInstance.LedgerAccount.AccountNumber.Eq(accountNumber)).First()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (ledger *LedgerAccountRepository) Update(data *model.LedgerAccount) error {
	if data.ID == "" {
		return errors.New("unable to update ledger account data without primary ID")
	}

	ledgerAccount := ledger.dbQuery.LedgerAccount.WithContext(context.Background())

	if err := ledgerAccount.Save(data); err != nil {
		return err
	}

	return nil
}
