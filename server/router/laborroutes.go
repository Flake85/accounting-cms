package router

import (
	"server/handlers"
)

func laborRoutes(handlers *handlers.Handler) Routes {
	return Routes{
		Route{
			Name: "GetLabors",
			Method: "GET",
			Pattern: "/labor",
			HandlerFunc: handlers.GetLabors,
		},
		Route{
			Name: "CreateLabor",
			Method: "POST",
			Pattern: "/labor",
			HandlerFunc: handlers.CreateLabor,
		},
		Route{
			Name: "Getlabor",
			Method: "GET",
			Pattern: "/labor/{id}",
			HandlerFunc: handlers.GetLabor,
		},
		Route{
			Name: "UpdateLabor",
			Method: "PUT",
			Pattern: "/labor/{id}",
			HandlerFunc: handlers.UpdateLabor,
		},
		Route{
			Name: "DeleteLabor",
			Method: "DELETE",
			Pattern: "/labor/{id}",
			HandlerFunc: handlers.DeleteLabor,
		},
	}
}