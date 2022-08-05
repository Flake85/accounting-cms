package router

import "server/handlers"

var invoiceRoutes = Routes{
	Route{
		"GetInvoices",
		"GET",
		"/invoice",
		handlers.GetInvoices,
	},
	Route{
		"CreateInvoices",
		"POST",
		"/invoice",
		handlers.CreateInvoice,
	},
	Route{
		"GetInvoice",
		"GET",
		"/invoice/{id}",
		handlers.GetInvoice,
	},
	Route{
		"UpdateInvoice",
		"PUT",
		"/invoice/{id}",
		handlers.UpdateInvoice,
	},
	Route{
		"DeleteInvoice",
		"DELETE",
		"/invoice/{id}",
		handlers.DeleteInvoice,
	},
}