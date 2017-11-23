package main

import (
	"log"
	"net/http"

	"github.com/eelcoh/sdlib"
)

// Endpoints = database
var endpoints []sdlib.Endpoint

func main() {

	endpoints = []sdlib.Endpoint{}
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
