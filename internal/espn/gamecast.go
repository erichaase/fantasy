package espn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type gamecastClient struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func NewGamecastClient(params ...interface{}) gamecastClient {
	var c *http.Client
	var u *url.URL

	if len(params) == 2 {
		c = params[0].(*http.Client)
		u = params[1].(*url.URL)
	} else {
		c = &http.Client{Timeout: 10 * time.Second}
		u = &url.URL{Scheme: "http", Host: "scores.espn.go.com"}
	}

	return gamecastClient{httpClient: c, baseURL: u}
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

func (c gamecastClient) GameLines(gid int) ([]GameLine, error) {
	c.baseURL.Path = "nba/gamecast12/master"
	c.baseURL.RawQuery = fmt.Sprintf("xhr=1&gameId=%d&lang=en&init=true&setType=true&confId=null", gid)

	resp, err := c.httpClient.Get(c.baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("gamecastClient.GameLines: http.Client.Get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("gamecastClient.GameLines: http.Client.Get: non-successful status code: %d", resp.StatusCode)
	}

	// should use json.Decode() here but the payload includes unicode chars which we need to strip, see below for details
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gamecastClient.GameLines: ioutil.ReadAll: %w", err)
	}

	s := string(body)
	t := regexp.MustCompile("[[:^print:]]").ReplaceAllLiteralString(s, "")
	b := []byte(t)

	var g struct {
		Gamecast struct {
			Stats struct {
				Player struct {
					Home []GameLine
					Away []GameLine
				}
			}
		}
	}

	err = json.Unmarshal(b, &g)
	if err != nil {
		return nil, fmt.Errorf("gamecastClient.GameLines: json.Unmarshal: %w", err)
	}

	var gls []GameLine
	for _, gl := range g.Gamecast.Stats.Player.Home {
		if gl.Id != 0 { // totals row has a 0 id
			gls = append(gls, gl)
		}
	}
	for _, gl := range g.Gamecast.Stats.Player.Away {
		if gl.Id != 0 { // totals row has a 0 id
			gls = append(gls, gl)
		}
	}

	return gls, nil
}
