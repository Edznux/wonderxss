package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HandlePayloadByID Serve payload by its ID (used only in subpath)
func (ui *UI) HandlePayloadByID(w http.ResponseWriter, req *http.Request) {
	var err error

	params := mux.Vars(req)
	id := params["id"]
	//TODO: change content type to stored one.
	payload, err := ui.api.GetPayload(id)
	if err != nil {
		fmt.Printf("Could not get payload to be served as a /p/%s, error : %s\n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", payload.ContentType)
	w.Write([]byte(payload.Content))
}
