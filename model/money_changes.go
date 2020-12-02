package model

type MoneyChange struct {
	ID      int64
	Operation
	Balance
	Account *Account
	Tx      *Transaction
}

type Operation string

const (
	Add      Operation = "ADD"
	Withdraw Operation = "WITHDRAW"
)

func NewMoneyChange(op Operation, balance Balance, accountID int64, tx *Transaction) *MoneyChange {
	mc := &MoneyChange{Operation: op, Account: &Account{ID: accountID}, Balance: balance, Tx: tx}
	tx.addMoneyChange(mc)
	return mc
}


