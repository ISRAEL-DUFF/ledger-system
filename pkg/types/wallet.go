package types

type PostRule struct {
	Debit  string `json:"debit"`
	Credit string `json:"credit"`
}

type EmitRule struct {
	Event     string   `json:"event"`
	To        string   `json:"to"`
	WithInput []string `json:"withInput"`
}

type WalletRuleType struct {
	Event     string     `json:"event"`
	Input     []string   `json:"input"`
	Rule      PostRule   `json:"rule"`
	EmitRules []EmitRule `json:"emitRules"`
}

type PostTransactionInput struct {
	EventName     string                 `json:"eventName"`
	AccountNumber string                 `json:"accountNumber"`
	MetaData      map[string]interface{} `json:"metaData"`
}
