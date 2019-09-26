package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func teamCreationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// TODO: properly retrieve JSON post data
	team := Team{
		ID:   "20",
		Name: vars["name"],
	}
	teams = append(teams, team)
	w.Write([]byte("Team successfully created"))
}

// teamDeletionHandler takes a team id and removes the matching team.
// If no matching team can be found, it returns a 404 page
func teamDeletionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for i, team := range teams {
		if team.ID == vars["id"] {
			teams = append(teams[:i], teams[i+1:]...)
			w.Write([]byte("Team successfully deleted"))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Team not found"))
}

func teamsListingHandler(w http.ResponseWriter, r *http.Request) {
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
