package http

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/edznux/wonder-xss/api"
	"github.com/edznux/wonder-xss/storage/models"
	"github.com/gorilla/mux"
)

func (httpapi *HTTPApi) Handler(router *mux.Router) {
	router.Use(httpapi.jsonMiddleware)
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
	router.HandleFunc(httpapi.UrlPrefix+"/aliases/{id}", httpapi.getAlias).Methods("GET")

	// Loots CRUD
	router.HandleFunc(httpapi.UrlPrefix+"/loots", httpapi.getLoots).Methods("GET")
}

func (httpapi *HTTPApi) NotImplementedYet(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	res.Code = 0
	res.Message = "Not implemented yet"
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) createPayload(w http.ResponseWriter, req *http.Request) {

	var data models.Payload
	var res api.Response

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		res.Code = 1
		res.Message = "Could not decode payload"
		_ = json.NewEncoder(w).Encode(&res)
		return
	}
	returnedPayload, err := api.AddPayload(data.Name, data.Content)
	if err != nil {
		res.Code = 1
		res.Message = "Could not save payload"
		fmt.Println("AddPayload returned an error: ", err)
		_ = json.NewEncoder(w).Encode(&res)
		return
	}

	res = api.Response{Code: 1, Message: "OK", Data: returnedPayload}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) createAlias(w http.ResponseWriter, req *http.Request) {

	var data api.Alias
	var res api.Response

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		res.Code = 1
		res.Message = "Could not decode Alias"
		fmt.Println(err)
		_ = json.NewEncoder(w).Encode(&res)
		return
	}
	fmt.Println("Data recieved & parsed : ", data)
	returnedAlias, err := api.AddAlias(data.Alias, data.PayloadID)
	if err != nil {
		res.Code = 1
		res.Message = "Could not save Alias"
		fmt.Println(err)
		_ = json.NewEncoder(w).Encode(&res)
		return
	}

	res = api.Response{Code: 1, Message: "OK", Data: returnedAlias}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getPayloads(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	fmt.Println("getPayloads")
	payloads, err := api.GetPayloads()
	if err != nil {
		res = api.Response{Code: 1, Message: "Error getting the payloads"}
		json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: payloads}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getLoots(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	fmt.Println("getLoots")
	payloads, err := api.GetLoots()
	if err != nil {
		res = api.Response{Code: 1, Message: "Error getting the payloads"}
		json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: payloads}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getAliases(w http.ResponseWriter, req *http.Request) {
	var res api.Response

	aliases, err := api.GetAliases()
	if err != nil {
		res = api.Response{Code: 1, Message: "Error getting the aliases"}
		json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: aliases}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getPayload(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	vars := mux.Vars(req)
	returnedPayload, err := api.GetPayload(vars["id"])
	fmt.Println(err.Error())
	if err != nil {
		res.Code = 1
		res.Message = "Could not find payload"
		_ = json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: returnedPayload}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getAlias(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	vars := mux.Vars(req)
	returnedAlias, err := api.GetAlias(vars["id"])
	if err != nil {
		res.Code = 1
		res.Message = "Could not find Alias"
		_ = json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: returnedAlias}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) updatePayload(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(api.Response{Code: 1, Message: "Not Implemented yet"})
	w.Write(res)
}

func (httpapi *HTTPApi) deletePayload(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(api.Response{Code: 2, Message: "Not Implemented yet"})
	w.Write(res)
}

func (httpapi *HTTPApi) healthz(w http.ResponseWriter, req *http.Request) {
	res := api.Response{Code: 1, Message: "OK"}
	json.NewEncoder(w).Encode(res)
}
