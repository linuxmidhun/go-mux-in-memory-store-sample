package object

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	data "../../models"
	"github.com/gorilla/mux"
)

// Get .
func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	cID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in converting given id to int", id)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		d, err := data.Get(cID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusInternalServerError)
		} else {
			if d.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(http.StatusNotFound)
			} else {
				json.NewEncoder(w).Encode(d)
			}
		}
	}
}

// Create .
func Create(w http.ResponseWriter, r *http.Request) {
	var d data.Data
	_ = json.NewDecoder(r.Body).Decode(&d)
	log.Println(d)
	data, err := data.Create(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		if data.ID == 0 {
			w.WriteHeader(http.StatusPreconditionFailed)
			json.NewEncoder(w).Encode(http.StatusPreconditionFailed)
		} else {
			json.NewEncoder(w).Encode(data)
		}
	}
}
