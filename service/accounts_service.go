package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go_homework_1/db"
	"go_homework_1/model"
)

func GetAccounts(clientID int64, dest *[]model.Account) error {
	err := db.Db.Select(dest, `SELECT a.id, a.balance, a.closed FROM accounts a WHERE a.client_id = $1`, clientID)
	return err
}

func GetAccount(id int64, dest *model.Account) error {
	err := db.Db.Get(dest, `SELECT a.id, a.balance, a.closed FROM accounts a WHERE a.id = $1`, id)
	return err
}

func CreateAccount(a *model.Account) error {
	rows, err := db.Db.NamedQuery(`INSERT INTO accounts (balance, client_id) VALUES (:balance, :client_id) RETURNING id`, map[string]interface{}{
		"balance":   a.Balance,
		"client_id": a.Client.ID,
	})
	if rows.Next() {
		rows.Scan(&a.ID)
	}
	return err
}

func CloseAccount(id int64) error {
	_, err := db.Db.NamedQuery(`UPDATE accounts a SET closed = TRUE WHERE a.id = :id`, map[string]interface{}{
		"id": id,
	})
	return err
}

func CurrentBalance(accountID int64, b *model.Balance) error {
	err := db.Db.Get(&b, `SELECT a.balance FROM accounts a WHERE a.id = $1`, accountID)
	return err
}

func Add(amount model.Balance, id int64) error {
	var account model.Account
	
	tx, err := db.Db.Beginx()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Add: %d to the id(%d) account.", amount, id)
	mtx := model.NewTransaction(msg, nil)
	model.NewMoneyChange(model.Add, amount, id, mtx)
	
	if err := CreateTransaction(mtx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Get(&account, `SELECT id, balance FROM accounts WHERE id = $1 FOR UPDATE`, id); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if err := account.Add(amount); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if _, err := tx.Exec(`UPDATE accounts SET balance = $1 WHERE id = $2`, account.Balance, account.ID); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	return tx.Commit()
}

func Withdraw(amount model.Balance, id int64) error {
	var account model.Account

	tx, err := db.Db.Beginx()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Withdraw: %d from the id(%d) account.", amount, id)
	mtx := model.NewTransaction(msg, nil)
	model.NewMoneyChange(model.Withdraw, amount, id, mtx)
	
	if err := CreateTransaction(mtx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Get(&account, `SELECT id, balance FROM accounts WHERE id = $1 FOR UPDATE`, id); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if err := account.Withdraw(amount); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if _, err := tx.Exec(`UPDATE accounts SET balance = $1 WHERE id = $2`, account.Balance, account.ID); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	return tx.Commit()
}

func Transfer(amount model.Balance, fromID int64, toID int64) error {
	var account model.Account
	var other model.Account
	
	tx, err := db.Db.Beginx()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Transfer: %d from id(%d) to id(%d) account.", amount, fromID, toID)
	mtx := model.NewTransaction(msg, nil)
	model.NewMoneyChange(model.Add, amount, toID, mtx)
	model.NewMoneyChange(model.Withdraw, amount, fromID, mtx)
	
	if err := CreateTransaction(mtx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Get(&account, `SELECT id, balance FROM accounts WHERE id = $1 FOR UPDATE`, fromID); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if err := tx.Get(&other, `SELECT id, balance FROM accounts WHERE id = $1 FOR UPDATE`, toID); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if err := account.TransferTo(&other, amount); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if _, err := tx.Exec(`UPDATE accounts SET balance = $1 WHERE id = $2`, account.Balance, account.ID); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	if _, err := tx.Exec(`UPDATE accounts SET balance = $1 WHERE id = $2`, other.Balance, other.ID); err != nil {
		return updateTxError(tx, err, mtx.ID)
	}
	return tx.Commit()
}

func updateTxError(tx *sqlx.Tx, err error, txID int64) error {
	if _, errn := tx.Exec(`UPDATE transactions SET error = $1 WHERE id = $2`, err.Error(), txID); errn != nil {
		tx.Rollback()
		return errn
	}
	if errn := tx.Commit(); errn != nil {
		return errn
	}
	return err
}
