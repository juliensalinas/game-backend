package main

import "time"

// Player represents a game player within a team
type Player struct {
	ID     string `json:"id"`
	Pseudo string `json:"pseudo"`
}

// Team represents a gaming team of players
type Team struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

// AddPlayer adds a player to the team
func (t *Team) AddPlayer(p Player) []Player {
	t.Players = append(t.Players, p)
	return t.Players
}

// Game represents a game matching 2 teams of equal sizes
// with a limited duration
type Game struct {
	Team1     Team      `json: "team1"`
	Team2     Team      `json: "team2"`
	Name      string    `json:"name"`
	StartTime time.Time `json: "startTime"`
	StopTime  time.Time `json: "stopTime"`
}

// TeamSizesAreValid checks that game teams have the right size (3 to 5 players)
// and both the same size
func (g *Game) TeamSizesAreValid() bool {
	if len(g.Team1.Players) == len(g.Team2.Players) && len(g.Team1.Players) <= 5 && len(g.Team1.Players) >= 3 {
		return true
	}
	return false
}

var teams []Team
var games []Game
