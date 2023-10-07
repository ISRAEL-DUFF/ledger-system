package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/model"
)

// var dbInstance = config.DbInstance().GetDBQuery()

type ChartOfAccountRepository struct {
}

type IChartOfAccount interface {
	Create(name string, accountNumber string, accountType string, description string) (*model.ChartOfAccount, error)
	FindById(id string) (*model.ChartOfAccount, error)
}

// type DAO interface {
//     Create(...)(*model.UserEntity, error)
//     FetchByID(...)(*model.UserEntity, error)
//     FetchAll(...)([]*model.UserEntity, error)
//     Update(...)(*model.UserEntity, error)
//     Delete(...) error
// }

func NewChartOfAccountRepository() *ChartOfAccountRepository {
	return &ChartOfAccountRepository{}
}

func (repo *ChartOfAccountRepository) Create(name string, accountNumber string, accountType string, description string) (*model.ChartOfAccount, error) {
	var dbInstance = config.DbInstance().GetDBQuery()
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
	var dbInstance = config.DbInstance().GetDBQuery()
	ctx := context.Background()
	chartOfAccount := dbInstance.ChartOfAccount.WithContext(ctx)

	data, err := chartOfAccount.Where(dbInstance.ChartOfAccount.ID.Eq(id)).First()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to retrieve item")
	}

	return data, nil
}
