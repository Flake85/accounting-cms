package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter () *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.HandlerFunc
		handler = route.HandlerFunc
		router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
	}
	return router
}