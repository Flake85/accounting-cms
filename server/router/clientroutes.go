package router

import "server/handlers"

var clientRoutes = Routes{
	Route{
		"GetClients",
		"GET",
		"/client",
		handlers.GetClients,
	},
	Route{
		"CreateClients",
		"POST",
		"/client",
		handlers.CreateClient,
	},
	Route{
		"GetClient",
		"GET",
		"/client/{id}",
		handlers.GetClient,
	},
	Route{
		"UpdateClient",
		"PUT",
		"/client/{id}",
		handlers.UpdateClient,
	},
	Route{
		"DeleteClient",
		"DELETE",
		"/client/{id}",
		handlers.DeleteClient,
	},
}