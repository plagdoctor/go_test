package main

import (
	"PlagDoctor/accounts"
	"fmt"
)

//	/home/miumiu/go/src/
func main() {
	account := accounts.NewAccount("pepe")
	fmt.Println(account)
	account.Deposit(100)
	fmt.Println(account.CheckBalance())
}
