package main

import (
	"os"
	"sort"

	"github.com/erichaase/fantasy/internal/espn"
	"github.com/erichaase/fantasy/internal/fantasy"
)

func main() {
	date := ""
	if len(os.Args) > 1 {
		date = os.Args[1]
	}
	gids := espn.GameIdsStarted(date)

	var espnLines []espn.GameLine
	for _, gid := range gids {
		for _, line := range espn.GameLines(gid) {
			espnLines = append(espnLines, line)
		}
	}

	var lines []fantasy.GameLine
	for _, espnLine := range espnLines {
		line := fantasy.BuildGameLine(espnLine)
		lines = append(lines, line)
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].Zsum < lines[j].Zsum
	})

	for _, l := range lines {
		fantasy.PrintGameLine(l)
	}
}
