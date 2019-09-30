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
var team1 Team

func TestTeamCreationHandler(t *testing.T) {
	// POST param to pass
	teamName := "Best Team Ever"

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
	json.Unmarshal([]byte(rr.Body.String()), &team1)

	// Team id should be not empty
	if team1.ID == "" {
		t.Errorf("handler returned unexpected team name in body: got %v want %v",
			team1.ID, "")
	}
	// Team name should be the one we passed earlier
	if team1.Name != teamName {
		t.Errorf("handler returned unexpected team name in body: got %v want %v",
			team1.Name, teamName)
	}
	// Team players should be empty
	if team1.Players != nil {
		t.Errorf("handler returned unexpected team name in body: got %v want %v",
			team1.Players, nil)
	}
}

func TestTeamDeletionHandler(t *testing.T) {
	// Delete the newly created team.
	//
	// Send the team id of the team we want to delete
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/teams/%s", team1.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", teamDeletionHandler)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
