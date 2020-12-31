package main

import (
	"log"
)

func main() {
	log.Printf("importing")

	// parse args / flags, date
	// gameline.Import()
}

/*
espnScoreboardClient := espn.NewScoreboardClient()
gids := espnScoreboardClient.GameIdsStarted()

for _, gid := gids {
	espnGamecastClient := espn.NewGamecastClient()
	egls := espnGamecastClient.GameLines(gid)
	for _, egl := egls {
		gl := gameline.NewGameLineFromEspn(egl)
		go gameline.Upsert(gl) // don't persist zscores
	}
}
*/
