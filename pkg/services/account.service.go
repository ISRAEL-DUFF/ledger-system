package services

import (
	"errors"
	"fmt"

	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type IAccountService interface {
}

type AccountService struct {
	chartOfAccountService ICoaService
	accountBlockRepo      repositories.IAccountBlock
	ledgerAccountRepo     repositories.ILedgerAccount
	journalEntryRepo      repositories.IJournalEntry
}

type CreateLedgerAccountOptionInput struct {
	CoaAccountInput    *types.CreateCoaAccount
	AccountBlockInput  *types.CreateAccountBlock
	LedgerAccountInput *types.CreateLedgerAccount
}

type LedgerAccountOption func(*CreateLedgerAccountOptionInput)

func NewAccountService(chartOfAccountService ICoaService, accountBlock repositories.IAccountBlock, ledgerAccount repositories.ILedgerAccount, journalEntry repositories.IJournalEntry) *AccountService {
	return &AccountService{
		chartOfAccountService: chartOfAccountService,
		accountBlockRepo:      accountBlock,
		ledgerAccountRepo:     ledgerAccount,
		journalEntryRepo:      journalEntry,
	}
}

func (accountService *AccountService) CreateLedgerAccount(name string, ownerId string, opts ...LedgerAccountOption) (string, error) {
	coaPayload := &types.CreateCoaAccount{
		AccountType: types.ASSET,
		Name:        name,
		Description: "Asset Account",
	}
	accountBlockPayload := &types.CreateAccountBlock{
		BlockSize: 50, // TODO: put in .env file
	}
	ledgerAccountPayload := &types.CreateLedgerAccount{
		Book:       types.CASH_RECEIPT,
		Particular: "Asset Account",
	}

	for _, opt := range opts {
		opt(&CreateLedgerAccountOptionInput{
			CoaAccountInput:    coaPayload,
			AccountBlockInput:  accountBlockPayload,
			LedgerAccountInput: ledgerAccountPayload,
		})
	}

	accountId, accountNumber := accountService.chartOfAccountService.CreateAccount(*coaPayload)

	if accountNumber == "" {
		return "", errors.New("unable to create coa account")
	}

	accountBlockPayload.AccountId = accountId
	accountBlockPayload.IsCurrentBlock = true
	accountBlockPayload.Status = types.OPEN
	accountBlockPayload.TransactionsCount = 0

	block, err := accountService.accountBlockRepo.Create(*accountBlockPayload)

	if err != nil {
		return "", errors.New("unable to create genesis block")
	}

	ledgerAccountPayload.OwnerId = ownerId
	ledgerAccountPayload.AccountNumber = accountNumber
	ledgerAccountPayload.CurrentActiveBlockId = block.ID
	ledgerAccountPayload.Status = types.APPROVED
	ledgerAccountPayload.BlockCount = 1

	_, acctErr := accountService.ledgerAccountRepo.Create(*ledgerAccountPayload)

	if acctErr != nil {
		return "", err
	}

	return accountNumber, nil
}

func (accountService *AccountService) CreateInternalAccounts(ownerId string) []*types.InternalAccount {
	accounts := make([]*types.InternalAccount, 4)

	accountIdGenerator := utils.NewAccountIdGenerator(8)
	accountId, err := accountIdGenerator()

	if err != nil {
		panic("unable to generate account Id!!!")
	}

	for i := 1; i <= 4; i++ {
		acctNumber, _ := accountService.CreateLedgerAccount(fmt.Sprintf("A%s-%s", fmt.Sprint(i), accountId), fmt.Sprintf("A%s", fmt.Sprint(i)), func(claoi *CreateLedgerAccountOptionInput) {
			claoi.LedgerAccountInput.Label = fmt.Sprintf("A%s", fmt.Sprint(i))
		})

		if i == 1 {
			accountId = acctNumber
		}

		accounts[i-1] = &types.InternalAccount{
			AccountNumber: acctNumber,
			Label:         fmt.Sprintf("A%s", fmt.Sprint(i)),
			OwnerId:       ownerId,
		}
	}

	return accounts

}

func (accountService *AccountService) GetAccountPairs(mainAccountNumber string) *types.AccountPairs {
	mainAccount, err := accountService.ledgerAccountRepo.FindByAccountNumber(mainAccountNumber)

	if err != nil {
		panic("invalid main account number")
	}

	a2CoaAccount, err2 := accountService.chartOfAccountService.FindByName("A2-" + mainAccount.AccountNumber)
	a3CoaAccount, err3 := accountService.chartOfAccountService.FindByName("A3-" + mainAccount.AccountNumber)
	a4CoaAccount, err4 := accountService.chartOfAccountService.FindByName("A4-" + mainAccount.AccountNumber)

	if err2 != nil || err3 != nil || err4 != nil {
		panic("incomplete account!!!")
	}

	a2Account, _ := accountService.ledgerAccountRepo.FindByAccountNumber(a2CoaAccount.AccountNumber)
	a3Account, _ := accountService.ledgerAccountRepo.FindByAccountNumber(a3CoaAccount.AccountNumber)
	a4Account, _ := accountService.ledgerAccountRepo.FindByAccountNumber(a4CoaAccount.AccountNumber)

	return &types.AccountPairs{
		A1: mainAccount.AccountNumber,
		A2: a2Account.AccountNumber,
		A3: a3Account.AccountNumber,
		A4: a4Account.AccountNumber,
	}
}

func (accountService *AccountService) AccountBalance(accountNumber string) int {
	account, err := accountService.ledgerAccountRepo.FindByAccountNumber(accountNumber)

	if err != nil {
		panic("invalid account")
	}

	block, blockErr := accountService.accountBlockRepo.FindById(account.CurrentActiveBlockID)

	if blockErr != nil {
		panic("invalid block")
	}

	entries, _ := accountService.journalEntryRepo.FindAllByBlockId(block.ID)

	balance := 0

	for _, entry := range entries {
		if entry.Type == string(types.DEBIT) {
			balance += int(entry.Amount)
		} else {
			balance += int(entry.Amount)
		}
	}

	return balance
}

func (accountService *AccountService) GetAccount(accountNumber string) *types.AccountRepresentation {
	account, err := accountService.ledgerAccountRepo.FindByAccountNumber(accountNumber)

	if err != nil {
		panic("invalid account number")
	}

	balance := accountService.AccountBalance(accountNumber)

	return &types.AccountRepresentation{
		AccountNumber:        accountNumber,
		OwnerId:              account.OwnerID,
		Book:                 types.LedgerBook(account.Book),
		CurrentActiveBlockId: account.CurrentActiveBlockID,
		Status:               types.LedgerAccountStatus(account.Status),
		Label:                account.Label,
		BlockCount:           int(account.BlockCount),
		Particular:           account.Particular,
		CreatedAt:            account.CreatedAt.String(),
		Balance:              balance,
	}
}
