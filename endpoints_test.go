package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Global variables use to keep state accross tests
var testTeam1, testTeam2, testTeam3, testTeam4 Team
var testPlayer1, testPlayer2, testPlayer3, testPlayer4, testPlayer5, testPlayer6, testPlayer7 Player

// TestTeamCreationHandler tests the creation of 4 teams
func TestTeamCreationHandler(t *testing.T) {
	for i := 1; i < 5; i++ {
		// POST param to pass
		teamName := fmt.Sprintf("Best Team Ever %d", i)

		// Create a request to pass to our handler. Pass the team name as form urlencoded POST parameters.
		params := url.Values{}
		params.Set("name", teamName)
		req, err := http.NewRequest("POST", "/teams", strings.NewReader(params.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(teamCreationHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		var team Team
		json.Unmarshal([]byte(rr.Body.String()), &team)

		// Team id should be not empty
		if team.ID == "" {
			t.Errorf("handler returned unexpected team name in body: got %v want %v",
				team.ID, "")
		}
		// Team name should be the one we passed earlier
		if team.Name != teamName {
			t.Errorf("handler returned unexpected team name in body: got %v want %v",
				team.Name, teamName)
		}
		// Team players should be empty
		if team.Players != nil {
			t.Errorf("handler returned unexpected team name in body: got %v want %v",
				team.Players, nil)
		}

		// Make team available to other tests by making it global
		switch i {
		case 1:
			testTeam1 = team
		case 2:
			testTeam2 = team
		case 3:
			testTeam3 = team
		case 4:
			testTeam4 = team
		}
	}
}

// TestPlayerCreationHandler tests the creation of 3 players in teams 1 and 2,
// 1 player in team 3, and 0 player in team 4
func TestPlayerCreationHandler(t *testing.T) {
	// Create and add 3 players to teams 1 and 2, and 1 player
	// to team 3
	for i := 1; i < 8; i++ {
		var endpoint string
		switch i {
		case 1:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam1.ID)
		case 2:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam1.ID)
		case 3:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam1.ID)
		case 4:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam2.ID)
		case 5:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam2.ID)
		case 6:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam2.ID)
		case 8:
			endpoint = fmt.Sprintf("/teams/%s/players", testTeam3.ID)
		}

		playerPseudo := fmt.Sprintf("killer%d", i)
		params := url.Values{}
		params.Set("pseudo", playerPseudo)
		req, err := http.NewRequest("POST", endpoint, strings.NewReader(params.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/teams/{id}/players", playerCreationHandler)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var player Player
		json.Unmarshal([]byte(rr.Body.String()), &player)

		// player id should be not empty
		if player.ID == "" {
			t.Errorf("handler returned unexpected team name in body: got %v want %v",
				player.ID, "")
		}
		// player pseudo should be the one we passed earlier
		if player.Pseudo != playerPseudo {
			t.Errorf("handler returned unexpected team name in body: got %v want %v",
				player.Pseudo, playerPseudo)
		}

		// Make player available to other tests by making him global
		switch i {
		case 1:
			testPlayer1 = player
		case 2:
			testPlayer2 = player
		case 3:
			testPlayer3 = player
		case 4:
			testPlayer4 = player
		case 5:
			testPlayer5 = player
		case 6:
			testPlayer6 = player
		case 7:
			testPlayer7 = player
		}

	}
}

// TestTeamsListingHandler tests the listing of all teams
func TestTeamsListingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(teamsListingHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("Created teams:")
	fmt.Println(rr.Body.String())
}

// TestTeamDeletionHandler test deletion of a team
func TestTeamDeletionHandler(t *testing.T) {
	fmt.Println(testTeam4)
	// Send the team id of the team we want to delete
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/teams/%s", testTeam4.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", teamDeletionHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
