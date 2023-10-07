package repositories

import (
	"context"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IAccountBlock interface {
	Create(input types.CreateAccountBlock) (*model.AccountBlock, error)
	FindById(id string) (*model.AccountBlock, error)
}

type AccountBlockRepository struct {
}

func NewAccountBlockRepository() *LedgerAccountRepository {
	return &LedgerAccountRepository{}
}

func (accountBlockRepo *AccountBlockRepository) Create(input types.CreateAccountBlock) (*model.AccountBlock, error) {
	dbInstance := config.DbInstance().GetDBQuery()
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	createdAccountBlock := &model.AccountBlock{
		IsCurrentBlock:    input.IsCurrentBlock,
		Status:            string(input.Status),
		TransactionsCount: int32(input.TransactionsCount),
		BlockSize:         int32(input.BlockSize),
		AccountID:         input.AccountId,
	}

	if err := accountBlock.Create(createdAccountBlock); err != nil {
		return nil, err
	}

	return createdAccountBlock, nil
}

func (accountBlockRep *AccountBlockRepository) FindById(id string) (*model.AccountBlock, error) {
	dbInstance := config.DbInstance().GetDBQuery()
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	acctBlock, err := accountBlock.Where(dbInstance.AccountBlock.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return acctBlock, nil
}
