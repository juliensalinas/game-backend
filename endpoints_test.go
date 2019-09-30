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
var team1, team2, team3, team4 Team

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

		// Make team available to other tests by making them globals
		switch i {
		case 1:
			team1 = team
		case 2:
			team2 = team
		case 3:
			team3 = team
		case 4:
			team4 = team
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
	fmt.Println(team4)
	// Send the team id of the team we want to delete
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/teams/%s", team4.ID), nil)
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
