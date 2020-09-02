package main

import (
	"fmt"
	"learngo/day1/accounts"
)

func main() {
	account := accounts.NewAccount("song")
	account.Deposit(10)
	fmt.Println(account)
}
