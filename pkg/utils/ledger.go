package utils

import (
	"crypto/rand"

	"github.com/aidarkhanov/nanoid"
)

func CustomAlphabet(alphabet string, size int) func() (string, error) {
	// generateBytesBuffer returns random bytes buffer
	generateBytesBuffer := func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := rand.Read(buffer); err != nil {
			return nil, err
		}

		return buffer, nil
	}

	generatorFn := func() (string, error) {
		id, err := nanoid.Format(generateBytesBuffer, alphabet, size)

		return id, err
	}

	return generatorFn
}

func NewAccountIdGenerator(size int) func() (string, error) {
	return CustomAlphabet("0123456789", size)
}

func NewApiKeyGenerator() func() (string, error) {
	return CustomAlphabet("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 64)
}

func DeleteArrayItem[T any](index int32, array []T) ([]T, bool) {
	arrayLen := len(array)

	if index < 0 || index >= int32(arrayLen) {
		return array, false
	}

	array = append(array[:index], array[index+1:]...)

	return array, true
}

func GetArrayItemIndex[T comparable](item T, arr []T) (int, bool) {
	for i, v := range arr {
		if item == v {
			return i, true
		}
	}

	return -1, false
}

func GenerateDefaultWalletRules() []map[string]interface{} {
	fundingRule := map[string]interface{}{
		"event": "fund",
		"input": []string{
			"accountNumber",
			"memo",
		},
		"rule": map[string]interface{}{
			"credit": "A1",
			"debit":  "A2",
		},
	}

	withdrawRule := map[string]interface{}{
		"event": "withdraw",
		"input": []string{
			"accountNumber",
			"memo",
		},
		"rule": map[string]interface{}{
			"credit": "A3",
			"debit":  "A1",
		},
	}

	fundTransferRule := map[string]interface{}{
		"event": "transfer",
		"input": []string{
			"fromAccountNumber",
			"memo",
		},
		"rule": map[string]interface{}{
			"credit": "A4",
			"debit":  "A1",
		},
		"emitRules": []map[string]interface{}{
			{
				"event": "transfer.receive",
				"to":    "toAccountNumber",
				"withInput": []string{
					"amount",
					"memo",
				},
			},
		},
	}

	fundTransferReceiveRule := map[string]interface{}{
		"event": "transfer.receive",
		"input": []string{
			"accountNumber",
			"memo",
		},
		"rule": map[string]interface{}{
			"credit": "A1",
			"debit":  "A4",
		},
	}

	defaultRules := []map[string]interface{}{
		fundingRule,
		withdrawRule,
		fundTransferRule,
		fundTransferReceiveRule,
	}

	return defaultRules
}
