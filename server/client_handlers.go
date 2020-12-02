package server

import (
	"net/http"
	"go_homework_1/model"
	"go_homework_1/service"
)

func clientsListHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	var clients []model.Client
	err := service.GetAllClients(&clients)
	writeResponse(w, clients, err)
}

func clientGetHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	id := params["id"].(int64)
	c := &model.Client{}
	err := service.GetClient(id, c)
	writeResponse(w, c, err)
}

func clientDeleteHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	id := params["id"].(int64)
	err := service.DeleteClient(id)
	writeResponse(w, "ok", err)
}

func clientCreateHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	c := params["client"].(*model.Client)
	err := service.CreateClient(c)
	writeResponse(w, c, err)
}

func clientUpdateHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	c := params["client"].(*model.Client)
	err := service.UpdateClient(c)
	writeResponse(w, c, err)
}
