package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func Init() error {
	var err error
	Db, err = sqlx.Connect("postgres", "dbname=go_homework_1 host=localhost port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	err := Db.Close()
	Db = nil
	return err
}
