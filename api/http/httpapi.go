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
	if status != api.Success {
		res.Error = status.Error()
	}
	res.Data = data

	err := json.NewEncoder(w).Encode(&res)
	return err
}

func (httpapi *HTTPApi) createInjection(w http.ResponseWriter, req *http.Request) {
	var data models.Injection

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedInjection, err := api.AddInjection(data.Name, data.Content)
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedInjection, w)
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

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedCollector, err := api.AddCollector(data.Data)
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

	log.Println("Data recieved & parsed : ", data)
	if data.Alias == "" || data.PayloadID == "" {
		log.Println("Empty data")
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedAlias, err := api.AddAlias(data.Alias, data.PayloadID)
	if err == models.AlreadyExist {
		log.Println("Returned alias error: ", err)
		sendResponse(api.AlreadyExist, nil, w)
		return
	}
	if err != nil {
		log.Println("Returned alias DatabaseError", err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) getUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("httpapi.getUser")
	vars := mux.Vars(req)
	user, err := api.GetUser((vars["id"]))
	if err != nil {
		log.Println(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, user, w)
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
func (httpapi *HTTPApi) getInjection(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedInjection, err := api.GetInjection(vars["injection"])
	if err != nil {
		log.Println(err)
		sendResponse(api.NotFound, nil, w)
		return
	}

	sendResponse(api.Success, returnedInjection, w)
}
func (httpapi *HTTPApi) getInjections(w http.ResponseWriter, req *http.Request) {
	returnedInjections, err := api.GetInjections()
	if err != nil {
		log.Println(err)
		sendResponse(api.NotFound, nil, w)
		return
	}

	sendResponse(api.Success, returnedInjections, w)
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
	res, _ := json.Marshal(api.Response{Error: "Not Implemented yet"})
	w.Write(res)
}

func (httpapi *HTTPApi) deletePayload(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	api.DeletePayload(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteExecution(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	api.DeleteExecution(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	api.DeleteUser(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteAlias(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	api.DeleteAlias(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteInjection(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	api.DeleteInjection(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteCollector(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	api.DeleteCollector(vars["id"])
	sendResponse(api.Success, "", w)
}
