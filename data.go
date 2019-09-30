package main

import (
	"time"
)

// AchievementsCalculator is an interface for player achievements calculation
type AchievementsCalculator interface {
	Calculate(Stats)
}

// Achievements represents which achievements have been
// reached by a player
type Achievements struct {
	Sharpshooter bool `json:"sharpshooter"`
	Bruiser      bool `json:"bruiser"`
	Veteran      bool `json:"veteran"`
	BigWinner    bool `json:"bigWinner"`
}

// CalculateAchievements calculates the achievements of a player
func (a *Achievements) CalculateAchievements(stats Stats) {
	if stats.NbHits != 0 && float64(stats.NbHits/stats.NbAttemptedAttacks) >= 0.75 {
		a.Sharpshooter = true
	}
	if stats.DamageDone+stats.SpellDamageDone >= 500 {
		a.Bruiser = true
	}
	if stats.TotalNbGamesPlayed >= 1000 {
		a.Veteran = true
	}
	if stats.TotalNbWins >= 200 {
		a.BigWinner = true
	}
}

// StatsIncrementer is an interface for player stat
// incrementation
type StatsIncrementer interface {
	IncrementStats(statName string) bool
}

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
	TotalTimePlayedInSeconds int `json:"totalTimePlayedInSeconds"`
	TotalNbGamesPlayed       int `json:"totalNbGamesPlayed"`
	TotalNbWins              int `json:"totalNbGamesWins"`
}

// CalculateGlobalStats calculates the global stats that a player
// is accumulating for his whole life
func (s *Stats) CalculateGlobalStats(gameDuration int) {
	s.TotalTimePlayedInSeconds = s.TotalTimePlayedInSeconds + gameDuration
	s.TotalNbGamesPlayed++
}

// IncrementStats increments one of the player stats based on
// the stat name provided
func (s *Stats) IncrementStats(statName string) bool {
	switch statName {
	case "nbAttemptedAttacks":
		s.NbAttemptedAttacks++
	case "nbHits":
		s.NbHits++
	case "damageDone":
		s.DamageDone++
	case "nbKills":
		s.NbKills++
	case "nbFirstHitKills":
		s.NbFirstHitKills++
	case "nbAssists":
		s.NbAssists++
	case "nbSpellCasts":
		s.NbSpellCasts++
	case "spellDamageDone":
		s.SpellDamageDone++
	default:
		return false
	}
	return true
}

// Player represents a game player within a team
type Player struct {
	ID           string       `json:"id"`
	Pseudo       string       `json:"pseudo"`
	Stats        Stats        `json:"stats"`
	Achievements Achievements `json:"achievements"`
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

// MarkAsWinner increments all the TotalNbWins of the players
// of the team
func (t *Team) MarkAsWinner() []Player {
	for _, p := range t.Players {
		p.Stats.TotalNbWins++
	}
	return t.Players
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
	if len(g.Team1.Players) == len(g.Team2.Players) && len(g.Team1.Players) <= 5 && len(g.Team1.Players) >= 3 {
		return true
	}
	return false
}

// Stop stops the game by filling in the stop time and computes the duration
// in seconds.
// It also updates all the players' TotalTimePlayedInSeconds, TotalNbGamesPlayed
// and TotalNbWins stat.
// It also calculates all the players achievements.
func (g *Game) Stop() {
	// Fill the stop time
	g.StopTime = time.Now()

	// Compute the duration and convert it to seconds
	gameDuration := int(g.StopTime.Sub(g.StartTime) / time.Second)

	// Update all the players TotalTimePlayedInSeconds and TotalNbGamesPlayed stats
	// and calculate their achievements
	var players1, players2 []Player
	for _, p := range g.Team1.Players {
		p.Stats.CalculateGlobalStats(gameDuration)
		p.Achievements.CalculateAchievements(p.Stats)
		players1 = append(players1, p)
	}
	g.Team1.Players = players1
	for _, p := range g.Team2.Players {
		p.Stats.CalculateGlobalStats(gameDuration)
		p.Achievements.CalculateAchievements(p.Stats)
		players2 = append(players2, p)
	}
	g.Team2.Players = players2
}

// Init current list of teams and list of games
var teams []Team
var games []Game
