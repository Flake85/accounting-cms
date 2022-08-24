package router

import (
	"net/http"
	"server/middleware"

	"github.com/gorilla/mux"
)

func NewRouter () *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		for _, route := range route {
		var handler http.HandlerFunc
		handler = route.HandlerFunc
		methods := []string{route.Method}
		if route.Method == "PUT" || route.Method == "DELETE" {
			methods = append(methods, "OPTIONS")
		}
		router.
				Methods(methods...).
				Path(route.Pattern).
				Name(route.Name).
				Handler(middleware.AddDefaultCORSHeaders(handler))
		}
	}
	return router
}