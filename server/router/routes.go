package router

import (
	"net/http"
	// handlers "server/handlers"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = []Routes{ 
	clientRoutes,
	invoiceRoutes,
}

	
// var routes = Routes{
// 	Route{
// 		"GetClients",
// 		"GET",
// 		"/client",
// 		handlers.GetClients,
// 	},
// 	Route{
// 		"CreateClients",
// 		"POST",
// 		"/client",
// 		handlers.CreateClient,
// 	},
// 	Route{
// 		"GetClient",
// 		"GET",
// 		"/client/{id}",
// 		handlers.GetClient,
// 	},
// 	Route{
// 		"UpdateClient",
// 		"PUT",
// 		"/client/{id}",
// 		handlers.UpdateClient,
// 	},
// 	Route{
// 		"DeleteClient",
// 		"DELETE",
// 		"/client/{id}",
// 		handlers.DeleteClient,
// 	},
// 	Route{
// 		"GetInvoices",
// 		"GET",
// 		"/invoice",
// 		handlers.GetInvoices,
// 	},
// 	Route{
// 		"CreateInvoices",
// 		"POST",
// 		"/invoice",
// 		handlers.CreateInvoice,
// 	},
// 	Route{
// 		"GetInvoice",
// 		"GET",
// 		"/invoice/{id}",
// 		handlers.GetInvoice,
// 	},
// 	Route{
// 		"UpdateInvoice",
// 		"PUT",
// 		"/invoice/{id}",
// 		handlers.UpdateInvoice,
// 	},
// 	Route{
// 		"DeleteInvoice",
// 		"DELETE",
// 		"/invoice/{id}",
// 		handlers.DeleteInvoice,
// 	},
// 	Route{
// 		"GetExpenses",
// 		"GET",
// 		"/expense",
// 		handlers.GetExpenses,
// 	},
// 	Route{
// 		"Createexpenses",
// 		"POST",
// 		"/expense",
// 		handlers.CreateExpense,
// 	},
// 	Route{
// 		"GetExpense",
// 		"GET",
// 		"/expense/{id}",
// 		handlers.GetExpense,
// 	},
// 	Route{
// 		"UpdateExpense",
// 		"PUT",
// 		"/expense/{id}",
// 		handlers.UpdateExpense,
// 	},
// 	Route{
// 		"DeleteExpense",
// 		"DELETE",
// 		"/expense/{id}",
// 		handlers.DeleteExpense,
// 	},
// 	Route{
// 		"GetLabors",
// 		"GET",
// 		"/labor",
// 		handlers.GetLabors,
// 	},
// 	Route{
// 		"CreateLabor",
// 		"POST",
// 		"/labor",
// 		handlers.CreateLabor,
// 	},
// 	Route{
// 		"Getlabor",
// 		"GET",
// 		"/labor/{id}",
// 		handlers.GetLabor,
// 	},
// 	Route{
// 		"UpdateLabor",
// 		"PUT",
// 		"/labor/{id}",
// 		handlers.UpdateLabor,
// 	},
// 	Route{
// 		"DeleteLabor",
// 		"DELETE",
// 		"/labor/{id}",
// 		handlers.DeleteLabor,
// 	},
// 	Route{
// 		"GetSales",
// 		"GET",
// 		"/sale",
// 		handlers.GetSales,
// 	},
// 	Route{
// 		"CreateSales",
// 		"POST",
// 		"/sale",
// 		handlers.CreateSale,
// 	},
// 	Route{
// 		"GetSale",
// 		"GET",
// 		"/sale/{id}",
// 		handlers.GetSale,
// 	},
// 	Route{
// 		"UpdateSale",
// 		"PUT",
// 		"/sale/{id}",
// 		handlers.UpdateSale,
// 	},
// 	Route{
// 		"DeleteSale",
// 		"DELETE",
// 		"/sale/{id}",
// 		handlers.DeleteSale,
// 	},
// }
