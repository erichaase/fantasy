package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"

	"github.com/erichaase/fantasy/internal/espn"
	"github.com/erichaase/fantasy/internal/fantasy"
)

func main() {
	lh := flag.Bool("localhost", false, "Configure HTTP server to listen to localhost for local development")
	port := flag.Int("port", 80, "The port that the HTTP server listens to for incoming requests")
	flag.Parse()

	addr := ""
	if *lh {
		addr = "localhost"
	}
	ap := fmt.Sprintf("%s:%d", addr, *port)
	log.Printf("Listening on '%s'", ap)

	http.HandleFunc("/lines", handler)
	log.Fatal(http.ListenAndServe(ap, nil))
}

var templates = template.Must(template.ParseFiles("web/template/game_lines.tmpl"))

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	q, _ := url.ParseQuery(r.URL.RawQuery)
	date := q.Get("date")

	var gids []int
	if date == "" {
		gids = espn.GameIdsStarted()
	} else {
		gids = espn.GameIdsStarted(date)
	}

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
