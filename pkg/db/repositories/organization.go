package repositories

import (
	"context"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IOrganizationRepository interface {
	types.IBaseRepository[IOrganizationRepository]
	Create(input types.CreateOrganization) (*model.Organization, error)
	FindById(id string) (*model.Organization, error)
	FindByOwnerId(id string) (*model.Organization, error)
}

type OrganizationRepository struct {
	dbQuery *dao.Query
}

func NewOrganizationRepository() *OrganizationRepository {
	dbInstance := config.DbInstance().GetDBQuery()

	return &OrganizationRepository{
		dbQuery: dbInstance,
	}
}

func (orgRepo *OrganizationRepository) WithTransaction(queryTx types.IDBTransaction) IOrganizationRepository {
	return &OrganizationRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (orgRepo *OrganizationRepository) BeginTransaction() types.IDBTransaction {
	return orgRepo.dbQuery.Begin()
}

func (orgRepo *OrganizationRepository) Create(input types.CreateOrganization) (*model.Organization, error) {
	dbInstance := orgRepo.dbQuery
	organization := dbInstance.Organization.WithContext(context.Background())

	// TODO: check whether this email address exists in the users table

	data := &model.Organization{
		Name:         input.Name,
		Address:      input.Address,
		EmailAddress: input.EmailAddress,
		PhoneNumber:  input.PhoneNumber,
		OwnerID:      input.OwnerID,
	}

	if err := organization.Create(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (apiRepo *OrganizationRepository) FindById(id string) (*model.Organization, error) {
	dbInstance := apiRepo.dbQuery
	org := dbInstance.Organization.WithContext(context.Background())

	orgData, err := org.Where(dbInstance.Organization.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return orgData, nil
}

func (apiRepo *OrganizationRepository) FindByOwnerId(id string) (*model.Organization, error) {
	dbInstance := apiRepo.dbQuery
	org := dbInstance.Organization.WithContext(context.Background())

	orgData, err := org.Where(dbInstance.Organization.OwnerID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return orgData, nil
}
