package main

// Team represents a gaming team of players
type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var teams = []Team{
	{
		ID:   "1",
		Name: "Amazing team",
	},
	{
		ID:   "2",
		Name: "Amazing team 2",
	},
}
