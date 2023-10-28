package repositories

import (
	"context"
	"encoding/json"

	"github.com/israel-duff/ledger-system/pkg/config"
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/types"
)

type IWalletRepository interface {
	types.IBaseRepository[IWalletRepository]
	Create(input types.CreateWallet) (*model.Wallet, error)
	FindById(id string) (*model.Wallet, error)
	FindByAccountNumber(accountNumber string) (*model.Wallet, error)
	// FindAllByBlockId(blockId string) ([]*model.JournalEntry, error)
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
