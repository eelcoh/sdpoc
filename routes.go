package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		index,
	},
	Route{
		"EndpointIndex",
		"GET",
		"/Endpoints",
		endpointIndex,
	},
	Route{
		"ServiceRegister",
		"POST",
		"/Services/{service}/Instance/{instance}",
		serviceRegister,
	},
	Route{
		"EndpointShow",
		"GET",
		"/Endpoints/{service}",
		endpointsForService,
	},
	Route{
		"Heath",
		"GET",
		"/health",
		health,
	},
}
