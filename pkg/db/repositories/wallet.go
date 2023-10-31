package repositories

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type IWalletRepository interface {
	types.IBaseRepository[IWalletRepository]
	Create(input types.CreateWallet) (*model.Wallet, error)
	FindById(id string) (*model.Wallet, error)
	FindByAccountNumber(accountNumber string) (*model.Wallet, error)
	FindAllByOwnerId(ownerId string) ([]*model.Wallet, error)
	Update(data *model.Wallet) error
	AddLedgerAccounts(accountNumber string, accounts []string) error
}

type WalletRepository struct {
	dbQuery *dao.Query
}

func NewWalletRepository() *WalletRepository {
	dbInstance := config.DbInstance().GetDBQuery()
	return &WalletRepository{
		dbQuery: dbInstance,
	}
}

func (walletRepo *WalletRepository) WithTransaction(queryTx types.IDBTransaction) IWalletRepository {
	return &WalletRepository{
		dbQuery: queryTx.(*dao.QueryTx).Query,
	}
}

func (repo *WalletRepository) BeginTransaction() types.IDBTransaction {
	return repo.dbQuery.Begin()
}

func (walletRepo *WalletRepository) Create(input types.CreateWallet) (*model.Wallet, error) {
	dbInstance := walletRepo.dbQuery
	wallet := dbInstance.Wallet.WithContext(context.Background())

	jsonStr, _ := json.Marshal(input.LedgerAccounts)

	createdWallet := &model.Wallet{
		Type:           string(input.WalletType),
		Name:           input.Name,
		OwnerID:        input.OwnerId,
		AccountNumber:  input.AccountNumber,
		LedgerAccounts: string(jsonStr),
	}

	if err := wallet.Create(createdWallet); err != nil {
		return nil, err
	}

	return createdWallet, nil
}

func (walletRepo *WalletRepository) FindById(id string) (*model.Wallet, error) {
	dbInstance := walletRepo.dbQuery
	wallet := dbInstance.Wallet.WithContext(context.Background())

	w, err := wallet.Where(dbInstance.Wallet.ID.Eq(id)).First()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (walletRepo *WalletRepository) FindByAccountNumber(accountNumber string) (*model.Wallet, error) {
	dbInstance := walletRepo.dbQuery
	wallet := dbInstance.Wallet.WithContext(context.Background())

	w, err := wallet.Where(dbInstance.Wallet.AccountNumber.Eq(accountNumber)).First()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (walletRepo *WalletRepository) FindAllByOwnerId(ownerId string) ([]*model.Wallet, error) {
	dbInstance := walletRepo.dbQuery
	wallet := dbInstance.Wallet.WithContext(context.Background())

	w, err := wallet.Where(dbInstance.Wallet.OwnerID.Eq(ownerId)).Find()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (ledger *WalletRepository) Update(data *model.Wallet) error {
	if data.ID == "" {
		return errors.New("unable to update wallet data without primary ID")
	}

	wallet := ledger.dbQuery.Wallet.WithContext(context.Background())

	if err := wallet.Save(data); err != nil {
		return err
	}

	return nil
}

func (walletRepo *WalletRepository) AddLedgerAccounts(accountNumber string, accounts []string) error {
	wallet, err := walletRepo.FindByAccountNumber(accountNumber)

	if err != nil {
		return err
	}

	var ledgerAccounts []string

	err = json.Unmarshal([]byte(wallet.LedgerAccounts), &ledgerAccounts)

	if err != nil {
		return err
	}

	for _, acct := range accounts {
		_, exists := utils.GetArrayItemIndex[string](acct, ledgerAccounts)

		if !exists {
			ledgerAccounts = append(ledgerAccounts, acct)
		}
	}

	jsonStr, err := json.Marshal(ledgerAccounts)

	if err != nil {
		return err
	}

	wallet.LedgerAccounts = string(jsonStr)
	err = walletRepo.Update(wallet)

	if err != nil {
		return err
	}

	return nil

}
