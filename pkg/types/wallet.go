package types

import (
	"errors"
	"strings"
)

type PostRule struct {
	Debit  string `json:"debit" validate:"required"`
	Credit string `json:"credit" validate:"required"`
}

type EmitRule struct {
	Event     string   `json:"event" validate:"required"`
	To        string   `json:"to" validate:"required"`
	WithInput []string `json:"withInput"`
}

type WalletRuleType struct {
	Event     string     `json:"event" validate:"required,min=2"`
	Input     []string   `json:"input" validate:"required"`
	Rule      PostRule   `json:"rule" validate:"required"`
	EmitRules []EmitRule `json:"emitRules" validate:"omitempty"`
}

type PostTransactionInput struct {
	EventName     string                 `json:"eventName" validate:"required"`
	AccountNumber string                 `json:"accountNumber" validate:"required"`
	MetaData      map[string]interface{} `json:"metaData" validate:"required"`
}

// custom validations
func (postRule PostRule) Validate() error {
	if postRule.Credit == "" || postRule.Debit == "" {
		return errors.New("empty character not allowed")
	}

	if strings.Contains(postRule.Credit, " ") || strings.Contains(postRule.Debit, " ") {
		return errors.New("space not allowed in account names")
	}

	if len(postRule.Credit) < 2 || len(postRule.Debit) < 2 {
		return errors.New("account names must be at least 2")
	}

	return nil
}

func (emitRule EmitRule) Validate() error {
	if emitRule.Event == "" {
		return errors.New("empty character not allowed in event name")
	}

	if strings.Contains(emitRule.Event, " ") {
		return errors.New("space not allowed in event names")
	}

	if len(emitRule.Event) < 3 {
		return errors.New("event names must be at least 3")
	}

	// _, err := regexp.MatchString("[\\d]", emitRule.To)
	// if err != nil {
	// 	return errors.New("invalid account number")
	// }

	if emitRule.To == "" {
		return errors.New("empty character not allowed in to-account")
	}

	if strings.Contains(emitRule.To, " ") {
		return errors.New("space not allowed in to-account")
	}

	return nil
}

func (walletRule WalletRuleType) Validate() error {
	if walletRule.Event == "" {
		return errors.New("empty character not allowed in event name")
	}

	if strings.Contains(walletRule.Event, " ") {
		return errors.New("space not allowed in event names")
	}

	if len(walletRule.Event) < 3 {
		return errors.New("event names must be at least 3")
	}

	err := walletRule.Rule.Validate()

	if err != nil {
		return err
	}

	for _, emitRule := range walletRule.EmitRules {
		err := emitRule.Validate()

		if err != nil {
			return err
		}
	}

	return nil
}
