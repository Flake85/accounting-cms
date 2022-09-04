package router

import (
	"server/handlers"
	"server/model"
)

var clientRoutes = Routes{
	model.Route{
		Name: "GetClients",
		Method: "GET",
		Pattern: "/client",
		HandlerFunc: handlers.GetClients,
	},
	model.Route{
		Name: "CreateClients",
		Method: "POST",
		Pattern: "/client",
		HandlerFunc: handlers.CreateClient,
	},
	model.Route{
		Name: "GetClient",
		Method: "GET",
		Pattern: "/client/{id}",
		HandlerFunc: handlers.GetClient,
	},
	model.Route{
		Name: "UpdateClient",
		Method: "PUT",
		Pattern: "/client/{id}",
		HandlerFunc: handlers.UpdateClient,
	},
	model.Route{
		Name: "DeleteClient",
		Method: "DELETE",
		Pattern: "/client/{id}",
		HandlerFunc: handlers.DeleteClient,
	},
	model.Route{
		Name: "GetDeletedClients",
		Method: "GET",
		Pattern: "/client_deleted",
		HandlerFunc: handlers.GetDeletedClients,
	},
	model.Route{
		Name: "GetDeletedClient",
		Method: "GET",
		Pattern: "/client_deleted/{id}",
		HandlerFunc: handlers.GetDeletedClient,
	},
	model.Route{
		Name: "UnDeleteClient",
		Method: "PUT",
		Pattern: "/client_deleted/{id}",
		HandlerFunc: handlers.UnDeleteClient,
	},
	model.Route{
		Name: "PermaDeleteClient",
		Method: "DELETE",
		Pattern: "/client/delete/{id}",
		HandlerFunc: handlers.PermaDeleteClient,
	},
}