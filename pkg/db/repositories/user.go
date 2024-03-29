package repositories

import (
	"context"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	types.IBaseRepository[IUserRepository]
	Create(input types.CreateUser) (*model.User, error)
	FindByPhoneNumber(phoneNumber string) (*model.User, error)
	FindByEmail(emailAddress string) (*model.User, error)
	FindByEmailAndPhone(emailAddress, phoneNumber string) (*model.User, error)
	FindByID(id string) (*model.User, error)
}

type UserRepository struct {
	dbQuery *dao.Query
}

func NewUserRepository() *UserRepository {
	dbInstance := config.DbInstance().GetDBQuery()

	return &UserRepository{
		dbQuery: dbInstance,
	}
}

func (userRepo *UserRepository) WithTransaction(queryTx types.IDBTransaction) IUserRepository {
	return &UserRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (userRepo *UserRepository) BeginTransaction() types.IDBTransaction {
	return userRepo.dbQuery.Begin()
}

func (userRepo *UserRepository) Create(input types.CreateUser) (*model.User, error) {
	dbInstance := config.DbInstance().GetDBQuery()
	user := dbInstance.User.WithContext(context.Background())

	hashedPassword, err := userRepo.hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	userData := &model.User{
		FullName:     input.FullName,
		EmailAddress: input.EmailAddress,
		Password:     hashedPassword,
		PhoneNumber:  input.PhoneNumber,
	}

	if err := user.Create(userData); err != nil {
		return nil, err
	}

	return userData, nil
}

func (userRepo *UserRepository) FindByEmail(emailAddress string) (*model.User, error) {
	dbInstance := userRepo.dbQuery
	user := dbInstance.User.WithContext(context.Background())

	userData, err := user.Where(dbInstance.User.EmailAddress.Eq(emailAddress)).First()

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (userRepo *UserRepository) FindByPhoneNumber(phoneNumber string) (*model.User, error) {
	dbInstance := userRepo.dbQuery
	user := dbInstance.User.WithContext(context.Background())

	userData, err := user.Where(dbInstance.User.PhoneNumber.Eq(phoneNumber)).First()

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (userRepo *UserRepository) FindByEmailAndPhone(emailAddress, phoneNumber string) (*model.User, error) {
	dbInstance := userRepo.dbQuery
	user := dbInstance.User.WithContext(context.Background())

	userData, err := user.Where(dbInstance.User.EmailAddress.Eq(emailAddress), dbInstance.User.EmailAddress.Eq(phoneNumber)).First()

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (userRepo *UserRepository) FindByID(id string) (*model.User, error) {
	dbInstance := userRepo.dbQuery
	user := dbInstance.User.WithContext(context.Background())

	userData, err := user.Where(dbInstance.User.PhoneNumber.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (userRepo *UserRepository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func (userRepo *UserRepository) comparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
