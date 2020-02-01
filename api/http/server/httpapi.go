package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/api/local"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/gorilla/mux"
)

type HTTPApi struct {
	UrlPrefix string
	local     api.API
}

func New() *HTTPApi {
	httpapi := HTTPApi{}
	httpapi.UrlPrefix = "/api/v1"
	log.Debugln("Connecting to local API from HTTPApi")
	httpapi.local = local.New()
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
	var data api.Injection

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedInjection, err := httpapi.local.AddInjection(data.Name, data.Content)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedInjection, w)
}
func (httpapi *HTTPApi) createPayload(w http.ResponseWriter, req *http.Request) {
	var data api.Payload

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedPayload, err := httpapi.local.AddPayload(data.Name, data.Content, data.ContentType)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedPayload, w)
}

func (httpapi *HTTPApi) createCollectors(w http.ResponseWriter, req *http.Request) {

	var data api.Collector

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedCollector, err := httpapi.local.AddCollector(data.Data)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedCollector, w)
}

func (httpapi *HTTPApi) createAlias(w http.ResponseWriter, req *http.Request) {

	var data api.Alias

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Warnln(err)
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	log.Debugln("Data recieved & parsed : ", data)
	if data.Alias == "" || data.PayloadID == "" {
		log.Warnln("Empty data")
		sendResponse(api.InvalidInput, nil, w)
		return
	}

	returnedAlias, err := httpapi.local.AddAlias(data.Alias, data.PayloadID)
	if err == models.AlreadyExist {
		log.Warnln("Returned alias error: ", err)
		sendResponse(api.AlreadyExist, nil, w)
		return
	}
	if err != nil {
		log.Warnln("Returned alias DatabaseError", err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) getUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user, err := httpapi.local.GetUser((vars["id"]))
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, user, w)
}
func (httpapi *HTTPApi) getPayloads(w http.ResponseWriter, req *http.Request) {
	payloads, err := httpapi.local.GetPayloads()
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, payloads, w)
}

func (httpapi *HTTPApi) getCollectors(w http.ResponseWriter, req *http.Request) {
	collectors, err := httpapi.local.GetCollectors()
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}

	sendResponse(api.Success, collectors, w)
}
func (httpapi *HTTPApi) getExecutions(w http.ResponseWriter, req *http.Request) {
	executions, err := httpapi.local.GetExecutions()
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}
	sendResponse(api.Success, executions, w)
}

func (httpapi *HTTPApi) getAliases(w http.ResponseWriter, req *http.Request) {
	aliases, err := httpapi.local.GetAliases()
	if err != nil {
		log.Warnln(err)
		sendResponse(api.DatabaseError, nil, w)
		return
	}
	sendResponse(api.Success, aliases, w)
}

func (httpapi *HTTPApi) getPayload(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedPayload, err := httpapi.local.GetPayload(vars["id"])
	if err != nil {
		log.Warnln(err)
		sendResponse(api.NotFound, nil, w)
		return
	}
	sendResponse(api.Success, returnedPayload, w)
}

func (httpapi *HTTPApi) getAlias(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedAlias, err := httpapi.local.GetAlias(vars["alias"])
	if err != nil {
		log.Warnln(err)
		sendResponse(api.NotFound, nil, w)
		return
	}

	sendResponse(api.Success, returnedAlias, w)
}
func (httpapi *HTTPApi) getInjection(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedInjection, err := httpapi.local.GetInjection(vars["injection"])
	if err != nil {
		log.Warnln(err)
		sendResponse(api.NotFound, nil, w)
		return
	}

	sendResponse(api.Success, returnedInjection, w)
}
func (httpapi *HTTPApi) getInjections(w http.ResponseWriter, req *http.Request) {
	returnedInjections, err := httpapi.local.GetInjections()
	if err != nil {
		log.Warnln(err)
		sendResponse(api.NotFound, nil, w)
		return
	}

	sendResponse(api.Success, returnedInjections, w)
}

func (httpapi *HTTPApi) getAliasByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedAlias, err := httpapi.local.GetAliasByID(vars["id"])
	if err != nil {
		log.Warnln(err)
		sendResponse(api.NotFound, nil, w)
		return
	}
	sendResponse(api.Success, returnedAlias, w)
}

func (httpapi *HTTPApi) getAliasByPayloadID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	returnedAlias, err := httpapi.local.GetAliasByPayloadID(vars["id"])
	if err != nil {
		log.Warnln(err)
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
	httpapi.local.DeletePayload(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteExecution(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	httpapi.local.DeleteExecution(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	httpapi.local.DeleteUser(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteAlias(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	httpapi.local.DeleteAlias(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteInjection(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	httpapi.local.DeleteInjection(vars["id"])
	sendResponse(api.Success, "", w)
}

func (httpapi *HTTPApi) deleteCollector(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	httpapi.local.DeleteCollector(vars["id"])
	sendResponse(api.Success, "", w)
}
