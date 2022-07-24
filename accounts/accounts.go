package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount build func
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 100}
	return &account
}

// Deposit method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw method
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Nope. You are poor.")
	}
	a.balance -= amount
	return nil
}

// Balance checking method
func (a Account) Balance() int {
	return a.balance
}

// ChangeOwner method
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner checking method
func (a Account) Owner() string {
	return a.owner
}
