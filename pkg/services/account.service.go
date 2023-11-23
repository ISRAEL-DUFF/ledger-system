package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/israel-duff/ledger-system/pkg/db/model"
	"github.com/israel-duff/ledger-system/pkg/db/repositories"
	"github.com/israel-duff/ledger-system/pkg/types"
	"github.com/israel-duff/ledger-system/pkg/types/datastructure"
	"github.com/israel-duff/ledger-system/pkg/utils"
)

type IAccountService interface {
	CreateLedgerAccount(name string, ownerId string, opts ...LedgerAccountOption) (string, error)
	AccountBalance(accountNumber string) int
	GetAccount(accountNumber string) *types.AccountRepresentation
	InitializeNewBlock(accountNumber string, dbQueryTx types.IDBTransaction) *types.AccountRepresentation
	ExtractTransactionEntries(input types.PostTransactionInput, entryList [][]types.TransactionInputEntry) ([][]types.TransactionInputEntry, error)
	CreateWalletType(ownerId string, name string) (*model.WalletType, error)
	CreateWallet(ownerId string, walletTypeId string) *Wallet
	ListWalletTypes(ownerId string) ([]*model.WalletType, error)
	GetWalletByAccountNumber(accountNumber string) (*Wallet, error)
	ListUserWallets(ownerId string) ([]*Wallet, error)
	UpdateWalletType(typeId string, walletTypeData map[string]interface{}) error
	WalletStatus(mainAccountNumber string) *types.AccountStatus
}

type AccountService struct {
	chartOfAccountService ICoaService
	accountBlockRepo      repositories.IAccountBlockRepository
	ledgerAccountRepo     repositories.ILedgerAccount
	journalEntryRepo      repositories.IJournalEntryRepository
	walletRepo            repositories.IWalletRepository
	walletTypeRepo        repositories.IWalletTypeRepository
}

type Wallet struct {
	accounts   map[string]model.LedgerAccount
	walletType WalletTypeStructure
	balance    float32
}

type WalletTypeStructure struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Events []types.WalletRuleType `json:"events"`
}

type CreateLedgerAccountOptionInput struct {
	CoaAccountInput    *types.CreateCoaAccount
	AccountBlockInput  *types.CreateAccountBlock
	LedgerAccountInput *types.CreateLedgerAccount
}

type LedgerAccountOption func(*CreateLedgerAccountOptionInput)

