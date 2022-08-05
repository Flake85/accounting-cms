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

	