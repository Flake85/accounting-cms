package router

import (
	"net/http"
	"server/handlers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes(h) {
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
