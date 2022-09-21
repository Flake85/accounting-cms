package router

import (
	"server/handlers"
)

func expenseRoutes(handlers *handlers.Handler) Routes {
	return Routes{
		Route{
			Name: "GetExpenses",
			Method: "GET",
			Pattern: "/expense",
			HandlerFunc: handlers.GetExpenses,
		},
		Route{
			Name: "CreateExpenses",
			Method: "POST",
			Pattern: "/expense",
			HandlerFunc: handlers.CreateExpense,
		},
		Route{
			Name: "GetExpense",
			Method: "GET",
			Pattern: "/expense/{id}",
			HandlerFunc: handlers.GetExpense,
		},
		Route{
			Name: "UpdateExpense",
			Method: "PUT",
			Pattern: "/expense/{id}",
			HandlerFunc: handlers.UpdateExpense,
		},
		Route{
			Name: "DeleteExpense",
			Method: "DELETE",
			Pattern: "/expense/{id}",
			HandlerFunc: handlers.DeleteExpense,
		},
	}
}