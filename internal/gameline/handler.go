package gameline

import (
	"fmt"
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

	lines, err := getGameLines(d)
	if err != nil {
		// report error to 3rd party
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rows := buildRows(lines)

	err = templates.ExecuteTemplate(w, "gamelines.tmpl", rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getGameLines(date string) ([]gameLine, error) {
	var gids []int
	var err error

	c := espn.NewScoreboardClient()
	if date == "" {
		gids, err = c.GameIDs()
		if err != nil {
			return nil, fmt.Errorf("getGameLines: %w", err)
		}
	} else {
		gids, err = c.GameIDs(date)
		if err != nil {
			return nil, fmt.Errorf("getGameLines: %w", err)
		}
	}

	var espnLines []espn.GameLine
	for _, gid := range gids {
		gls, err := espn.NewGamecastClient().GameLines(gid)
		if err != nil {
			log.Printf("WARNING: getGameLines: gamecastClient.GameLines: %s\n", err.Error())
			continue
		}

		for _, gl := range gls {
			espnLines = append(espnLines, gl)
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

	return lines, nil
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
