package router

import (
	"server/handlers"
	"server/model"
)

var salesRoutes = Routes{
	model.Route{
		Name: "GetSales",
		Method: "GET",
		Pattern: "/sale",
		HandlerFunc: handlers.GetSales,
	},
	model.Route{
		Name: "CreateSales",
		Method: "POST",
		Pattern: "/sale",
		HandlerFunc: handlers.CreateSale,
	},
	model.Route{
		Name: "GetSale",
		Method: "GET",
		Pattern: "/sale/{id}",
		HandlerFunc: handlers.GetSale,
	},
	model.Route{
		Name: "UpdateSale",
		Method: "PUT",
		Pattern: "/sale/{id}",
		HandlerFunc: handlers.UpdateSale,
	},
	model.Route{
		Name: "DeleteSale",
		Method: "DELETE",
		Pattern: "/sale/{id}",
		HandlerFunc: handlers.DeleteSale,
	},
}