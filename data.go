package main

import "time"

// Stats represents the game statistics of a player.
// It can be extended by adding new stats below.
type Stats struct {
	NbAttemptedAttacks       int `json:"nbAttemptedAttacks"`
	NbHits                   int `json:"nbHits"`
	DamageDone               int `json:"damageDone"`
	NbKills                  int `json:"nbKills"`
	NbFirstHitKills          int `json:"nbFirstHitKills"`
	NbAssists                int `json:"nbAssists"`
	NbSpellCasts             int `json:"nbSpellCasts"`
	SpellDamageDone          int `json:"spellDamageDone"`
	TotalTimePlayedInMinutes int `json:"totalTimePlayedInMinutes"`
}

// Player represents a game player within a team
type Player struct {
	ID     string `json:"id"`
	Pseudo string `json:"pseudo"`
	Stats  Stats  `json:"stats"`
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

// RemovePlayer removes a player from the team and reports
// whether player removal was successful or not
func (t *Team) RemovePlayer(id string) (bool, []Player) {
	// Look for the right player in this team and remove him
	for i, p := range t.Players {
		if p.ID == id {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			return true, t.Players
		}
	}
	return false, t.Players
}

// Game represents a game matching 2 teams of equal sizes
// with a limited duration
type Game struct {
	ID        string    `json:"id"`
	Team1     Team      `json:"team1"`
	Team2     Team      `json:"team2"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	StopTime  time.Time `json:"stopTime"`
}

// TeamSizesAreValid checks that game teams have the right size (3 to 5 players)
// and both the same size
func (g *Game) TeamSizesAreValid() bool {
	// TODO(juliensalinas): switch 1 to 3
	if len(g.Team1.Players) == len(g.Team2.Players) && len(g.Team1.Players) <= 5 && len(g.Team1.Players) >= 1 {
		return true
	}
	return false
}

// Stop stops the game by filling in the stop time
func (g *Game) Stop() {
	g.StopTime = time.Now()
}

var teams []Team
var games []Game
