package repositories

import (
	"context"
	"errors"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IApikeyRepository interface {
	types.IBaseRepository[IApikeyRepository]
	Create(input types.CreateAPIKEY) (*model.APIKey, error)
	Update(data *model.APIKey) error
	FindById(id string) (*model.APIKey, error)
	FindByOrgId(id string) (*model.APIKey, error)
}

type ApikeyRepository struct {
	dbQuery *dao.Query
}

func NewApikeyRepository() *ApikeyRepository {
	dbInstance := config.DbInstance().GetDBQuery()

	return &ApikeyRepository{
		dbQuery: dbInstance,
	}
}

func (apiRepo *ApikeyRepository) WithTransaction(queryTx types.IDBTransaction) IApikeyRepository {
	return &ApikeyRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (apiRepo *ApikeyRepository) BeginTransaction() types.IDBTransaction {
	return apiRepo.dbQuery.Begin()
}

func (apiRepo *ApikeyRepository) Create(input types.CreateAPIKEY) (*model.APIKey, error) {
	dbInstance := config.DbInstance().GetDBQuery()
	apikey := dbInstance.APIKey.WithContext(context.Background())

	userData := &model.APIKey{
		OwnerID:       input.OrganizationID,
		TestSecretKey: input.TestSecretKey,
		TestPublicKey: input.TestPublicKey,
		LivePublicKey: input.LivePublicKey,
		LiveSecretKey: input.LiveSecretKey,
	}

	if err := apikey.Create(userData); err != nil {
		return nil, err
	}

	return userData, nil
}

func (apiRepo *ApikeyRepository) FindById(id string) (*model.APIKey, error) {
	dbInstance := apiRepo.dbQuery
	apikey := dbInstance.APIKey.WithContext(context.Background())

	apiKey, err := apikey.Where(dbInstance.APIKey.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (apiRepo *ApikeyRepository) FindByOrgId(id string) (*model.APIKey, error) {
	dbInstance := apiRepo.dbQuery
	apikey := dbInstance.APIKey.WithContext(context.Background())

	apiKey, err := apikey.Where(dbInstance.APIKey.OwnerID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (apiRepo *ApikeyRepository) Update(data *model.APIKey) error {
	dbInstance := apiRepo.dbQuery
	apikey := dbInstance.APIKey.WithContext(context.Background())

	if data.ID == "" {
		return errors.New("unable to update api key data without primary ID")
	}

	if err := apikey.Save(data); err != nil {
		return err
	}

	return nil
}
