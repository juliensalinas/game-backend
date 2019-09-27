package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
		w.Write([]byte("Player could not be created because of malformed POST parameters"))
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
	for i, t := range teams {
		if t.ID == vars["id"] {
			p := Player{ID: uuid.New().String(), Pseudo: pseudo}
			t.AddPlayer(p)
			teams[i] = t
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

	// Look for the right team
	for i, t := range teams {
		// Look for the right player and remove him if found
		if t.ID == vars["teamId"] {
			ok, _ := t.RemovePlayer(vars["playerId"])
			if ok {
				teams[i] = t
				w.Write([]byte("Player successfully deleted"))
				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Team or player not found"))
}

// gameCreationHandler creates a player and affects him to a team based
// on the team id received
func gameCreationHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Game could not be created because of malformed POST parameters"))
		return
	}
	name := r.Form.Get("name")
	team1Id := r.Form.Get("team1Id")
	team2Id := r.Form.Get("team2Id")
	if name == "" || team1Id == "" || team2Id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Game could not be created because of empty POST parameter"))
		return
	}

	// Create the game and set a starting time
	g := Game{ID: uuid.New().String(), Name: name, StartTime: time.Now()}

	// Affect team 1 and team 2 to the game.
	// If at least of the teams cannot be found, stop here and return an error.
	for _, t := range teams {
		if t.ID == team1Id {
			g.Team1 = t
		}
	}
	if g.Team1.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No team 1 could be found with this id"))
		return
	}
	for _, t := range teams {
		if t.ID == team2Id {
			g.Team2 = t
		}
	}
	if g.Team2.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No team 1 could be found with this id"))
		return
	}

	// Check that teams are not equal
	if g.Team1.ID == g.Team2.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The teams should not be equal"))
		return
	}

	// Check team sizes
	if !g.TeamSizesAreValid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The teams should have from 3 to 5 players and have the same size"))
		return
	}

	games = append(games, g)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)

}

// gameStopHandler stops a game by setting a stop time
func gameStopHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Look for the right game
	for i, g := range games {
		// Stop the game if found
		if g.ID == vars["id"] {
			g.Stop()
			games[i] = g
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Team or player not found"))
}

// gamesListingHandler returns a json encoded list of all the games
func gamesListingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func main() {
	// Declare HTTP routes
	r := mux.NewRouter()
	r.HandleFunc("/teams", teamCreationHandler).Methods("POST")
	r.HandleFunc("/teams/{id}", teamDeletionHandler).Methods("DELETE")
	r.HandleFunc("/teams", teamsListingHandler).Methods("GET")
	r.HandleFunc("/teams/{id}/players", playerCreationHandler).Methods("POST")
	r.HandleFunc("/teams/{teamId}/players/{playerId}", playerDeletionHandler).Methods("DELETE")
	r.HandleFunc("/games", gameCreationHandler).Methods("POST")
	r.HandleFunc("/games/{id}", gameStopHandler).Methods("DELETE")
	r.HandleFunc("/games", gamesListingHandler).Methods("GET")

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8000", r))
}
