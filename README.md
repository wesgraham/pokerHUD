## Project Status
This is meant to be an academic project. There is very little development occuring for this project presently. This is due to hurdles encountered in bot policy on most active poker sites. I will continue to maintain this repository by thoroughly reviewing any open source contributions. I will provide support and guidance for anyone looking to continue development. 

## What is PokerHUD?
PokerHUD is a poker heads up display (HUD). The HUD allows poker players to enter hand data on their opponents, and displays vital statistics such as **VPiP** (% of hands with funds voluntarily put in pot), **PFR** (Pre-flop aggressor rate), and **average bet siz**e. These metrics can be used to analyze a player's aggression and playing habits. 

PokerHUD consists of a Postgres instance, a set of Go functions, and a MVP front end. 

## Getting Started (MacOS)

**Requirements**
- [golang v1.11+](https://golang.org/)
- [PostgreSQL 9.4+](https://www.postgresql.org/download/)
- [PostgREST 6.0.2+] (http://postgrest.org/en/v6.0/tutorials/tut0.html)

Ensure that all of the above are installed correctly before going any further. Note that PostgREST exists as a single binary, which should be present in your working directory.

**Running Postgrest**

`./postgrest pokerhud.conf`

You should see:
`Listening on port 3000
Attempting to connect to the database...
Connection successful`

*Note:* Default port is set to :5432
Please follow the above postgREST installation tutorial to set up authenticator role before running postgrest server.

**Running the Application**
go run /cmd/main.go

You should see: 
`Connection successful
Listening on port :8080
`
