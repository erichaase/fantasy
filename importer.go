package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	day := os.Args[1]
	gids := getStartedGameIds(day)

	var glss []GameLine
	for _, gid := range gids {
		egls := getEspnGameLines(gid)
		gls := mapEspnGameLines(egls)
		for _, gl := range gls {
			glss = append(glss, gl)
		}
	}

	sort.SliceStable(glss, func(i, j int) bool {
		return glss[i].Zsum < glss[j].Zsum
	})

	for _, gl := range glss {
		printGameLine(gl)
	}
}

//######################### ESPN Scoreboard API Client #########################

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
func getStartedGameIds(day string) []int {
	q := "lang=en&region=us&calendartype=blacklist&limit=100"
	if day != "" {
		q = fmt.Sprintf("%s&dates=%s", q, day)
	}

	u := &url.URL{
		Scheme: "http",
		Host: "site.api.espn.com",
		Path: "apis/site/v2/sports/basketball/nba/scoreboard",
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

//########################## ESPN Gamecast API Client ##########################

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
	Home []EspnGameLine
	Away []EspnGameLine
}

type EspnGameLine struct {
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

func getEspnGameLines(gid int) []EspnGameLine {
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

	var egls []EspnGameLine
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

//################################ Some Library ################################

type GameLine struct {
	EspnId int
	FirstName string
	LastName string
	Min int
	Fgm int
	Fga int
	Ftm int
	Fta int
	Tpm int
	Tpa int
	Pts int
	Reb int
	Ast int
	Stl int
	Blk int
	To int
	Zfg float64
	Zft float64
	Ztp float64
	Zpts float64
	Zreb float64
	Zast float64
	Zstl float64
	Zblk float64
	Zto float64
	Zsum float64
}

func mapEspnGameLines(egls []EspnGameLine) []GameLine {
	var gls []GameLine
	for _, egl := range egls {
		gl := mapEspnGameLine(egl)
		gls = append(gls, gl)
	}
	return gls
}

func mapEspnGameLine(egl EspnGameLine) GameLine {
	min, _ := strconv.Atoi(egl.Minutes)
	pts, _ := strconv.Atoi(egl.Points)
	reb, _ := strconv.Atoi(egl.Rebounds)
	ast, _ := strconv.Atoi(egl.Assists)
	stl, _ := strconv.Atoi(egl.Steals)
	blk, _ := strconv.Atoi(egl.Blocks)
	to, _ := strconv.Atoi(egl.Turnovers)

	var fgm, fga int
	fmt.Sscanf(egl.Fg, "%d/%d", &fgm, &fga)
	var ftm, fta int
	fmt.Sscanf(egl.Ft, "%d/%d", &ftm, &fta)
	var tpm, tpa int
	fmt.Sscanf(egl.Threept, "%d/%d", &tpm, &tpa)

	// better way to calculate this?
	// double check the following using hashtag
	zfg := 0.0
	if fga != 0 {
		zfg = (((float64(fgm) / float64(fga)) - 0.457) / 0.06236986313) * (float64(fga) / (5.011649158 * 2))
	}
	zft := 0.0
	if fta != 0 {
		zft = (((float64(ftm) / float64(fta)) - 0.775) / 0.08132524135) * (float64(fta) / (1.819730344 * 2))
	}
	ztp := (float64(tpm) - 1.5) / 0.9
	zpts := (float64(pts) - 13.1) / 6.35
	zreb := (float64(reb) - 4.8) / 2.4
	zast := (float64(ast) - 2.5) / 1.8
	zstl := (float64(stl) - 0.9) / 0.4
	zblk := (float64(blk) - 0.5) / 0.5
	zto := -((float64(to) - 1.55) / 0.85)
	zsum := zfg + zft + ztp + zpts + zreb + zast + zstl + zblk + zto

	return GameLine{
		EspnId: egl.Id,
		FirstName: egl.FirstName,
		LastName: egl.LastName,
		Min: min,
		Fgm: fgm,
		Fga: fga,
		Ftm: ftm,
		Fta: fta,
		Tpm: tpm,
		Tpa: tpa,
		Pts: pts,
		Reb: reb,
		Ast: ast,
		Stl: stl,
		Blk: blk,
		To: to,
		Zfg: zfg,
		Zft: zft,
		Ztp: ztp,
		Zpts: zpts,
		Zreb: zreb,
		Zast: zast,
		Zstl: zstl,
		Zblk: zblk,
		Zto: zto,
		Zsum: zsum,
	}
}

func printGameLine(l GameLine) {
	fmt.Printf("%s %s,|,%dm,%d-%d,%d-%d,%d-%d,%d-%d-%d,%d-%d-%d,|,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,|,%.1f\n",
		l.FirstName,
		l.LastName,
		l.Min,
		l.Fgm,
		l.Fga,
		l.Ftm,
		l.Fta,
		l.Tpm,
		l.Tpa,
		l.Pts,
		l.Reb,
		l.Ast,
		l.Stl,
		l.Blk,
		l.To,
		l.Zfg,
		l.Zft,
		l.Ztp,
		l.Zpts,
		l.Zreb,
		l.Zast,
		l.Zstl,
		l.Zblk,
		l.Zto,
		l.Zsum,
	)
}