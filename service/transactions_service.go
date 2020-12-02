package service

import (
	"github.com/jmoiron/sqlx"
	"go_homework_1/model"
	"log"
)

func CreateTransaction(t *model.Transaction, tx *sqlx.Tx) error {
	rows, err := tx.NamedQuery(`INSERT INTO transactions (message, error) VALUES (:msg, :err) RETURNING id`, map[string]interface{}{
		"msg": t.Message, 
		"err": t.ErrorMsg(),
	})
	if err != nil {
		return err
	}
	if rows.Next() {
		rows.Scan(&t.ID)
	}
	rows.Close()
	for _, mc := range t.Details {
		if err = createMoneyChange(mc, tx); err != nil {
			return err
		}
	}
	return nil
}

func createMoneyChange(mc *model.MoneyChange, tx *sqlx.Tx) error {
	_, err := tx.Exec(`INSERT INTO money_changes (operation, account_id, amount, transaction_id) VALUES ($1, $2, $3, $4)`, 
			mc.Operation, mc.Account.ID, mc.Balance, mc.Tx.ID)
	return err
}
