package server

import (
	"net/http"
	"strconv"
	"log"
	"io"
	"encoding/json"
	"go_homework_1/model"
)

func Start() {
	r := newRouter()
	r.register("/clients", "GET", clientsListHandler, nil)
	r.register("/client", "GET", clientGetHandler, intParam("id"))
	r.register("/client", "DELETE", clientDeleteHandler, intParam("id"))
	r.register("/client", "POST", clientCreateHandler, jsonParam("client", model.NewEmptyClient))
	r.register("/client", "PUT", clientUpdateHandler, jsonParam("client", model.NewEmptyClient))
	r.register("/accounts", "GET", accountsListHandler, intParam("client_id"))
	r.register("/account", "GET", accountGetHandler, intParam("id"))
	r.register("/account", "DELETE", accountCloseHandler, intParam("id"))
	r.register("/account", "POST", accountCreateHandler, jsonParam("account", model.NewEmptyAccount))
	r.register("/account", "PUT", accountBalanceUpdateHandler, jsonParam("balance_operation", newEmptyBalanceOpParams))
	
	http.HandleFunc("/", r.handler)
	http.ListenAndServe(":8000", nil)
}

func handleError(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(500)
	io.WriteString(w, err.Error())
}

func intParam(name string) paramsExtractor {
	return func(r *http.Request) (map[string]interface{}, error) {
		value := r.FormValue(name)
		converted, err := strconv.Atoi(value)
		return map[string]interface{}{name: int64(converted)}, err
	}
}

func jsonParam(name string, targetSupplier func() interface{}) paramsExtractor {
	return func(r *http.Request) (map[string]interface{}, error) {
		target := targetSupplier()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(target)
		return map[string]interface{}{name: target}, err
	}
}

func writeResponse(w http.ResponseWriter, data interface{}, err error) {
	var serialized []byte
	if err != nil {
		handleError(err, w)
		return
	}
	if serialized, err = json.Marshal(data); err != nil {
        handleError(err, w)
        return
    }
	io.WriteString(w, string(serialized))
}
