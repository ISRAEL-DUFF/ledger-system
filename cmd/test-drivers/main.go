package main

import (
	"fmt"
	"sync"

	commonpkg "github.com/israel-duff/ledger-system/cmd/common-pkg"
	// "github.com/israel-duff/ledger-system/cmd/commonpkg"
)

func main() {
	fmt.Println("Hello new driver")

	user1 := commonpkg.Users[0]
	user2 := commonpkg.Users[1]
	user3 := commonpkg.Users[2]

	var wg sync.WaitGroup
	wg.Add(6)

	user1Account := commonpkg.GetUserWallet(user1)
	user2Account := commonpkg.GetUserWallet(user2)
	user3Account := commonpkg.GetUserWallet(user3)
	frequency := 1000

	go commonpkg.RandomFund(user1Account.AccountNumber, frequency, &wg)
	go commonpkg.RandomFund(user2Account.AccountNumber, frequency, &wg)
	go commonpkg.RandomFund(user3Account.AccountNumber, frequency, &wg)

	go commonpkg.RandomWithDraw(user1Account.AccountNumber, frequency, &wg)
	go commonpkg.RandomWithDraw(user2Account.AccountNumber, frequency, &wg)
	go commonpkg.RandomWithDraw(user3Account.AccountNumber, frequency, &wg)

	wg.Wait()
}
