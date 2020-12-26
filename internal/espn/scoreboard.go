package espn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type scoreboard struct {
	Events []event
}

type event struct {
	Id     string
	Status status
}

type status struct {
	Type statusType
}

type statusType struct {
	State string // values: "pre", "in", "post"
}

// convert into api client
func GameIdsStarted(day string) []int {
	q := "lang=en&region=us&calendartype=blacklist&limit=100"
	if day != "" {
		q = fmt.Sprintf("%s&dates=%s", q, day)
	}

	u := &url.URL{
		Scheme:   "http",
		Host:     "site.api.espn.com",
		Path:     "apis/site/v2/sports/basketball/nba/scoreboard",
		RawQuery: q,
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
	}

	sb := scoreboard{}
	err = json.Unmarshal(body, &sb)
	if err != nil {
		log.Fatalln(err)
	}

	var ids []int
	for _, e := range sb.Events {
		if e.Status.Type.State != "pre" {
			id, err := strconv.Atoi(e.Id)
			if err != nil {
				log.Fatalln(err)
			}
			ids = append(ids, id)
		}
	}

	return ids
}
