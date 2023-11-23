package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IAccountBlockRepository interface {
	types.IBaseRepository[IAccountBlockRepository]
	Create(input types.CreateAccountBlock) (*model.AccountBlock, error)
	FindById(id string) (*model.AccountBlock, error)
	Update(data *model.AccountBlock) error
	FindByDate(transactionDate int64, accountId, direction string) (*model.AccountBlock, error)
	// FindAllBlocksInBetween(startDate, endDate int32) ([]*model.AccountBlock, error)
	GetCurrentOpenBlock(accountID string) (*model.AccountBlock, error)
	FindAllByIDs(accountBlockIDs []string) ([]*model.AccountBlock, error)
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

func (accountBlockRepo *AccountBlockRepository) WithTransaction(queryTx types.IDBTransaction) IAccountBlockRepository {
	return &AccountBlockRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (accountBlockRepo *AccountBlockRepository) BeginTransaction() types.IDBTransaction {
	return accountBlockRepo.dbQuery.Begin()
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

func (accountBlockRepo *AccountBlockRepository) FindByDate(transactionDate int64, accountId, direction string) (*model.AccountBlock, error) {
	dbInstance := accountBlockRepo.dbQuery
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	var acctBlock *model.AccountBlock
	var err error

	txDate := time.UnixMilli(int64(transactionDate))

	fmt.Println("TXDATE:", txDate)

	if direction == "left" {
		acctBlock, err = accountBlock.Where(dbInstance.AccountBlock.AccountID.Eq(accountId), dbInstance.AccountBlock.CreatedAt.Lte(txDate)).Order(dbInstance.AccountBlock.CreatedAt.Desc()).First()

		if err != nil {
			return nil, err
		}

	} else {
		acctBlock, err = accountBlock.Where(dbInstance.AccountBlock.AccountID.Eq(accountId), dbInstance.AccountBlock.CreatedAt.Gte(txDate)).Order(dbInstance.AccountBlock.CreatedAt).First()

		if err != nil {
			return nil, err
		}
	}

	return acctBlock, nil
}

func (accountBlockRepo *AccountBlockRepository) GetCurrentOpenBlock(accountID string) (*model.AccountBlock, error) {
	dbInstance := accountBlockRepo.dbQuery
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	// TODO: check this IsCurrentBlock condition
	acctBlock, err := accountBlock.Where(dbInstance.AccountBlock.AccountID.Eq(accountID), dbInstance.AccountBlock.IsCurrentBlock, dbInstance.AccountBlock.Status.Eq(string(types.OPEN))).Order(dbInstance.AccountBlock.CreatedAt.Desc()).First()

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

func (accountBlockRepo *AccountBlockRepository) FindAllByIDs(accountBlockIDs []string) ([]*model.AccountBlock, error) {
	dbInstance := accountBlockRepo.dbQuery
	accountBlock := dbInstance.AccountBlock.WithContext(context.Background())

	acctBlocks, err := accountBlock.Where(dbInstance.AccountBlock.ID.In(accountBlockIDs...)).Order(dbInstance.AccountBlock.CreatedAt).Find()

	if err != nil {
		return nil, err
	}

	return acctBlocks, nil
}
