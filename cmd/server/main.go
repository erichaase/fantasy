package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

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

	lines := buildGameLines(date)
	data := buildTemplateData(lines)

	err := templates.ExecuteTemplate(w, "game_lines.tmpl", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func buildGameLines(date string) []fantasy.GameLine {
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

	return lines
}

type templateData struct {
	Players []player
}

type player struct {
	Line  fantasy.GameLine
	Class string
}

func buildTemplateData(lines []fantasy.GameLine) templateData {
	playerIds := os.Getenv("PLAYER_IDS")
	m := make(map[int]bool)
	for _, pid := range strings.Split(playerIds, ",") {
		p, _ := strconv.Atoi(pid)
		m[p] = true
	}

	var players []player
	for _, line := range lines {
		p := player{Line: line}
		if m[line.EspnId] {
			p.Class = "team"
		} else if line.Zsum >= 0.0 {
			p.Class = "good"
		} else {
			p.Class = "bad"
		}
		players = append(players, p)
	}

	return templateData{Players: players}
}
