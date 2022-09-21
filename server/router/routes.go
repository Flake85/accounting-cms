package router

import (
	"net/http"
	"server/handlers"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func routes(h *handlers.Handler) []Routes { 
	return []Routes{ 
		clientRoutes(h),
		invoiceRoutes(h),
		expenseRoutes(h),
		laborRoutes(h),
		salesRoutes(h),
	}
}