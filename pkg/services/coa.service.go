package services

import (
	"errors"

	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type CoaService struct {
}

func NewCoaService() *CoaService {
	return &CoaService{}
}

func (coaService *CoaService) CreateAccount(accountType types.LedgerAccountType, name string, description string) (id string, accountNumber string) {
	accountNum := generateAccountNumber(accountType)

	coaRepo := repositories.NewChartOfAccountRepository()

	account, err := coaRepo.Create(name, accountNum, string(accountType), description)

	if err != nil {
		return "", ""
	}

	return account.ID, account.AccountNumber

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
