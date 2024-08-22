package rest

import (
	"fmt"
	"net/http"
)

type Router struct {
	Mux http.ServeMux
}

type Mapper struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRouter(mappers []Mapper) *Router {
	r := &Router{}

	for _, m := range mappers {
		pattern := fmt.Sprintf("%s %s", m.Method, m.Path)
		r.Mux.HandleFunc(pattern, m.Handler)
	}
	return r
}
