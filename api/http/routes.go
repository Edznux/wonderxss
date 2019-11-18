package http

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/edznux/wonderxss/api"
	"github.com/gorilla/mux"
)

func (httpapi *HTTPApi) Routes(router *mux.Router) {
	router.Use(httpapi.jsonMiddleware)
	router.Use(httpapi.CORSMiddleware)
	// Return real payload
	router.HandleFunc("/p/{id}", httpapi.handlePayloadByID)

	// HealthZ endpoint.
	router.HandleFunc(httpapi.UrlPrefix+"/healthz", httpapi.healthz)

	// Authentication
	router.HandleFunc("/login", httpapi.NotImplementedYet).Methods("POST")
	router.HandleFunc("/logout", httpapi.NotImplementedYet).Methods("POST")

	// Payload CRUD
	router.HandleFunc(httpapi.UrlPrefix+"/payloads", httpapi.createPayload).Methods("POST")
	router.HandleFunc(httpapi.UrlPrefix+"/payloads", httpapi.getPayloads).Methods("GET")
	router.HandleFunc(httpapi.UrlPrefix+"/payloads/{id}", httpapi.getPayload).Methods("GET")
	router.HandleFunc(httpapi.UrlPrefix+"/payloads/{id}", httpapi.updatePayload).Methods("PUT")
	router.HandleFunc(httpapi.UrlPrefix+"/payloads/{id}", httpapi.deletePayload).Methods("DELETE")

	// Aliases CRUD
	router.HandleFunc(httpapi.UrlPrefix+"/aliases", httpapi.createAlias).Methods("POST")
	router.HandleFunc(httpapi.UrlPrefix+"/aliases", httpapi.getAliases).Methods("GET")
	router.HandleFunc(httpapi.UrlPrefix+"/aliases/{alias}", httpapi.getAlias).Methods("GET")
	router.HandleFunc(httpapi.UrlPrefix+"/aliases/id/{id}", httpapi.getAliasByID).Methods("GET")
	router.HandleFunc(httpapi.UrlPrefix+"/aliases/payload/{id}", httpapi.getAliasByPayloadID).Methods("GET")

	// Executions CRUD
	router.HandleFunc(httpapi.UrlPrefix+"/executions", httpapi.getExecutions).Methods("GET")

	// Colletors CRUD
	router.HandleFunc(httpapi.UrlPrefix+"/collectors", httpapi.getCollectors).Methods("GET")
	router.HandleFunc(httpapi.UrlPrefix+"/collectors", httpapi.createCollectors).Methods("POST")
}

func (httpapi *HTTPApi) NotImplementedYet(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	res.Code = 0
	res.Message = "Not implemented yet"
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) healthz(w http.ResponseWriter, req *http.Request) {
	res := api.Response{Code: 1, Message: "OK"}
	json.NewEncoder(w).Encode(res)
}
