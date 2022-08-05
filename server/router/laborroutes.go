package router

import (
	"server/handlers"
	"server/model"
)

var laborRoutes = Routes{
	model.Route{
		Name: "GetLabors",
		Method: "GET",
		Pattern: "/labor",
		HandlerFunc: handlers.GetLabors,
	},
	model.Route{
		Name: "CreateLabor",
		Method: "POST",
		Pattern: "/labor",
		HandlerFunc: handlers.CreateLabor,
	},
	model.Route{
		Name: "Getlabor",
		Method: "GET",
		Pattern: "/labor/{id}",
		HandlerFunc: handlers.GetLabor,
	},
	model.Route{
		Name: "UpdateLabor",
		Method: "PUT",
		Pattern: "/labor/{id}",
		HandlerFunc: handlers.UpdateLabor,
	},
	model.Route{
		Name: "DeleteLabor",
		Method: "DELETE",
		Pattern: "/labor/{id}",
		HandlerFunc: handlers.DeleteLabor,
	},
}