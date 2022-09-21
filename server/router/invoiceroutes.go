package router

import (
	"server/handlers"
)

func invoiceRoutes(handlers *handlers.Handler) Routes {
	return Routes{
		Route{
			Name: "GetInvoices",
			Method: "GET",
			Pattern: "/invoice",
			HandlerFunc: handlers.GetInvoices,
		},
		Route{
			Name: "CreateInvoices",
			Method: "POST",
			Pattern: "/invoice",
			HandlerFunc: handlers.CreateInvoice,
		},
		Route{
			Name: "GetInvoice",
			Method: "GET",
			Pattern: "/invoice/{id}",
			HandlerFunc: handlers.GetInvoice,
		},
		Route{
			Name: "UpdateInvoice",
			Method: "PUT",
			Pattern: "/invoice/{id}",
			HandlerFunc: handlers.UpdateInvoice,
		},
		Route{
			Name: "DeleteInvoice",
			Method: "DELETE",
			Pattern: "/invoice/{id}",
			HandlerFunc: handlers.DeleteInvoice,
		},
	}
}