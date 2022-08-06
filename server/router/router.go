package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter () *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		for _, route := range route {
		var handler http.HandlerFunc
		handler = route.HandlerFunc
		router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	return router
}