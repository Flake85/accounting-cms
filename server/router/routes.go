package router

import (
	"server/model"
)

type Routes []model.Route

var routes = []Routes{ 
	clientRoutes,
	invoiceRoutes,
	expenseRoutes,
	laborRoutes,
	salesRoutes,
}
