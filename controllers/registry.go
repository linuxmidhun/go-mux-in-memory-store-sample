package controllers

import (
	"log"
	"net/http"

	object "./object"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New .
func New() http.Handler {
	log.Println("Setting up end points")
	r := mux.NewRouter()

	r.HandleFunc("/object/{id}", object.Get).Methods("GET")
	r.HandleFunc("/object", object.Create).Methods("POST")
	return handlers.CORS(handlers.AllowedMethods([]string{"GET",
		"POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r)
}
