package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/gorilla/mux"
)

type HTTPApi struct {
	UrlPrefix string
}

func New() *HTTPApi {
	httpapi := HTTPApi{}
	httpapi.UrlPrefix = "/api/v1"
	return &httpapi
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

func (httpapi *HTTPApi) createCollectors(w http.ResponseWriter, req *http.Request) {

	var data models.Collector
	var res api.Response
	vars := mux.Vars(req)

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		res.Code = 1
		res.Message = "Could not decode Collector"
		_ = json.NewEncoder(w).Encode(&res)
		return
	}

	// override the payloadID from the JSON if it's provided in the URL param.
	payloadID := vars["id"]
	if payloadID == "" {
		payloadID = data.PayloadID
	}

	returnedCollector, err := api.AddCollector(payloadID, data.Data)
	if err != nil {
		res.Code = 1
		res.Message = "Could not save Collector"
		fmt.Println("AddCollector returned an error: ", err)
		_ = json.NewEncoder(w).Encode(&res)
		return
	}

	res = api.Response{Code: 1, Message: "OK", Data: returnedCollector}
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
	fmt.Println("httpapi.getPayloads")
	payloads, err := api.GetPayloads()
	if err != nil {
		res = api.Response{Code: 1, Message: "Error getting the payloads"}
		json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: payloads}
	json.NewEncoder(w).Encode(&res)
}
func (httpapi *HTTPApi) getCollectors(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	fmt.Println("getCollectors")
	collectors, err := api.GetCollectors()
	if err != nil {
		res = api.Response{Code: 1, Message: "Error getting the collectors"}
		json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: collectors}
	json.NewEncoder(w).Encode(&res)
}
func (httpapi *HTTPApi) getExecutions(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	fmt.Println("getExecutions")
	executions, err := api.GetExecutions()
	if err != nil {
		res = api.Response{Code: 1, Message: "Error getting the executions"}
		json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: executions}
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
	returnedAlias, err := api.GetAlias(vars["alias"])
	if err != nil {
		res.Code = 1
		res.Message = "Could not find Alias"
		_ = json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: returnedAlias}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getAliasByID(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	vars := mux.Vars(req)
	returnedAlias, err := api.GetAliasByID(vars["id"])
	if err != nil {
		res.Code = 1
		res.Message = "Could not find Alias"
		_ = json.NewEncoder(w).Encode(&res)
		return
	}
	res = api.Response{Code: 1, Message: "OK", Data: returnedAlias}
	json.NewEncoder(w).Encode(&res)
}

func (httpapi *HTTPApi) getAliasByPayloadID(w http.ResponseWriter, req *http.Request) {
	var res api.Response
	vars := mux.Vars(req)
	returnedAlias, err := api.GetAliasByPayloadID(vars["id"])
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
