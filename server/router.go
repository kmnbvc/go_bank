package server

import (
	"net/http"
	"errors"
)

type route struct {
	path string
	method string
	target func(http.ResponseWriter, *http.Request, map[string]interface{})
	paramsExtractor
}

type router struct {
	routes []*route
}

type paramsExtractor func(*http.Request) (map[string]interface{}, error)

func (rr *router) register(path string, method string, target func(http.ResponseWriter, *http.Request, map[string]interface{}), pe paramsExtractor) {
	rr.routes = append(rr.routes, &route{path, method, target, pe})
}

func (rr *router) handler(w http.ResponseWriter, r *http.Request) {
	var matched []*route
	var params map[string]interface{}
	var err error
	for _, rt := range rr.routes {
		if r.Method == rt.method && r.URL.Path == rt.path {
			matched = append(matched, rt)
		}
	}
	if len(matched) == 0 {
		http.NotFound(w, r)
		return
	}
	if len(matched) > 1 {
		handleError(errors.New("more than one handler found for the specified path"), w)
		return
	}
	rt := matched[0] 
	if rt.paramsExtractor != nil {
		params, err = rt.paramsExtractor(r)
	}
	if err != nil {
		handleError(err, w)
		return
	}
	rt.target(w, r, params)
}

func newRouter() *router {
	return &router{}
}

