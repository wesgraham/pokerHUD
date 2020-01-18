# PokerHUD
[![Go Report](https://goreportcard.com/badge/github.com/FourthState/plasma-mvp-sidechain)](https://goreportcard.com/report/github.com/wesgraham/pokerHUD)
[![license](https://img.shields.io/aur/license/android-studio)](https://github.com/wesgraham/pokerHUD/blob/master/LICENSE)

## Project Status
This is meant to be an academic project. There is very little development occuring for this project presently. This is due to hurdles encountered in bot policy on most active poker sites. I will continue to maintain this repository by thoroughly reviewing any open source contributions. I will provide support and guidance for anyone looking to continue development. 

## What is PokerHUD?
PokerHUD is a poker heads up display (HUD). The HUD allows poker players to enter hand data on their opponents, and displays vital statistics such as **VPiP** (% of hands with funds voluntarily put in pot), **PFR** (Pre-flop aggressor rate), and **average bet siz**e. These metrics can be used to analyze a player's aggression and playing habits. 

PokerHUD consists of a Postgres instance, a set of Go functions, and a MVP front end. 

## Getting Started (MacOS)

**Requirements**
- [golang v1.11+](https://golang.org/)
- [PostgreSQL 9.4+](https://www.postgresql.org/download/)
- [PostgREST 6.0.2+](http://postgrest.org/en/v6.0/tutorials/tut0.html)

Ensure that all of the above are installed correctly before going any further. Note that PostgREST exists as a single binary, which should be present in your working directory.

**Instantiating Table In Postgres**
```
CREATE DATABASE pokerHands;
CREATE SCHEMA player;
CREATE TABLE player.hands (uname text, handid integer, balance integer, hand text, potsize integer, action text, amount integer, board text[], threebet boolean, fourplusbet boolean);
```

**Running PostgREST**

`./postgrest pokerhud.conf`

You should see:

```
Listening on port 3000
Attempting to connect to the database...
Connection successful
```

*Note:* Default port is set to :5432, and authenticated user is set to role titled: *authenticator*
Please follow the above postgREST installation tutorial to set up authenticator role before running postgrest server.

**Running the Application**

 `go run /cmd/main.go`

You should see: 
```
Listening on port :8080
```

The application should then be accessible at localhost:8080
