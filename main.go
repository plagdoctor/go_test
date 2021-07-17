package main

import (
	"PlagDoctor/accounts"
	"fmt"
	"log"
)

//	/home/miumiu/go/src/
func main() {
	account := accounts.NewAccount("pepe")

	account.Deposit(100)
	fmt.Println(account.Owner())
	fmt.Println(account.CheckBalance())

	fmt.Println(account)
	err := account.Withdraw(150)
	//err := account.Withdraw(50)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.CheckBalance())
}
