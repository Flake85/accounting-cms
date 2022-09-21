package router

import "server/handlers"

func clientRoutes(handlers *handlers.Handler) Routes {
	return Routes{
		Route{
			Name: "GetClients",
			Method: "GET",
			Pattern: "/client",
			HandlerFunc: handlers.GetClients,
		},
		Route{
			Name: "CreateClients",
			Method: "POST",
			Pattern: "/client",
			HandlerFunc: handlers.CreateClient,
		},
		Route{
			Name: "GetClient",
			Method: "GET",
			Pattern: "/client/{id}",
			HandlerFunc: handlers.GetClient,
		},
		Route{
			Name: "UpdateClient",
			Method: "PUT",
			Pattern: "/client/{id}",
			HandlerFunc: handlers.UpdateClient,
		},
		Route{
			Name: "DeleteClient",
			Method: "DELETE",
			Pattern: "/client/{id}",
			HandlerFunc: handlers.DeleteClient,
		},
		Route{
			Name: "GetDeletedClients",
			Method: "GET",
			Pattern: "/client_deleted",
			HandlerFunc: handlers.GetDeletedClients,
		},
		Route{
			Name: "GetDeletedClient",
			Method: "GET",
			Pattern: "/client_deleted/{id}",
			HandlerFunc: handlers.GetDeletedClient,
		},
		Route{
			Name: "UnDeleteClient",
			Method: "PUT",
			Pattern: "/client_deleted/{id}",
			HandlerFunc: handlers.UnDeleteClient,
		},
		Route{
			Name: "PermaDeleteClient",
			Method: "DELETE",
			Pattern: "/client/delete/{id}",
			HandlerFunc: handlers.PermaDeleteClient,
		},
	}
}