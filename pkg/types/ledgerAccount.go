package types

type LedgerBook string
type LedgerAccountStatus string
type AccountBlockStatus string
type JournalType string
type TransactionStatus string

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

const (
	TRANSACTION_PENDING  TransactionStatus = "pending"
	TRANSACTION_APPROVED TransactionStatus = "approved"
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

type CreateLedgerTransaction struct {
	Status LedgerAccountStatus
}

type InternalAccount struct {
	AccountNumber string
	OwnerId       string
	Label         string
}

type AccountPairs struct {
	A1 string
	A2 string
	A3 string
	A4 string
}

type AccountRepresentation struct {
	ID                   string
	AccountNumber        string
	OwnerId              string
	Book                 LedgerBook
	CurrentActiveBlockId string
	Status               LedgerAccountStatus
	Label                string
	BlockCount           int
	Particular           string
	CreatedAt            string
	Balance              int
}

type AccountStatusInfo struct {
	AccountNumber string
	Type          string
	Balance       int
}

type AccountStatus struct {
	Balanced bool
	Accounts []AccountStatusInfo
}

type TransactionEntry struct {
	AccountNumber string
	Amount        int
	Type          JournalType
}

type TransactionInputEntry struct {
	TransactionEntry
	Memo           string
	OwnerId        string
	OrganizationId string
}
type TransactionInput struct {
	Entries []TransactionInputEntry
}

type TransactionResponse struct {
	Entries       []TransactionEntry
	Status        TransactionStatus
	TransactionId string
}

type CreateBlockMetum struct {
	AccountId      string
	BlockTxLimit   int
	TransitionTxId string
	OpeningDate    string
	ClosingDate    string
}

type DependencyQueueItem struct {
	Tx           TransactionInput
	LockId       string
	Dependencies []string
}
