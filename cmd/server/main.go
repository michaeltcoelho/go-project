package main

import (
	"fmt"

	"maisdindin/src/domain/accounts"
)

func main() {
	accountID := accounts.AccountID("blah")
	fmt.Println(accountID)
	fmt.Println("Main.. worked")
}
