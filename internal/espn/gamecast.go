package espn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

type game struct {
	Gamecast gamecast
}

type gamecast struct {
	Stats stats
}

type stats struct {
	Player player
}

type player struct {
	Home []GameLine
	Away []GameLine
}

type GameLine struct {
	Id             int
	FirstName      string
	LastName       string
	PositionAbbrev string
	Jersey         string
	Active         string
	IsStarter      string
	Fg             string
	Ft             string
	Threept        string
	Rebounds       string
	Assists        string
	Steals         string
	Fouls          string
	Points         string
	Minutes        string
	Blocks         string
	Turnovers      string
	PlusMinus      string
	Dnp            bool
	EnteredGame    bool
}

// convert into api client
func GameLines(gid int) []GameLine {
	query := fmt.Sprintf("xhr=1&gameId=%d&lang=en&init=true&setType=true&confId=null", gid)
	u := &url.URL{
		Scheme:   "http",
		Host:     "scores.espn.go.com",
		Path:     "nba/gamecast12/master",
		RawQuery: query,
	}

	// var myClient = &http.Client{Timeout: 10 * time.Second}
	resp, err := http.Get(u.String())
	// check for non successful http status
	if err != nil { // shorthand to make this terse
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	s := string(body)
	re := regexp.MustCompile("[[:^print:]]")
	t := re.ReplaceAllLiteralString(s, "")
	b := []byte(t)

	g := game{}
	err = json.Unmarshal(b, &g)
	if err != nil {
		log.Fatalln(err)
	}

	var egls []GameLine
	for _, gl := range g.Gamecast.Stats.Player.Home {
		if gl.Id != 0 { // totals row has a 0 id
			egls = append(egls, gl)
		}
	}
	for _, gl := range g.Gamecast.Stats.Player.Away {
		if gl.Id != 0 { // totals row has a 0 id
			egls = append(egls, gl)
		}
	}

	return egls
}
