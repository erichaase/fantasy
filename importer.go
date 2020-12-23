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
	ids := get_game_ids()
	fmt.Println(ids)
}

type statusType struct {
	State string // values: "pre", "in", "post"
}

type status struct {
	Type statusType
}

type event struct {
	Id string
	Status status
}

type scoreboard struct {
	Events []event
}

// move into espn scoreboard api client
func get_game_ids() []int {
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