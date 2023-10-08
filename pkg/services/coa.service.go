package services

import (
	"errors"

	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type CoaService struct {
	CoaRepository repositories.IChartOfAccount
}

type ICoaService interface {
	CreateAccount(input types.CreateCoaAccount) (id string, accountNumber string)
	FindByName(accountName string) (*types.ChartOfAccount, error)
	FindByAccountNumber(accountNumber string)
	ListAll() ([]*types.ChartOfAccount, error)
}

func NewCoaService(CoaRepository repositories.IChartOfAccount) *CoaService {
	return &CoaService{
		CoaRepository: CoaRepository,
	}
}

func (coaService *CoaService) CreateAccount(input types.CreateCoaAccount) (id string, accountNumber string) {
	accountNum := generateAccountNumber(input.AccountType)
	account, err := coaService.CoaRepository.Create(input.Name, accountNum, string(input.AccountType), input.Description)

	if err != nil {
		return "", ""
	}

	return account.ID, account.AccountNumber

}

func (coaService *CoaService) FindByName(accountName string) (*types.ChartOfAccount, error) {
	coaAccount, err := coaService.CoaRepository.FindByName(accountName)

	if err != nil {
		return nil, err
	}

	return &types.ChartOfAccount{
		Id:            coaAccount.ID,
		AccountNumber: coaAccount.AccountNumber,
		CreatedAt:     coaAccount.CreatedAt.String(),
	}, nil

}

func (coaService *CoaService) FindByAccountNumber(accountNumber string) (*types.ChartOfAccount, error) {
	coaAccount, err := coaService.CoaRepository.FindByAccountNumber(accountNumber)

	if err != nil {
		return nil, err
	}

	return &types.ChartOfAccount{
		Id:            coaAccount.ID,
		AccountNumber: coaAccount.AccountNumber,
		CreatedAt:     coaAccount.CreatedAt.String(),
	}, nil

}

func (coaService *CoaService) ListAll() ([]*types.ChartOfAccount, error) {
	coaAccounts, err := coaService.CoaRepository.FindAll()

	if err != nil {
		return nil, err
	}

	accounts := make([]*types.ChartOfAccount, len(coaAccounts))

	for index, acct := range coaAccounts {
		accounts[index] = &types.ChartOfAccount{
			Id:            acct.ID,
			AccountNumber: acct.AccountNumber,
			CreatedAt:     acct.CreatedAt.String(),
		}
	}

	return accounts, nil

}

func generateAccountNumber(accountType types.LedgerAccountType) string {
	accountPrefix, err := getAccountTypePrefix(accountType)

	if err != nil {
		panic(err)
	}

	accountGenerator := utils.CustomAlphabet("0123456789", 10)
	accountId, _ := accountGenerator()

	return accountPrefix + accountId
}

func getAccountTypePrefix(accountType types.LedgerAccountType) (string, error) {
	switch accountType {
	case types.ASSET:
		return "1", nil
	case types.EQUITY:
		return "2", nil
	case types.EXPENSIS:
		return "3", nil
	case types.LIABILITY:
		return "4", nil
	case types.REVENUE:
		return "5", nil
	default:
		return "", errors.New("invalid account type")
	}
}
