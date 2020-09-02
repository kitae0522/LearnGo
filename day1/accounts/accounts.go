package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (theAccount *Account) Deposit(amount int) {
	theAccount.balance += amount
}

// Withdraw x amount from your account
func (theAccount *Account) Withdraw(amount int) error {
	if theAccount.balance < amount {
		return errors.New("Can't Withdraw")
	}

	theAccount.balance -= amount
	return nil
}

// Balance of your Account
func (theAccount Account) Balance() int {
	return theAccount.balance
}

// ChangeOwner of the Account
func (theAccount *Account) changeOwner(newOwner string) {
	theAccount.owner = newOwner
}

// Owner of Account
func (theAccount Account) Owner() string {
	return theAccount.owner
}

func (theAccount Account) String() string {
	return fmt.Sprint(theAccount.Owner(), "'s Account\nHas: $", theAccount.Balance())
}
