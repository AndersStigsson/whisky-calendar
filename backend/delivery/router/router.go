package router

import "net/http"

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
	ParsePathVariable(r *http.Request, name string) (interface{}, error)
	GetBody(r *http.Request) ([]byte, error)
}
