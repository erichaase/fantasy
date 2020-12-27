package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/erichaase/fantasy/internal/espn"
	"github.com/erichaase/fantasy/internal/fantasy"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

var templates = template.Must(template.ParseFiles("web/template/lines.tmpl"))

type LinesData struct {
	Lines []fantasy.GameLine
}

func handler(w http.ResponseWriter, r *http.Request) {
	gids := espn.GameIdsStarted("")

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
		return lines[i].Zsum > lines[j].Zsum
	})

	linesData := &LinesData{
		Lines: lines,
	}
	err := templates.ExecuteTemplate(w, "lines.tmpl", linesData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
