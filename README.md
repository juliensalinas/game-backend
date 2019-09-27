# Description

This is a Go backend for an external game.

In order for this backend to easily interface with any external game developped in any language (in a microservice spirit), this backend is built as a RESTful API.
The following elements can be managed through the API: players, teams, games, achievements, and stats.

Achievements and statistics cannot be extended through the API but it can be done manually by altering tje `achievements.json` and `stats.json` config files.

In order to automatically populate the backend and easily test it, a second program called `driver` is made available.

## API Endpoints Available

### Teams

* `POST /teams` with `name` parameter: create a team by providing a team name, and return the team created
* `DELETE /teams/{id}`: delete a team by providing its team id
* `GET /teams`: list all teams

### Players

* `POST /teams/{id}/players` with `pseudo` parameter: create a player and affect him to a team by providing a pseudo and a team id, and return the player created
* `DELETE /teams/{id}/players/{id}`: remove a player by providing his id and its team id

### Games

* `POST /games` with `name`, `team1Id` and `team2Id` parameters: create a new game by giving a name and affect 2 teams to this game by providing their team ids, and return the game created
* `DELETE /games/{id}`: stop a game by providing the game id
* `GET /games`: list all games

### Achievements

* `GET /players/{id}/achievements`: list all achievements from a player by providing the player id

### Stats

* `GET /players/{id}/stats`: list all available stats from a player by providing the player id
* `PUT /players/{id}/stats/{id}`: increment by 1 the stat of a player by providing the stat id

## Backend Usage

1. get the `osmo_test` binary file sent by email
1. launch it. For example on Linux open the terminal and launch `./osmo_test` in console.
1. the backend is now running at the following address `http://127.0.0.1:8000`, and can be consumed through curl, Postman, an internet browser...

## Driver Usage

