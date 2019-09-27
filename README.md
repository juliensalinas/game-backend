# Description

This is a Go backend for an external game.

In order for this backend to easily interface with any external game developped in any language (in a microservice spirit), this backend is built as a RESTful API.
The following elements can be managed through the API: players, teams, games, achievements, and stats.

Achievements and statistics are admin operations so they cannot be extended through the API but it can be done manually by altering the `achievements.json` and `stats.json` config files.

In order to automatically populate the backend and easily test it, a second program called `driver` is made available.

## API Endpoints Available

### Teams

* `POST /teams` with `name` parameter: create a team by providing a team name, and return the team created
* `DELETE /teams/{id}`: delete a team by providing its team id
* `GET /teams`: list all teams

### Players

* `POST /teams/{id}/players` with `pseudo` parameter: create a player and affect him to a team by providing a pseudo and a team id, and return the player created
* `DELETE /teams/{teamId}/players/{playerId}`: remove a player by providing his id and its team id

### Games

* `POST /games` with `name`, `team1Id` and `team2Id` parameters: create a new game by giving a name and affect 2 teams to this game by providing their team ids, and return the game created
* `PUT /games/{id}` with `teamId` parameter: stop a game by providing the game id and the team id of the winning team
* `GET /games`: list all games

### Achievements

* `GET /games/{gameId}/players/{playerId}/achievements`: list all achievements from a player by providing the player id

### Stats

* `GET /games/{gameId}/players/{playerId}/stats`: list all stats from a player in a game by providing the game id and player id
* `PUT /games/{gameId}/players/{playerId}/stats` with `name` parameter: increment by 1 the stat of a player in a game by providing the stat name (choices are: `nbAttemptedAttacks`, `nbHits`, `damageDone`, `nbKills`, `nbFirstHitKills`, `nbAssists`, `nbSpellCasts`, `spellDamageDone`)

## Backend Usage

1. get the `osmo_test` binary file sent by email
1. launch it. For example on Linux open the terminal and launch `./osmo_test` in console.
1. the backend is now running at the following address `http://127.0.0.1:8000`, and can be consumed through curl, Postman, an internet browser...

## Driver Usage

## TODO

Implement the possibility for a player to move to another team or play another game (which would reinitialize his personal stats except the total time played, but the player stats inside a game would be kept).
Check that the same user is not in the 2 teams at the same time during a game.
Check that the same user is not playing 2 games at the same time.
