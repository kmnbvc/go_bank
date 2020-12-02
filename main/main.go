package main

import (
	"go_homework_1/db"
	"go_homework_1/model"
	"go_homework_1/service"
	"log"
	"go_homework_1/server"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatalln(err)
		return
	}
	defer db.Close()

	c := model.NewClient("client1", "mail@mail.com", "phone712313")
	service.CreateClient(c)

	a1 := model.NewAccount(100, c)
	a2 := model.NewAccount(200, c)
	service.CreateAccount(a1)
	service.CreateAccount(a2)

	service.Add(15, a1.ID)
	service.Add(25, a2.ID)
	service.Transfer(5, a1.ID, a2.ID)
	service.Withdraw(5, a2.ID)

	server.Start()
}
