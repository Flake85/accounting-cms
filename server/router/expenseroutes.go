package router

import (
	"server/handlers"
	"server/model"
)

var expenseRoutes = Routes{
	model.Route{
		Name: "GetExpenses",
		Method: "GET",
		Pattern: "/expense",
		HandlerFunc: handlers.GetExpenses,
	},
	model.Route{
		Name: "CreateExpenses",
		Method: "POST",
		Pattern: "/expense",
		HandlerFunc: handlers.CreateExpense,
	},
	model.Route{
		Name: "GetExpense",
		Method: "GET",
		Pattern: "/expense/{id}",
		HandlerFunc: handlers.GetExpense,
	},
	model.Route{
		Name: "UpdateExpense",
		Method: "PUT",
		Pattern: "/expense/{id}",
		HandlerFunc: handlers.UpdateExpense,
	},
	model.Route{
		Name: "DeleteExpense",
		Method: "DELETE",
		Pattern: "/expense/{id}",
		HandlerFunc: handlers.DeleteExpense,
	},
}