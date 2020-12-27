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

var templates = template.Must(template.ParseFiles("web/template/game_lines.tmpl"))

func handler(w http.ResponseWriter, r *http.Request) {
	gids := espn.GameIdsStarted("20201225")

	var espnLines []espn.GameLine
	for _, gid := range gids {
		for _, line := range espn.GameLines(gid) {
			espnLines = append(espnLines, line)
		}
	}

	var lines []fantasy.GameLine
	for _, espnLine := range espnLines {
		line := fantasy.NewGameLineFromEspn(espnLine)
		lines = append(lines, line)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].Zsum > lines[j].Zsum
	})

	data := struct{ Lines []fantasy.GameLine }{lines}
	err := templates.ExecuteTemplate(w, "game_lines.tmpl", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
