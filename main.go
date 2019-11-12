package main

import (
	"fmt"
	"log"
	"net/http"

	httpApi "github.com/edznux/wonder-xss/api/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting web server")

	r := mux.NewRouter()
	api := httpApi.New()
	api.Handler(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
