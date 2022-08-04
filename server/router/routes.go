package router

import (
	"net/http"
	handlers "server/handlers"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetClients",
		"GET",
		"/clients",
		handlers.GetClients,
	},
	Route{
		"CreateClients",
		"POST",
		"/clients",
		handlers.CreateClient,
	},
	Route{
		"GetClient",
		"GET",
		"/clients/{id}",
		handlers.GetClient,
	},
	Route{
		"UpdateClient",
		"PUT",
		"/clients/{id}",
		handlers.UpdateClient,
	},
	Route{
		"DeleteClient",
		"DELETE",
		"/clients/{id}",
		handlers.DeleteClient,
	},
	Route{
		"GetInvoices",
		"GET",
		"/invoices",
		handlers.GetInvoices,
	},
	Route{
		"CreateInvoices",
		"POST",
		"/invoices",
		handlers.CreateInvoice,
	},
	Route{
		"GetInvoice",
		"GET",
		"/invoices/{id}",
		handlers.GetInvoice,
	},
	Route{
		"UpdateInvoice",
		"PUT",
		"/invoices/{id}",
		handlers.UpdateInvoice,
	},
	Route{
		"DeleteInvoice",
		"DELETE",
		"/invoices/{id}",
		handlers.DeleteInvoice,
	},
	Route{
		"GetExpenses",
		"GET",
		"/expenses",
		handlers.GetExpenses,
	},
	Route{
		"Createexpenses",
		"POST",
		"/expenses",
		handlers.CreateExpense,
	},
	Route{
		"GetExpense",
		"GET",
		"/expenses/{id}",
		handlers.GetExpense,
	},
	Route{
		"UpdateExpense",
		"PUT",
		"/expenses/{id}",
		handlers.UpdateExpense,
	},
	Route{
		"DeleteExpense",
		"DELETE",
		"/expenses/{id}",
		handlers.DeleteExpense,
	},
}