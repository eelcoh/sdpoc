package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/eelcoh/sdlib"
	"github.com/gorilla/mux"
)

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"Hostname":"h2", "IP":"1.2.3.4", "Version":"0.0.1", "Endpoints": [{"Method":"GET", "Path":"/"}]}' http://localhost:8080/Service/svcname/Instance/instname
*/
func serviceRegister(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var service sdlib.Manifest

	var svc string
	var inst string
	var err error

	svc = vars["service"]
	inst = vars["instance"]

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &service); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := serviceToEndpoints(service, svc, inst)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func serviceToEndpoints(service sdlib.Manifest, svc string, instance string) []sdlib.Endpoint {

	for _, p := range service.Paths {
		CreateEndpoint(instance, service.Hostname, service.IP, svc, p.Method, p.Path)
	}

	return endpoints

}
