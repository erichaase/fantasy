package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	// ids := get_started_game_ids()
	// fmt.Println(ids)
	print_game_stats(401266806)
}

type scoreboard struct {
	Events []event
}

type event struct {
	Id string
	Status status
}

type status struct {
	Type statusType
}

type statusType struct {
	State string // values: "pre", "in", "post"
}

// move into espn scoreboard api client
func get_started_game_ids() []int {
	u := &url.URL{
		Scheme: "http",
		Host: "site.api.espn.com",
		Path: "apis/site/v2/sports/basketball/nba/scoreboard",
		RawQuery: "lang=en&region=us&calendartype=blacklist&limit=100",
	}

	// var myClient = &http.Client{Timeout: 10 * time.Second}
	resp, err := http.Get(u.String())
	// check for non successful http status
	if err != nil { // shorthand to make this terse
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// json.NewDecoder(r.Body).Decode(target)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		// exit?
	}

	sb := scoreboard{}
	err = json.Unmarshal(body, &sb)
	if err != nil {
		log.Fatalln(err)
		// exit?
	}

	var ids []int
	for _, e := range sb.Events {
		fmt.Printf("%s: %s\n", e.Id, e.Status.Type.State)
		if e.Status.Type.State != "pre" {
			id, err := strconv.Atoi(e.Id)
			if err != nil {
				log.Fatalln(err)
				// exit?
			}
			ids = append(ids, id)
		}
	}

	return ids
}

type game struct {
	Gamecast gamecast
}

type gamecast struct {
	Stats stats
}

type stats struct {
	Player players
}

type players struct {
	Home []player
	Away []player
}

type player struct {
	Id int
	FirstName string
	LastName string
	PositionAbbrev string
	Jersey string
	Active string
	IsStarter string
	Fg string
	Ft string
	Threept string
	Rebounds string
	Assists string
	Steals string
	Fouls string
	Points string
	Minutes string
	Blocks string
	Turnovers string
	PlusMinus string
	Dnp bool
	EnteredGame bool
}

func print_game_stats(gid int) {
	query := fmt.Sprintf("xhr=1&gameId=%d&lang=en&init=true&setType=true&confId=null", gid)
	u := &url.URL{
		Scheme: "http",
		Host: "scores.espn.go.com",
		Path: "nba/gamecast12/master",
		RawQuery: query,
	}

	// var myClient = &http.Client{Timeout: 10 * time.Second}
	resp, err := http.Get(u.String())
	// check for non successful http status
	if err != nil { // shorthand to make this terse
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// json.NewDecoder(r.Body).Decode(target)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		// exit?
	}

	g := game{}
	err = json.Unmarshal(body, &g)
	if err != nil {
		log.Fatalln(err)
		// exit?
	}

	// skip totals player row
	for _, p := range g.Gamecast.Stats.Player.Home {
		fmt.Printf("%s %s: %s\n", p.FirstName, p.LastName, p.Points)
	}

	for _, p := range g.Gamecast.Stats.Player.Away {
		fmt.Printf("%s %s: %s\n", p.FirstName, p.LastName, p.Points)
	}

    // build new player struct for my use
}