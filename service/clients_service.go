package service

import (
	"go_homework_1/db"
	"go_homework_1/model"
)

func GetAllClients(dest *[]model.Client) error {
	err := db.Db.Select(dest, `SELECT * FROM clients`)
	return err
}

func GetClient(id int64, dest *model.Client) error {
	err := db.Db.Get(dest, `SELECT * FROM clients c WHERE c.id = $1`, id)
	return err
}

func CreateClient(c *model.Client) error {
	rows, err := db.Db.NamedQuery(`INSERT INTO clients (name, email, phone) VALUES (:name, :email, :phone) RETURNING id`, map[string]interface{}{
		"name":  c.Name,
		"email": c.Email,
		"phone": c.Phone,
	})
	if rows.Next() {
		rows.Scan(&c.ID)
	}
	return err
}

func UpdateClient(c *model.Client) error {
	_, err := db.Db.NamedQuery(`UPDATE clients SET name = :name, email = :email, phone = :phone WHERE id = :id`, map[string]interface{}{
		"id": c.ID,
		"name":  c.Name,
		"email": c.Email,
		"phone": c.Phone,
	})
	return err
}

func DeleteClient(id int64) error {
	_, err := db.Db.NamedExec(`DELETE FROM clients c WHERE c.id = :id`, map[string]interface{}{
		"id": id,
	})
	return err
}
