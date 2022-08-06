package router

import (
	"server/handlers"
	"server/model"
)

var invoiceRoutes = Routes{
	model.Route{
		Name: "GetInvoices",
		Method: "GET",
		Pattern: "/invoice",
		HandlerFunc: handlers.GetInvoices,
	},
	model.Route{
		Name: "CreateInvoices",
		Method: "POST",
		Pattern: "/invoice",
		HandlerFunc: handlers.CreateInvoice,
	},
	model.Route{
		Name: "GetInvoice",
		Method: "GET",
		Pattern: "/invoice/{id}",
		HandlerFunc: handlers.GetInvoice,
	},
	model.Route{
		Name: "UpdateInvoice",
		Method: "PUT",
		Pattern: "/invoice/{id}",
		HandlerFunc: handlers.UpdateInvoice,
	},
	model.Route{
		Name: "DeleteInvoice",
		Method: "DELETE",
		Pattern: "/invoice/{id}",
		HandlerFunc: handlers.DeleteInvoice,
	},
}