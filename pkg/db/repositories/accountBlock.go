package repositories

import (
	"context"
	"errors"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IAccountBlockRepository interface {
	WithTransaction(queryTx *dao.QueryTx) *AccountBlockRepository
	Create(input types.CreateAccountBlock) (*model.AccountBlock, error)
	FindById(id string) (*model.AccountBlock, error)
	Update(data *model.AccountBlock) error
}

type AccountBlockRepository struct {
	dbQuery *dao.Query
}

func NewAccountBlockRepository() *AccountBlockRepository {
	dbInstance := config.DbInstance().GetDBQuery()

	return &AccountBlockRepository{
		dbQuery: dbInstance,
	}
}

func (accountBlockRepo *AccountBlockRepository) WithTransaction(queryTx *dao.QueryTx) *AccountBlockRepository {
	return &AccountBlockRepository{
		dbQuery: queryTx.Query,
	}
}

func (accountBlockRepo *AccountBlockRepository) Create(input types.CreateAccountBlock) (*model.AccountBlock, error) {
	dbInstance := accountBlockRepo.dbQuery
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

func (accountBlockRepo *AccountBlockRepository) FindById(id string) (*model.AccountBlock, error) {
	dbInstance := accountBlockRepo.dbQuery
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	acctBlock, err := accountBlock.Where(dbInstance.AccountBlock.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return acctBlock, nil
}

func (accountBlockRep *AccountBlockRepository) Update(data *model.AccountBlock) error {
	if data.ID == "" {
		return errors.New("can't update account block without primary ID")
	}

	dbInstance := accountBlockRep.dbQuery
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	if err := accountBlock.Save(data); err != nil {
		return err
	}

	return nil
}
