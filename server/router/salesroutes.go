package router

import "server/handlers"

func salesRoutes(handlers *handlers.Handler) Routes {
	return Routes{
		Route{
			Name: "GetSales",
			Method: "GET",
			Pattern: "/sale",
			HandlerFunc: handlers.GetSales,
		},
		Route{
			Name: "CreateSales",
			Method: "POST",
			Pattern: "/sale",
			HandlerFunc: handlers.CreateSale,
		},
		Route{
			Name: "GetSale",
			Method: "GET",
			Pattern: "/sale/{id}",
			HandlerFunc: handlers.GetSale,
		},
		Route{
			Name: "UpdateSale",
			Method: "PUT",
			Pattern: "/sale/{id}",
			HandlerFunc: handlers.UpdateSale,
		},
		Route{
			Name: "DeleteSale",
			Method: "DELETE",
			Pattern: "/sale/{id}",
			HandlerFunc: handlers.DeleteSale,
		},
	}	
}