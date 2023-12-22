package main

import (
	"fmt"

	commonpkg "github.com/israel-duff/ledger-system/cmd/common-pkg"
	"github.com/israel-duff/ledger-system/pkg/config"
)

func main() {
	fmt.Println("Resetting Database...")
	databaseObject, err := config.NewDBConnection()

	truncateSql := `
	TRUNCATE TABLE public.account_blocks CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.block_meta CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.chart_of_accounts CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.journal_entries CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.ledger_accounts CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.ledger_transactions CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.users CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.wallet CONTINUE IDENTITY RESTRICT;
	TRUNCATE TABLE public.wallet_type CONTINUE IDENTITY RESTRICT;
	`

	if err != nil {
		panic(err)
	}

	dbTx := databaseObject.GetDBConnection().Exec(truncateSql)

	dbTx.Commit()

	fmt.Println("Done.")

	fmt.Println("Creating Wallets...")
	w1 := commonpkg.CreateUserWallet(commonpkg.Users[0])
	w2 := commonpkg.CreateUserWallet(commonpkg.Users[1])
	w3 := commonpkg.CreateUserWallet(commonpkg.Users[2])
	fmt.Println("Done.")

	fmt.Println(w1.Accounts["A1"].AccountNumber, w2.Accounts["A1"].AccountNumber, w3.Accounts["A1"].AccountNumber)

	// 11398884261 11811706910 12567176796
}
