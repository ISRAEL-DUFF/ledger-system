package types

type LedgerAccountType string

const (
	ASSET     LedgerAccountType = "asset"
	LIABILITY LedgerAccountType = "liability"
	EQUITY    LedgerAccountType = "equity"
	REVENUE   LedgerAccountType = "revenue"
	EXPENSIS  LedgerAccountType = "expensis"
)

type ChartOfAccount struct {
	Id            string
	AccountNumber string
	CreatedAt     string
}
