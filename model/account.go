package model

import (
	"fmt"
)

type Account struct {
	ID int64
	Balance Balance
	Client *Client
	Closed bool
}

type Balance int64

func NewEmptyAccount() interface{} {
	return &Account{}
}

func NewAccount(balance Balance, client *Client) *Account {
	a := &Account{Balance: balance, Client: client}
	return a
}

func (b Balance) Positive() error {
	if b < 1 {
		return fmt.Errorf("error: amount should be positive")
	}
	return nil
}

func (a *Account) Available(amount Balance) error {
	if a.Balance < amount {
		return fmt.Errorf("error: current balance is less than specified amount")
	}
	return nil
}

func (a *Account) Add(amount Balance) error {
	if err := amount.Positive(); err != nil {
		return err
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount Balance) error {
	if err := amount.Positive(); err != nil {
		return err
	}

	if err := a.Available(amount); err != nil {
		return err
	}

	a.Balance -= amount
	return nil
}

func (a *Account) TransferTo(other *Account, amount Balance) error {
	if err := amount.Positive(); err != nil {
		return err
	}

	if err := a.Available(amount); err != nil {
		return err
	}

	a.Balance -= amount
	other.Balance += amount

	return nil
}



