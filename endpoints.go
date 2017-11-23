package main

import (
	"sync"

	"github.com/eelcoh/sdlib"
)

var mu sync.Mutex

// Endpoints is the global in memory database :)
//var Endpoints []sdlib.Endpoint

// FindEndpoints is the function that searches for all endpoints for a given service
func FindEndpoints(svc string) []sdlib.Endpoint {
	endpoints := []sdlib.Endpoint{}

	for _, e := range endpoints {
		if e.Service == svc {
			endpoints = append(endpoints, e)
		}
	}
	// return empty Endpoint if not found
	return endpoints
}

// CreateEndpoint and add it to the database
func CreateEndpoint(inst, hostname, ip, service, method, path string) []sdlib.Endpoint {

	e := sdlib.Endpoint{Instance: inst,
		Hostname: hostname,
		IP:       ip,
		Service:  service,
		Method:   method,
		Path:     path,
	}

	mu.Lock()
	endpoints = append(endpoints, e)
	mu.Unlock()

	return endpoints
}