func NewAccountService(chartOfAccountService ICoaService,
	accountBlock repositories.IAccountBlockRepository,
	ledgerAccount repositories.ILedgerAccount,
	journalEntry repositories.IJournalEntryRepository,
	walletRepo repositories.IWalletRepository,
	walletTypeRepo repositories.IWalletTypeRepository,
) *AccountService {
	return &AccountService{
		chartOfAccountService: chartOfAccountService,
		accountBlockRepo:      accountBlock,
		ledgerAccountRepo:     ledgerAccount,
		journalEntryRepo:      journalEntry,
		walletRepo:            walletRepo,
		walletTypeRepo:        walletTypeRepo,
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

	ledgerAccount, acctErr := accountService.ledgerAccountRepo.Create(*ledgerAccountPayload)

	if acctErr != nil {
		return "", acctErr
	}

	block.AccountID = ledgerAccount.ID

	err = accountService.accountBlockRepo.Update(block)

	if err != nil {
		return "", err
	}

	return accountNumber, nil
}

func (accountService *AccountService) CreateWallet(ownerId string, walletTypeId string) *Wallet {
	walletType := accountService.walletTypeRepo.GetWalletRulesByTypeId(walletTypeId)

	accountIdGenerator := utils.NewAccountIdGenerator(8)
	accountId, err := accountIdGenerator()

	if err != nil {
		panic("unable to generate account Id!!!")
	}

	accountIndex := 1

	// fmt.Sprintf("A%s", fmt.Sprint(accountIndex))
	acctNumber, _ := accountService.CreateLedgerAccount(fmt.Sprintf("A%s-%s", fmt.Sprint(accountIndex), accountId), ownerId, func(claoi *CreateLedgerAccountOptionInput) {
		claoi.LedgerAccountInput.Label = fmt.Sprintf("A%s", fmt.Sprint(accountIndex))
	})

	account, aErr := accountService.ledgerAccountRepo.FindByAccountNumber(acctNumber)

	if aErr != nil {
		panic("unable to retrieve account")
	}

	_, wErr := accountService.walletRepo.Create(types.CreateWallet{
		AccountNumber: acctNumber,
		LedgerAccounts: []string{
			acctNumber,
		},
		Name:       acctNumber,
		OwnerId:    ownerId,
		WalletType: walletTypeId,
	})

	if wErr != nil {
		panic("unable to create wallet")
	}

	wallet := &Wallet{
		accounts: map[string]model.LedgerAccount{
			"A1": *account,
		},
		walletType: WalletTypeStructure{
			Name:   walletType.Name,
			ID:     walletTypeId,
			Events: walletType.Rules,
		},
	}

	return wallet
}

func (accountService *AccountService) GetWalletByAccountNumber(accountNumber string) (*Wallet, error) {
	wallet, err := accountService.walletRepo.FindByAccountNumber(accountNumber)

	if err != nil {
		panic("invalid wallet account number")
	}

	walletType := accountService.walletTypeRepo.GetWalletRulesByTypeId(wallet.Type)
	jsonStr := wallet.LedgerAccounts

	var ledgerAccounts []string

	err = json.Unmarshal([]byte(jsonStr), &ledgerAccounts)

	if err != nil {
		return nil, err
	}

	accounts, aErr := accountService.ledgerAccountRepo.FindAllByAccountNumbers(ledgerAccounts)

	if aErr != nil {
		return nil, aErr
	}

	balance := accountService.AccountBalance(accountNumber)

	walletData := &Wallet{
		accounts: map[string]model.LedgerAccount{},
		// walletType: walletType.Rules,
		walletType: WalletTypeStructure{
			Name:   walletType.Name,
			ID:     wallet.Type,
			Events: walletType.Rules,
		},
		balance: float32(balance),
	}
	setOfAccountLabels := walletType.AccountLabels()
	createdAccountLabels := datastructure.NewSet[string]()

	for _, account := range accounts {
		walletData.accounts[account.Label] = *account
		createdAccountLabels.Add(account.Label)
	}

	uncreatedAccountLabels := setOfAccountLabels.Difference(createdAccountLabels).Values()
	generatedAccountNumbers := []string{}

	for _, acctLable := range uncreatedAccountLabels {
		// TODO: Create and attach the missing account to this wallet
		acctName := fmt.Sprintf("A%s-%s", acctLable, wallet.AccountNumber)
		acctNumber, _ := accountService.CreateLedgerAccount(acctName, wallet.OwnerID, func(claoi *CreateLedgerAccountOptionInput) {
			claoi.LedgerAccountInput.Label = acctLable
		})

		newAccount, err := accountService.ledgerAccountRepo.FindByAccountNumber(acctNumber)

		if err != nil {
			return nil, err
		}

		walletData.accounts[newAccount.Label] = *newAccount
		generatedAccountNumbers = append(generatedAccountNumbers, acctNumber)
	}

	if len(generatedAccountNumbers) > 0 {
		err := accountService.walletRepo.AddLedgerAccounts(accountNumber, generatedAccountNumbers)

		if err != nil {
			return nil, err
		}
	}

	return walletData, nil
}

func (accountService *AccountService) ListUserWallets(ownerId string) ([]*Wallet, error) {
	wallets, err := accountService.walletRepo.FindAllByOwnerId(ownerId)

	if err != nil {
		panic("invalid wallet account number")
	}

	walletList := []*Wallet{}

	for _, wallet := range wallets {
		walletType := accountService.walletTypeRepo.GetWalletRulesByTypeId(wallet.Type)
		jsonStr := wallet.LedgerAccounts

		var ledgerAccounts []string

		err = json.Unmarshal([]byte(jsonStr), &ledgerAccounts)

		if err != nil {
			return nil, err
		}

		accounts, aErr := accountService.ledgerAccountRepo.FindAllByAccountNumbers(ledgerAccounts)

		if aErr != nil {
			return nil, aErr
		}

		balance := accountService.AccountBalance(wallet.AccountNumber)

		walletData := &Wallet{
			accounts: map[string]model.LedgerAccount{},
			walletType: WalletTypeStructure{
				Name:   walletType.Name,
				ID:     wallet.Type,
				Events: walletType.Rules,
			},
			balance: float32(balance),
		}
		setOfAccountLabels := walletType.AccountLabels()
		createdAccountLabels := datastructure.NewSet[string]()

		for _, account := range accounts {
			walletData.accounts[account.Label] = *account
			createdAccountLabels.Add(account.Label)
		}

		uncreatedAccountLabels := setOfAccountLabels.Difference(createdAccountLabels).Values()
		generatedAccountNumbers := []string{}

		for _, acctLable := range uncreatedAccountLabels {
			// TODO: Create and attach the missing account to this wallet
			acctName := fmt.Sprintf("A%s-%s", acctLable, wallet.AccountNumber)
			acctNumber, _ := accountService.CreateLedgerAccount(acctName, wallet.OwnerID, func(claoi *CreateLedgerAccountOptionInput) {
				claoi.LedgerAccountInput.Label = acctLable
			})

			newAccount, err := accountService.ledgerAccountRepo.FindByAccountNumber(acctNumber)

			if err != nil {
				return nil, err
			}

			walletData.accounts[newAccount.Label] = *newAccount
			generatedAccountNumbers = append(generatedAccountNumbers, acctNumber)
		}

		if len(generatedAccountNumbers) > 0 {
			err := accountService.walletRepo.AddLedgerAccounts(wallet.AccountNumber, generatedAccountNumbers)

			if err != nil {
				return nil, err
			}
		}

		walletList = append(walletList, walletData)
	}

	return walletList, nil
}

func (accountService *AccountService) CreateWalletType(ownerId string, name string) (*model.WalletType, error) {
	defaultRules := utils.GenerateDefaultWalletRules()
	typeRule := map[string]interface{}{
		"name":  name,
		"rules": defaultRules,
	}

	walletType, err := accountService.walletTypeRepo.Create(types.CreateWalletType{
		Name:    name,
		Rules:   typeRule,
		OwnerId: ownerId,
	})

	if err != nil {
		return nil, err
	}

	return walletType, nil
}

func (accountService *AccountService) UpdateWalletType(typeId string, walletTypeData map[string]interface{}) error {
	err := accountService.walletTypeRepo.UpdateWalletType(typeId, walletTypeData)

	if err != nil {
		return err
	}

	return nil
}

func (accountService *AccountService) ListWalletTypes(ownerId string) ([]*model.WalletType, error) {
	wt, err := accountService.walletTypeRepo.FindByOwnerId(ownerId)

	if err != nil {
		return nil, err
	}

	return wt, nil
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
			balance -= int(entry.Amount)
		} else {
			balance += int(entry.Amount)
		}
	}

	// time.Sleep(time.Millisecond * 200)

	return balance
}

