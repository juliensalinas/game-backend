package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Team represents gaming team of players
type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func teamCreationHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func teamDeletionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func teamsListingHandler(w http.ResponseWriter, r *http.Request) {
	var teams = []Team{
		{
			ID:   "1",
			Name: "Amazing team",
		},
		{
			ID:   "2",
			Name: "Amazing team 2",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/teams", teamCreationHandler).Methods("POST")
	r.HandleFunc("/teams/{id}", teamDeletionHandler).Methods("DELETE")
	r.HandleFunc("/teams", teamsListingHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
