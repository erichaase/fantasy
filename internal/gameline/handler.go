package gameline

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/erichaase/fantasy/internal/espn"
)

type row struct {
	Line  gameLine
	Class string
}

var templates = template.Must(template.ParseFiles("web/template/gamelines.tmpl"))

func LinesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	q, _ := url.ParseQuery(r.URL.RawQuery)
	d := q.Get("date")

	lines := getGameLines(d)
	rows := buildRows(lines)

	err := templates.ExecuteTemplate(w, "gamelines.tmpl", rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getGameLines(date string) []gameLine {
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

	var lines []gameLine
	for _, espnLine := range espnLines {
		line := newGameLineFromEspn(espnLine)
		lines = append(lines, line)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].Zsum > lines[j].Zsum
	})

	return lines
}

func buildRows(gameLines []gameLine) []row {
	pids := os.Getenv("PLAYER_IDS")
	m := make(map[int]bool)
	for _, pid := range strings.Split(pids, ",") {
		p, _ := strconv.Atoi(pid)
		m[p] = true
	}

	var rows []row
	for _, gl := range gameLines {
		r := row{Line: gl}
		if m[gl.EspnId] {
			r.Class = "team"
		} else if gl.Zsum >= 0.0 {
			r.Class = "good"
		} else {
			r.Class = "bad"
		}
		rows = append(rows, r)
	}

	return rows
}
