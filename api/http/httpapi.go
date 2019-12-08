package http

import (
	"encoding/json"
	"fmt"
	"log"
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

func sendResponse(status api.APIError, data interface{}, w http.ResponseWriter) error {
	var res api.Response
	res.Code = int(status)
	res.Message = status.Error()
	res.Data = data
	err := json.NewEncoder(w).Encode(&res)
	return err
}

func (httpapi *HTTPApi) createPayload(w http.ResponseWriter, req *http.Request) {
	var data models.Payload

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedPayload, err := api.AddPayload(data.Name, data.Content)
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedPayload, w)
}

func (httpapi *HTTPApi) createCollectors(w http.ResponseWriter, req *http.Request) {

	var data models.Collector
	vars := mux.Vars(req)

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	// override the payloadID from the JSON if it's provided in the URL param.
	payloadID := vars["id"]
	if payloadID == "" {
		payloadID = data.PayloadID
	}

	returnedCollector, err := api.AddCollector(payloadID, data.Data)
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedCollector, w)
}

func (httpapi *HTTPApi) createAlias(w http.ResponseWriter, req *http.Request) {

	var data api.Alias

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}
	fmt.Println("Data recieved & parsed : ", data)
	returnedAlias, err := api.AddAlias(data.Alias, data.PayloadID)
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) getPayloads(w http.ResponseWriter, req *http.Request) {
	fmt.Println("httpapi.getPayloads")
	payloads, err := api.GetPayloads()
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, payloads, w)
}

func (httpapi *HTTPApi) getCollectors(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getCollectors")
	collectors, err := api.GetCollectors()
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, collectors, w)
}
func (httpapi *HTTPApi) getExecutions(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getExecutions")
	executions, err := api.GetExecutions()
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}
	sendResponse(api.Success, executions, w)
}

func (httpapi *HTTPApi) getAliases(w http.ResponseWriter, req *http.Request) {
	aliases, err := api.GetAliases()
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}
	sendResponse(api.Success, aliases, w)
}

func (httpapi *HTTPApi) getPayload(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedPayload, err := api.GetPayload(vars["id"])
	if err != nil {
		log.Println(err)
		sendResponse(api.NotFound, nil, w)
		return
	}
	sendResponse(api.Success, returnedPayload, w)
}

func (httpapi *HTTPApi) getAlias(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedAlias, err := api.GetAlias(vars["alias"])
	if err != nil {
		log.Println(err)
		sendResponse(api.NotFound, nil, w)
		return
	}

	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) getAliasByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedAlias, err := api.GetAliasByID(vars["id"])
	if err != nil {
		log.Println(err)
		sendResponse(api.NotFound, nil, w)
		return
	}
	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) getAliasByPayloadID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedAlias, err := api.GetAliasByPayloadID(vars["id"])
	if err != nil {
		log.Println(err)
		sendResponse(api.NotFound, nil, w)
		return
	}
	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) updatePayload(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(api.Response{Code: 1, Message: "Not Implemented yet"})
	w.Write(res)
}

func (httpapi *HTTPApi) deletePayload(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(api.Response{Code: 2, Message: "Not Implemented yet"})
	w.Write(res)
}
