package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Can't withdraw you are poor")

// Account Struct
type account struct {
	owner   string
	balance int
}

// New Account creates Account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// (a account) 부분은 리시버라 한다
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a account) Balance() int {
	return a.balance
}

// Withdraw x amount in your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of your account
func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.owner,
		"'s account.\nHas: ",
		a.balance)
}
