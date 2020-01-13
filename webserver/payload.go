package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (ui *UI) HandlePayloadByID(w http.ResponseWriter, req *http.Request) {
	var err error

	params := mux.Vars(req)
	id := params["id"]
	text, err := ui.api.ServePayload(id)
	if err != nil {
		fmt.Printf("Could not get payload to be served as a /p/%s, error : %s\n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(text))
}
