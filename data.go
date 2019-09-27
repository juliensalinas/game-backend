package main

// Team represents a gaming team of players
type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Player represents a game player within a team
type Player struct {
	Team   `json:"team"`
	ID     string `json:"id"`
	Pseudo string `json:"pseudo"`
}

var teams []Team
var players []Player
