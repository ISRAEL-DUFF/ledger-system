package commonpkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const SERVER_URL string = "http://localhost:5050"

type WalletAccount struct {
	ID            string `json:"id"`
	AccountNumber string `json:"account_number"`
}

type WalletType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Wallet struct {
	Accounts map[string]WalletAccount `json:"accounts"`
	Type     WalletType               `json:"type"`
	Balance  int                      `json:"balance"`
}

type WalletList struct {
	Wallets []Wallet `json:"data"`
}

type CreateWalletTypeReponse struct {
	WalletType WalletType `json:"data"`
}

type CreateWalletResponse struct {
	Wallet Wallet `json:"data"`
}

func ListWallets(ownerID string) []Wallet {
	resp, err := http.Get(SERVER_URL + "/wallet/list/" + ownerID)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))
	// var jsonBody map[string]interface{}
	var walletList WalletList

	err = json.Unmarshal(body, &walletList)

	if err != nil {
		panic(err)
	}

	// return jsonBody
	return walletList.Wallets
}

func ListWalletTypes(ownerID string) map[string]interface{} {
	resp, err := http.Get(SERVER_URL + "/wallet/type/list/" + ownerID)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))
	var jsonBody map[string]interface{}

	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		panic(err)
	}

	return jsonBody
}

func CreateWalletType(ownerID string) WalletType {
	var payload map[string]interface{} = map[string]interface{}{
		"ownerId": ownerID,
		"name":    "Test Wallet",
	}

	jsonStr, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	// create an http request
	req, err := http.NewRequest("POST", SERVER_URL+"/wallet/type/create", bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respByteStr, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var dataResp CreateWalletTypeReponse

	err = json.Unmarshal(respByteStr, &dataResp)

	if err != nil {
		fmt.Println("String here json", string(respByteStr), len(respByteStr))
		panic(err)
	}

	return dataResp.WalletType
}

func CreateWallet(ownerID, walletTypeId string) Wallet {
	var payload map[string]interface{} = map[string]interface{}{
		"ownerId":    ownerID,
		"walletType": walletTypeId,
	}

	jsonStr, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	// create an http request
	req, err := http.NewRequest("POST", SERVER_URL+"/wallet/create", bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respByteStr, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var dataResp CreateWalletResponse

	err = json.Unmarshal(respByteStr, &dataResp)

	if err != nil {
		panic(err)
	}

	return dataResp.Wallet
}

func CreateUserWallet(ownerID string) Wallet {
	walletType := CreateWalletType(ownerID)
	wallet := CreateWallet(ownerID, walletType.ID)

	return wallet
}

func GetUserWallet(ownerID string) WalletAccount {
	walletsResp := ListWallets(ownerID)
	account := walletsResp[0].Accounts["A1"]

	return account
}

type PostTransactionOpts struct {
	PostType      string
	AccountNumber string
	MetaData      map[string]interface{}
}

func PostTransaction(opts PostTransactionOpts) map[string]interface{} {
	fmt.Println("posting transaction", opts)
	var body map[string]interface{} = map[string]interface{}{
		"eventName":     opts.PostType,
		"accountNumber": opts.AccountNumber,
		"metaData":      opts.MetaData,
	}

	jsonStr, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", SERVER_URL+"/wallet/transaction/post", bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	jsonRespBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonResp map[string]interface{}

	err = json.Unmarshal(jsonRespBytes, &jsonResp)

	if err != nil {
		panic(err)
	}

	fmt.Println(jsonResp)

	return jsonResp
}

func FundWallet(accountNumber string, amount int32) map[string]interface{} {
	return PostTransaction(PostTransactionOpts{
		PostType:      "fund",
		AccountNumber: accountNumber,
		MetaData: map[string]interface{}{
			"amount": amount,
			"memo":   "Fund wallet",
		},
	})
}

func Withdraw(accountNumber string, amount int32) map[string]interface{} {
	return PostTransaction(PostTransactionOpts{
		PostType:      "withdraw",
		AccountNumber: accountNumber,
		MetaData: map[string]interface{}{
			"amount": amount,
			"memo":   "Withdrawal from wallet",
		},
	})
}

func FundsTransfer(accountNumber1, accountNumber2 string, amount int32) map[string]interface{} {
	return PostTransaction(PostTransactionOpts{
		PostType:      "transfer",
		AccountNumber: accountNumber1,
		MetaData: map[string]interface{}{
			"memo":            "Test fundstrasfer",
			"amount":          amount,
			"toAccountNumber": accountNumber2,
		},
	})
}

func RandNumberBetween(max, min int) int {
	return rand.Int()%max + min
}

func RandomFund(accountNumber string, frequency int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= frequency; i++ {
		FundWallet(accountNumber, int32(RandNumberBetween(10, 50)))

		time.Sleep(time.Duration(100))
	}
}

func RandomWithDraw(accountNumber string, frequency int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= frequency; i++ {
		Withdraw(accountNumber, int32(RandNumberBetween(10, 50)))

		time.Sleep(time.Duration(100))
	}
}

func RandomFundsTransfer(accountNumbers []string, frequency int, wg *sync.WaitGroup) {
	defer wg.Done()
	// n := len(accountNumbers)

	for i := 0; i <= frequency; i++ {
		FundsTransfer(accountNumbers[0], accountNumbers[1], int32(RandNumberBetween(10, 50)))

		time.Sleep(time.Duration(100))
	}
}

var Users []string = []string{
	"b95467cb-e488-4310-b81e-33b3a96b3d2b",
	"3eeefe91-0494-43b1-9d1b-6bc1106c57cd",
	"42d6371d-e5c4-496c-be4c-5945f1e8ce32",
}
