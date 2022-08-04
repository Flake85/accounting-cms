package router

import (
	"net/http"
	handlers "server/handlers"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetClients",
		"GET",
		"/clients",
		handlers.GetClients,
	},
	Route{
		"CreateClients",
		"POST",
		"/clients",
		handlers.CreateClient,
	},
	Route{
		"GetClient",
		"GET",
		"/clients/{id}",
		handlers.GetClient,
	},
	Route{
		"UpdateClient",
		"PUT",
		"/clients/{id}",
		handlers.UpdateClient,
	},
	Route{
		"DeleteClient",
		"DELETE",
		"/clients/{id}",
		handlers.DeleteClient,
	},
}