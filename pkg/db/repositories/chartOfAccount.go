package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
)

// var dbInstance = config.DbInstance().GetDBQuery()

type ChartOfAccountRepository struct {
	dbQuery *dao.Query
}

type IChartOfAccount interface {
	WithTransaction(queryTx *dao.QueryTx) *ChartOfAccountRepository
	Create(name string, accountNumber string, accountType string, description string) (*model.ChartOfAccount, error)
	FindById(id string) (*model.ChartOfAccount, error)
	FindByName(accountName string) (*model.ChartOfAccount, error)
	FindByAccountNumber(accountName string) (*model.ChartOfAccount, error)
	FindAll() ([]*model.ChartOfAccount, error)
}

// type DAO interface {
//     Create(...)(*model.UserEntity, error)
//     FetchByID(...)(*model.UserEntity, error)
//     FetchAll(...)([]*model.UserEntity, error)
//     Update(...)(*model.UserEntity, error)
//     Delete(...) error
// }

func NewChartOfAccountRepository() *ChartOfAccountRepository {
	var dbInstance = config.DbInstance().GetDBQuery()
	return &ChartOfAccountRepository{
		dbQuery: dbInstance,
	}
}

func (repo *ChartOfAccountRepository) WithTransaction(queryTx *dao.QueryTx) *ChartOfAccountRepository {
	return &ChartOfAccountRepository{
		dbQuery: queryTx.Query,
	}
}

func (repo *ChartOfAccountRepository) Create(name string, accountNumber string, accountType string, description string) (*model.ChartOfAccount, error) {
	var dbInstance = repo.dbQuery
	ctx := context.Background()
	chartOfAccount := dbInstance.ChartOfAccount.WithContext(ctx)

	createdModel := &model.ChartOfAccount{
		Name:          name,
		AccountNumber: accountNumber,
		Type:          accountType,
		Description:   description,
	}

	if err := chartOfAccount.Create(createdModel); err != nil {
		return nil, err
	}

	return createdModel, nil
}

func (repo *ChartOfAccountRepository) FindById(id string) (*model.ChartOfAccount, error) {
	var dbInstance = repo.dbQuery
	ctx := context.Background()
	chartOfAccount := dbInstance.ChartOfAccount.WithContext(ctx)

	data, err := chartOfAccount.Where(dbInstance.ChartOfAccount.ID.Eq(id)).First()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to retrieve item")
	}

	return data, nil
}

func (repo *ChartOfAccountRepository) FindByName(accountName string) (*model.ChartOfAccount, error) {
	var dbInstance = repo.dbQuery
	ctx := context.Background()
	chartOfAccount := dbInstance.ChartOfAccount.WithContext(ctx)

	data, err := chartOfAccount.Where(dbInstance.ChartOfAccount.Name.Eq(accountName)).First()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to retrieve item")
	}

	return data, nil
}

func (repo *ChartOfAccountRepository) FindByAccountNumber(accountNumber string) (*model.ChartOfAccount, error) {
	var dbInstance = repo.dbQuery
	ctx := context.Background()
	chartOfAccount := dbInstance.ChartOfAccount.WithContext(ctx)

	data, err := chartOfAccount.Where(dbInstance.ChartOfAccount.AccountNumber.Eq(accountNumber)).First()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to retrieve item")
	}

	return data, nil
}

func (repo *ChartOfAccountRepository) FindAll() ([]*model.ChartOfAccount, error) {
	var dbInstance = repo.dbQuery
	ctx := context.Background()
	chartOfAccount := dbInstance.ChartOfAccount.WithContext(ctx)

	data, err := chartOfAccount.Find()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to retrieve item")
	}

	return data, nil
}
