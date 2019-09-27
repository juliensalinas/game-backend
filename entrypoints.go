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
	t := Team{ID: uuid.New().String(), Name: name}
	teams = append(teams, t)

	// Return the created team to user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

// teamDeletionHandler takes a team id and removes the matching team.
// If no matching team can be found, it returns a 404 page.
func teamDeletionHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve url positional arguments
	vars := mux.Vars(r)

	// Look for a team with the id retrieved from user and delete it
	// if found
	for i, t := range teams {
		if t.ID == vars["id"] {
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

// playerCreationHandler creates a player and affects him to a team based
// on the team id received
func playerCreationHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Payer could not be created because of malformed POST parameters"))
		return
	}
	pseudo := r.Form.Get("pseudo")
	if pseudo == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Player could not be created because of empty POST parameter"))
		return
	}

	vars := mux.Vars(r)

	// Find the correct team, create the new player, and affect him to the team
	for _, t := range teams {
		if t.ID == vars["id"] {
			p := Player{Team: t, ID: uuid.New().String(), Pseudo: pseudo}
			players = append(players, p)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Team not found"))

}

// playerDeletionHandler deletes a player based on the player id received
func playerDeletionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for i, p := range players {
		if p.ID == vars["id"] {
			players = append(players[:i], players[i+1:]...)
			w.Write([]byte("Player successfully deleted"))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Player not found"))
}

// playersListingHandler displays all players
func playersListingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func main() {
	// Declare HTTP routes
	r := mux.NewRouter()
	r.HandleFunc("/teams", teamCreationHandler).Methods("POST")
	r.HandleFunc("/teams/{id}", teamDeletionHandler).Methods("DELETE")
	r.HandleFunc("/teams", teamsListingHandler).Methods("GET")
	r.HandleFunc("/teams/{id}/players", playerCreationHandler).Methods("POST")
	r.HandleFunc("/players/{id}", playerDeletionHandler).Methods("DELETE")
	r.HandleFunc("/players", playersListingHandler).Methods("GET")

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8000", r))
}
