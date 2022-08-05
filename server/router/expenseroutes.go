package router

import "server/handlers"

var expenseRoutes = Routes{
	Route{
		"GetExpenses",
		"GET",
		"/expense",
		handlers.GetExpenses,
	},
	Route{
		"Createexpenses",
		"POST",
		"/expense",
		handlers.CreateExpense,
	},
	Route{
		"GetExpense",
		"GET",
		"/expense/{id}",
		handlers.GetExpense,
	},
	Route{
		"UpdateExpense",
		"PUT",
		"/expense/{id}",
		handlers.UpdateExpense,
	},
	Route{
		"DeleteExpense",
		"DELETE",
		"/expense/{id}",
		handlers.DeleteExpense,
	},
}