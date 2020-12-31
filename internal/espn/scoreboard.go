package espn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type scoreboardClient struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func NewScoreboardClient(params ...interface{}) scoreboardClient {
	var c *http.Client
	var u *url.URL

	if len(params) == 2 {
		c = params[0].(*http.Client)
		u = params[1].(*url.URL)
	} else {
		c = &http.Client{Timeout: 10 * time.Second}
		u = &url.URL{Scheme: "http", Host: "site.api.espn.com"}
	}

	return scoreboardClient{httpClient: c, baseURL: u}
}

func (c scoreboardClient) GameIDs(params ...string) ([]int, error) {
	c.baseURL.Path = "apis/site/v2/sports/basketball/nba/scoreboard"

	qs := "lang=en&region=us&calendartype=blacklist&limit=100"
	if len(params) == 1 {
		qs = fmt.Sprintf("%s&dates=%s", qs, params[0])
	}
	c.baseURL.RawQuery = qs

	resp, err := c.httpClient.Get(c.baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("scoreboardClient.GameIDs: http.Client.Get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("scoreboardClient.GameIDs: http.Client.Get: non-successful status code: %d", resp.StatusCode)
	}

	var sb struct {
		Events []struct {
			Id     string
			Status struct {
				Type struct {
					State string // values: "pre", "in", "post"
				}
			}
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&sb)
	if err != nil {
		return nil, fmt.Errorf("scoreboardClient.GameIDs: json.Decoder.Decode: %w", err)
	}

	var ids []int
	for _, e := range sb.Events {
		if e.Status.Type.State != "pre" {
			id, err := strconv.Atoi(e.Id)
			if err != nil {
				log.Printf("WARNING: scoreboardClient.GameIDs: strconv.Atoi: error converting game ID: '%s'\n", e.Id)
				continue
			}
			ids = append(ids, id)
		}
	}

	return ids, nil
}
