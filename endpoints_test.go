package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

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
	var team Team
	json.Unmarshal([]byte(rr.Body.String()), &team)
	// Team ID should be not empty
	if team.ID == "" {
		t.Errorf("handler returned unexpected team name in body: got %v want %v",
			team.ID, "")
	}
	// Team name should be the one we passed earlier
	if team.Name != teamName {
		t.Errorf("handler returned unexpected team name in body: got %v want %v",
			team.Name, teamName)
	}
	// Team Plauers should be empty
	if team.Players != nil {
		t.Errorf("handler returned unexpected team name in body: got %v want %v",
			team.Players, nil)
	}
}
