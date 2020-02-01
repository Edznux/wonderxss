package http

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"encoding/json"

	"github.com/edznux/wonderxss/api"
	"github.com/gorilla/mux"
)

func (httpapi *HTTPApi) Routes(router *mux.Router) {
	router.Use(httpapi.jsonMiddleware)
	router.Use(httpapi.CORSMiddleware)
	router.Use(httpapi.authMiddleware)

	// HealthZ endpoint.
	router.HandleFunc("/healthz", httpapi.healthz)
	// User crud
	router.HandleFunc("/users/{id}", httpapi.getUser).Methods("GET")
	// Payload CRUD
	router.HandleFunc("/payloads", httpapi.createPayload).Methods("POST")
	router.HandleFunc("/payloads", httpapi.getPayloads).Methods("GET")
	router.HandleFunc("/payloads/{id}", httpapi.getPayload).Methods("GET")
	router.HandleFunc("/payloads/{id}", httpapi.updatePayload).Methods("PUT")
	router.HandleFunc("/payloads/{id}", httpapi.deletePayload).Methods("DELETE")

	// Aliases CRUD
	router.HandleFunc("/aliases", httpapi.createAlias).Methods("POST")
	router.HandleFunc("/aliases", httpapi.getAliases).Methods("GET")
	router.HandleFunc("/aliases/{alias}", httpapi.getAlias).Methods("GET")
	router.HandleFunc("/aliases/id/{id}", httpapi.getAliasByID).Methods("GET")
	router.HandleFunc("/aliases/payload/{id}", httpapi.getAliasByPayloadID).Methods("GET")
	router.HandleFunc("/aliases/{id}", httpapi.deleteAlias).Methods("DELETE")

	// Executions CRUD
	router.HandleFunc("/executions", httpapi.getExecutions).Methods("GET")
	router.HandleFunc("/executions/{id}", httpapi.deleteExecution).Methods("DELETE")

	// Colletors CRUD
	router.HandleFunc("/collectors", httpapi.getCollectors).Methods("GET")
	router.HandleFunc("/collectors", httpapi.createCollectors).Methods("POST")
	router.HandleFunc("/collectors/{id}", httpapi.deleteCollector).Methods("DELETE")

	// Colletors CRUD
	router.HandleFunc("/injections/{name}", httpapi.getInjection).Methods("GET")
	router.HandleFunc("/injections", httpapi.getInjections).Methods("GET")
	router.HandleFunc("/injections", httpapi.createInjection).Methods("POST")
	router.HandleFunc("/injections/{id}", httpapi.deleteInjection).Methods("DELETE")
}

func (httpapi *HTTPApi) NotImplementedYet(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	res.Error = "Not implemented yet"
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) healthz(w http.ResponseWriter, req *http.Request) {
	res := api.Response{Data: "OK"}
	json.NewEncoder(w).Encode(res)
}
