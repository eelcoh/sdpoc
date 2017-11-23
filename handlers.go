package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eelcoh/sdlib"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up!\n")
}

func endpointIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(endpoints); err != nil {
		panic(err)
	}
}

func endpointsForService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var endpoints []sdlib.Endpoint
	var svc string

	svc = vars["service"]

	endpoints = FindEndpoints(svc)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(endpoints); err != nil {
		panic(err)
	}

}
