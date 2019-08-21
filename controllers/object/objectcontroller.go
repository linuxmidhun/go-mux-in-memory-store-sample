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
		if cID == 0 {
			log.Println("looking for invalid index '0'")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(http.StatusBadRequest)
		} else {
			d, err := data.Get(cID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
