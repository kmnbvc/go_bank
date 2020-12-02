package model

import (
	"log"
)

type Transaction struct {
	ID      int64
	Message string
	Details []*MoneyChange
	Err     error
}

func NewTransaction(msg string, err error) *Transaction {
	log.Println(msg)
	t := &Transaction{Message: msg, Err: err}
	return t
}

func (t *Transaction) ErrorMsg() string {
	if t.Err != nil {
		return t.Err.Error()
	}
	return ""
}

func (t *Transaction) addMoneyChange(mc *MoneyChange) {
	t.Details = append(t.Details, mc)
}