func (accountService *AccountService) GetAccount(accountNumber string) *types.AccountRepresentation {
	account, err := accountService.ledgerAccountRepo.FindByAccountNumber(accountNumber)

	if err != nil {
		panic("invalid account number")
	}

	balance := accountService.AccountBalance(accountNumber)

	return &types.AccountRepresentation{
		ID:                   account.ID,
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

func (accountService *AccountService) InitializeNewBlock(accountNumber string, dbQueryTx types.IDBTransaction) *types.AccountRepresentation {
	fmt.Println("Initializing a new block for " + accountNumber)
	accountRepo := accountService.ledgerAccountRepo.WithTransaction(dbQueryTx)
	account, err := accountRepo.FindByAccountNumber(accountNumber)

	if err != nil {
		panic("invalid account number")
	}

	balance := accountService.AccountBalance(accountNumber)

	accountBlockRepo := accountService.accountBlockRepo.WithTransaction(dbQueryTx)
	block, blckErr := accountBlockRepo.FindById(account.CurrentActiveBlockID)

	if blckErr != nil {
		panic("invalid block")
	}

	if block.TransactionsCount < block.BlockSize {
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

	block.Status = string(types.CLOSE)
	block.IsCurrentBlock = false

	accountBlockRepo.Update(block)

	newBlock, createErr := accountBlockRepo.Create(types.CreateAccountBlock{
		IsCurrentBlock:    true,
		Status:            types.OPEN,
		TransactionsCount: 0,
		BlockSize:         50,
		AccountId:         account.ID,
	})

	if createErr != nil {
		panic("unable to create new block")
	}

	account.CurrentActiveBlockID = newBlock.ID
	account.BlockCount += 1

	accountRepo.Update(account)

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

func (accountService *AccountService) WalletStatus(mainAccountNumber string) *types.AccountStatus {
	wallet, err := accountService.GetWalletByAccountNumber(mainAccountNumber)

	if err != nil {
		panic(err)
	}

	netBalance := 0
	accounts := make([]types.AccountStatusInfo, 0)

	for _, account := range wallet.accounts {
		balance := accountService.AccountBalance(account.AccountNumber)

		accounts = append(accounts, types.AccountStatusInfo{
			AccountNumber: account.AccountNumber,
			Balance:       balance,
			Type:          account.Label,
		})
		netBalance += balance
	}

	return &types.AccountStatus{
		Balanced:  netBalance == 0,
		NetAmount: netBalance,
		Accounts:  accounts[:],
	}
}

func (accountService *AccountService) ExtractTransactionEntries(input types.PostTransactionInput, entryList [][]types.TransactionInputEntry) ([][]types.TransactionInputEntry, error) {
	amountData, amtExists := input.MetaData["amount"]
	memoData := input.MetaData["memo"]

	if !amtExists {
		return nil, errors.New("amount is missing in meta data")
	}

	amountFloat := amountData.(float64)
	amount := int(amountFloat)
	memo := memoData.(string)

	wallet, err := accountService.GetWalletByAccountNumber(input.AccountNumber)

	if err != nil {
		return nil, err
	}

	walletRuleType, typeExists := wallet.GetWalletRuleType(input.EventName)

	if !typeExists {
		return nil, errors.New("the supplied wallet does not have the post event")
	}

	postRule := walletRuleType.Rule
	debitAccountLabel := postRule.Debit
	creditAccountLabel := postRule.Credit
	debitAccount := wallet.accounts[debitAccountLabel]
	creditAccount := wallet.accounts[creditAccountLabel]

	txEntries := []types.TransactionInputEntry{
		{
			TransactionEntry: types.TransactionEntry{
				AccountNumber: creditAccount.AccountNumber,
				Amount:        amount,
				Type:          types.CREDIT,
			},
			Memo:    memo,
			OwnerId: creditAccount.OwnerID,
		},

		{
			TransactionEntry: types.TransactionEntry{
				AccountNumber: debitAccount.AccountNumber,
				Amount:        amount,
				Type:          types.DEBIT,
			},
			Memo:    memo,
			OwnerId: debitAccount.OwnerID,
		},
	}

	entryList = append(entryList, txEntries)

	if walletRuleType.EmitRules != nil {
		for _, emitRule := range walletRuleType.EmitRules {
			toAccountNum := input.MetaData[emitRule.To].(string)

			if toAccountNum == "" {
				return nil, errors.New("invalid to-account number")
			}

			l, er := accountService.ExtractTransactionEntries(types.PostTransactionInput{
				AccountNumber: toAccountNum,
				EventName:     emitRule.Event,
				MetaData:      input.MetaData,
			}, entryList)

			if er != nil {
				return nil, er
			}

			entryList = l
		}
	}

	return entryList, nil
}

// Wallet Methods

func (wallet *Wallet) GetWalletRuleType(eventName string) (*types.WalletRuleType, bool) {
	for _, evt := range wallet.walletType.Events {
		if evt.Event == eventName {
			return &evt, true
		}
	}

	return nil, false
}

func (wallet *Wallet) GetWalletType() WalletTypeStructure {
	return wallet.walletType
}

func (wallet *Wallet) GetAccounts() map[string]model.LedgerAccount {
	return wallet.accounts
}

func (wallet *Wallet) GetBalance() float32 {
	return wallet.balance
}
