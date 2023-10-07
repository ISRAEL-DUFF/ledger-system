package types

type LedgerBook string
type LedgerAccountStatus string
type AccountBlockStatus string
type JournalType string

const (
	CASH_RECEIPT   LedgerBook = "cash_receipt"
	GENERAL_LEDGER LedgerBook = "general_ledger"
)

const (
	PENDING  LedgerAccountStatus = "pending"
	APPROVED LedgerAccountStatus = "approved"
)

const (
	OPEN  AccountBlockStatus = "open"
	CLOSE AccountBlockStatus = "close"
)

const (
	DEBIT  JournalType = "debit"
	CREDIT JournalType = "credit"
)

type CreateLedgerAccount struct {
	AccountNumber        string
	OwnerId              string
	Book                 LedgerBook
	CurrentActiveBlockId string
	Status               LedgerAccountStatus
	Label                string
	BlockCount           int
	Particular           string
}

type CreateAccountBlock struct {
	IsCurrentBlock    bool
	Status            AccountBlockStatus
	TransactionsCount int
	BlockSize         int
	AccountId         string
}

type CreateJournalEntry struct {
	Amount         int
	Type           JournalType
	BlockId        string
	TransactionId  string
	Name           string
	AccountNumber  string
	Memo           string
	OwnerId        string
	OrganizationId string
}
