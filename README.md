# Description

This is a Go backend for an external game.

In order for this backend to easily interface with any external game developped in any language (in a microservice spirit), this backend is built as a RESTful API.
The following elements can be managed through the API: players, teams, games, achievements, and stats.

Achievements and statistics cannot be extended through the API but it can be done manually by altering tje `achievements.json` and `stats.json` config files.

In order to automatically populate the backend and easily test it, a second program called `driver` is made available.


## API Endpoints Available

### Teams

* `POST /teams` with `name` parameter: create a team by providing a team name, and returns the team id of the team created
* `DELETE /teams/{id}`: delete a team by providing its team id
* `GET /teams`: list all team names

### Players

* `POST /teams/{id}/players` with `pseudo` parameter: create a player and affect him to a team by providing a pseudo and a team id, and return the player id of the player created
* `DELETE /players/{id}`: remove a player by providing his id
* `GET /players` (GET): list all players from a team by providing a team id

### Games

* `POST /games` with `team1Id` and `team2Id` parameters: create a new game and affect 2 teams to this game by providing their team ids, and returns the game id of the game created
* `DELETE /games/{id}`: stop a game by providing the game id

### Achievements

* `GET /players/{id}/achievements`: list all achievements from a player by providing the player id

### Stats

* `GET /players/{id}/stats`: list all available stats from a player by providing the player id
* `PUT /players/{id}/stats/{id}`: increment by 1 the stat of a player by providing the stat id

## Backend Installation

## Driver Usage

