package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// teamCreationHandler creates a new team based on the team name provided by user.
// The team id is a randomly generated id.
func teamCreationHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user new team name as a multipart form POST
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Team could not be created because of malformed POST parameters"))
		return
	}
	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Team could not be created because of empty POST parameter"))
		return
	}

	// Generate a UUID to avoid ids collision
	team := Team{ID: uuid.New().String(), Name: name}
	teams = append(teams, team)

	// Return the created team to user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}

// teamDeletionHandler takes a team id and removes the matching team.
// If no matching team can be found, it returns a 404 page.
func teamDeletionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Look for a team with the id retrieved from user and delete it
	// if found
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

// teamsListingHandler returns a json encoded list of all the teams
func teamsListingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func main() {
	// Declare HTTP routes
	r := mux.NewRouter()
	r.HandleFunc("/teams", teamCreationHandler).Methods("POST")
	r.HandleFunc("/teams/{id}", teamDeletionHandler).Methods("DELETE")
	r.HandleFunc("/teams", teamsListingHandler).Methods("GET")

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8000", r))
}
