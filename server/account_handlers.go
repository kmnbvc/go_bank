package server

import (
	"errors"
	"net/http"
	"go_homework_1/model"
	"go_homework_1/service"
)

func accountsListHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	var accounts []model.Account
	cid := params["client_id"].(int64)
	err := service.GetAccounts(cid, &accounts)
	writeResponse(w, accounts, err)
}

func accountGetHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	id := params["id"].(int64)
	a := &model.Account{}
	err := service.GetAccount(id, a)
	writeResponse(w, a, err)
}

func accountCloseHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	id := params["id"].(int64)
	err := service.CloseAccount(id)
	writeResponse(w, "ok", err)
}

func accountCreateHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	a := params["account"].(*model.Account)
	err := service.CreateAccount(a)
	writeResponse(w, a, err)
}

func accountBalanceUpdateHandler(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
	bp := params["balance_operation"].(*balanceOperationParams)
	var err error
	switch bp.Operation {
	case "ADD":
		err = service.Add(bp.Amount, bp.AccountID)
	case "WITHDRAW":
		err = service.Withdraw(bp.Amount, bp.AccountID)
	case "TRANSFER":
		err = service.Transfer(bp.Amount, bp.AccountID, bp.TransferTargetID)
	default:
		err = errors.New("Error: operation not supported. Supported are ADD|WITHDRAW|TRANSFER.")
	}
	writeResponse(w, "ok", err)
}

type balanceOperationParams struct {
	AccountID        int64
	TransferTargetID int64
	Amount           model.Balance
	Operation        string
}

func newEmptyBalanceOpParams() interface{} {
	return &balanceOperationParams{}
}
