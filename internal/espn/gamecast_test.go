package espn_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/erichaase/fantasy/internal/espn"
)

func TestGamecastClient_GameLines(t *testing.T) {
	// Given
	tdPath := filepath.Join("testdata", "gamecast-401266799.json")
	td, err := ioutil.ReadFile(tdPath)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(td)
	}))
	defer ts.Close()

	url, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// When
	c := espn.NewGamecastClient(ts.Client(), url)
	gls, err := c.GameLines(401266799)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Then
	if !reflect.DeepEqual(gls, expectedGameLines) {
		t.Errorf("Expected: %v, Got: %v", expectedGameLines, gls)
	}
}

var expectedGameLines = []espn.GameLine{
	{Id: 2581018, FirstName: "Kentavious", LastName: "Caldwell-Pope", PositionAbbrev: "SG", Jersey: "1", Active: "false", IsStarter: "true", Fg: "3/7", Ft: "0/0", Threept: "2/6", Rebounds: "0", Assists: "2", Steals: "1", Fouls: "4", Points: "8", Minutes: "24", Blocks: "2", Turnovers: "0", PlusMinus: "+14", Dnp: false, EnteredGame: true},
	{Id: 3032979, FirstName: "Dennis", LastName: "Schroder", PositionAbbrev: "PG", Jersey: "17", Active: "false", IsStarter: "true", Fg: "7/11", Ft: "3/4", Threept: "1/3", Rebounds: "2", Assists: "6", Steals: "0", Fouls: "5", Points: "18", Minutes: "25", Blocks: "0", Turnovers: "4", PlusMinus: "+8", Dnp: false, EnteredGame: true},
	{Id: 1966, FirstName: "LeBron", LastName: "James", PositionAbbrev: "SF", Jersey: "23", Active: "false", IsStarter: "true", Fg: "8/18", Ft: "3/4", Threept: "3/8", Rebounds: "7", Assists: "10", Steals: "1", Fouls: "0", Points: "22", Minutes: "31", Blocks: "0", Turnovers: "4", PlusMinus: "+16", Dnp: false, EnteredGame: true},
	{Id: 6583, FirstName: "Anthony", LastName: "Davis", PositionAbbrev: "PF", Jersey: "3", Active: "false", IsStarter: "true", Fg: "10/16", Ft: "5/7", Threept: "3/5", Rebounds: "8", Assists: "5", Steals: "2", Fouls: "2", Points: "28", Minutes: "30", Blocks: "0", Turnovers: "2", PlusMinus: "+16", Dnp: false, EnteredGame: true},
	{Id: 3206, FirstName: "Marc", LastName: "Gasol", PositionAbbrev: "C", Jersey: "14", Active: "false", IsStarter: "true", Fg: "0/1", Ft: "2/2", Threept: "0/1", Rebounds: "9", Assists: "0", Steals: "0", Fouls: "2", Points: "2", Minutes: "20", Blocks: "0", Turnovers: "1", PlusMinus: "+8", Dnp: false, EnteredGame: true},
	{Id: 3201, FirstName: "Jared", LastName: "Dudley", PositionAbbrev: "SF", Jersey: "10", Active: "true", IsStarter: "false", Fg: "0/0", Ft: "0/0", Threept: "0/0", Rebounds: "3", Assists: "0", Steals: "0", Fouls: "0", Points: "0", Minutes: "2", Blocks: "0", Turnovers: "0", PlusMinus: "+4", Dnp: false, EnteredGame: true},
	{Id: 4032, FirstName: "Wesley", LastName: "Matthews", PositionAbbrev: "SG", Jersey: "9", Active: "false", IsStarter: "false", Fg: "1/3", Ft: "0/0", Threept: "0/2", Rebounds: "1", Assists: "0", Steals: "1", Fouls: "1", Points: "2", Minutes: "16", Blocks: "0", Turnovers: "0", PlusMinus: "-3", Dnp: false, EnteredGame: true},
	{Id: 6461, FirstName: "Markieff", LastName: "Morris", PositionAbbrev: "PF", Jersey: "88", Active: "false", IsStarter: "false", Fg: "3/5", Ft: "0/0", Threept: "3/5", Rebounds: "7", Assists: "2", Steals: "1", Fouls: "0", Points: "9", Minutes: "15", Blocks: "0", Turnovers: "0", PlusMinus: "+3", Dnp: false, EnteredGame: true},
	{Id: 2530923, FirstName: "Alfonzo", LastName: "McKinnie", PositionAbbrev: "SF", Jersey: "28", Active: "true", IsStarter: "false", Fg: "1/1", Ft: "0/0", Threept: "1/1", Rebounds: "1", Assists: "0", Steals: "0", Fouls: "0", Points: "3", Minutes: "2", Blocks: "0", Turnovers: "0", PlusMinus: "+4", Dnp: false, EnteredGame: true},
	{Id: 3134907, FirstName: "Kyle", LastName: "Kuzma", PositionAbbrev: "PF", Jersey: "0", Active: "false", IsStarter: "false", Fg: "5/11", Ft: "0/0", Threept: "3/4", Rebounds: "6", Assists: "2", Steals: "0", Fouls: "1", Points: "13", Minutes: "22", Blocks: "1", Turnovers: "0", PlusMinus: "+7", Dnp: false, EnteredGame: true},
	{Id: 4396991, FirstName: "Talen", LastName: "Horton-Tucker", PositionAbbrev: "SG", Jersey: "5", Active: "true", IsStarter: "false", Fg: "1/2", Ft: "2/2", Threept: "1/1", Rebounds: "1", Assists: "2", Steals: "0", Fouls: "2", Points: "5", Minutes: "7", Blocks: "1", Turnovers: "1", PlusMinus: "+9", Dnp: false, EnteredGame: true},
	{Id: 2991055, FirstName: "Montrezl", LastName: "Harrell", PositionAbbrev: "PF", Jersey: "15", Active: "true", IsStarter: "false", Fg: "10/13", Ft: "2/5", Threept: "0/1", Rebounds: "7", Assists: "2", Steals: "1", Fouls: "4", Points: "22", Minutes: "28", Blocks: "0", Turnovers: "2", PlusMinus: "+15", Dnp: false, EnteredGame: true},
	{Id: 2566745, FirstName: "Quinn", LastName: "Cook", PositionAbbrev: "PG", Jersey: "2", Active: "true", IsStarter: "false", Fg: "0/1", Ft: "0/0", Threept: "0/0", Rebounds: "0", Assists: "0", Steals: "0", Fouls: "0", Points: "0", Minutes: "2", Blocks: "0", Turnovers: "0", PlusMinus: "+4", Dnp: false, EnteredGame: true},
	{Id: 2991350, FirstName: "Alex", LastName: "Caruso", PositionAbbrev: "SG", Jersey: "4", Active: "false", IsStarter: "false", Fg: "2/2", Ft: "0/0", Threept: "2/2", Rebounds: "1", Assists: "2", Steals: "2", Fouls: "2", Points: "6", Minutes: "13", Blocks: "0", Turnovers: "2", PlusMinus: "+10", Dnp: false, EnteredGame: true},
	{Id: 2581190, FirstName: "Josh", LastName: "Richardson", PositionAbbrev: "SG", Jersey: "0", Active: "false", IsStarter: "true", Fg: "6/12", Ft: "3/3", Threept: "2/5", Rebounds: "2", Assists: "2", Steals: "2", Fouls: "4", Points: "17", Minutes: "27", Blocks: "0", Turnovers: "0", PlusMinus: "-10", Dnp: false, EnteredGame: true},
	{Id: 3945274, FirstName: "Luka", LastName: "Doncic", PositionAbbrev: "PG", Jersey: "77", Active: "false", IsStarter: "true", Fg: "9/19", Ft: "7/8", Threept: "2/5", Rebounds: "4", Assists: "7", Steals: "0", Fouls: "0", Points: "27", Minutes: "34", Blocks: "1", Turnovers: "3", PlusMinus: "-14", Dnp: false, EnteredGame: true},
	{Id: 2528210, FirstName: "Tim", LastName: "Hardaway Jr.", PositionAbbrev: "SG", Jersey: "11", Active: "false", IsStarter: "true", Fg: "4/12", Ft: "1/2", Threept: "1/5", Rebounds: "4", Assists: "2", Steals: "0", Fouls: "1", Points: "10", Minutes: "28", Blocks: "0", Turnovers: "0", PlusMinus: "-9", Dnp: false, EnteredGame: true},
	{Id: 2578185, FirstName: "Dorian", LastName: "Finney-Smith", PositionAbbrev: "SF", Jersey: "10", Active: "false", IsStarter: "true", Fg: "4/9", Ft: "0/0", Threept: "2/6", Rebounds: "2", Assists: "0", Steals: "3", Fouls: "4", Points: "10", Minutes: "29", Blocks: "0", Turnovers: "1", PlusMinus: "-19", Dnp: false, EnteredGame: true},
	{Id: 2531367, FirstName: "Dwight", LastName: "Powell", PositionAbbrev: "C", Jersey: "7", Active: "false", IsStarter: "true", Fg: "4/4", Ft: "2/4", Threept: "1/1", Rebounds: "3", Assists: "1", Steals: "3", Fouls: "2", Points: "11", Minutes: "24", Blocks: "0", Turnovers: "1", PlusMinus: "-11", Dnp: false, EnteredGame: true},
	{Id: 3999, FirstName: "James", LastName: "Johnson", PositionAbbrev: "PF", Jersey: "16", Active: "false", IsStarter: "false", Fg: "1/1", Ft: "2/3", Threept: "0/0", Rebounds: "1", Assists: "4", Steals: "1", Fouls: "0", Points: "4", Minutes: "12", Blocks: "0", Turnovers: "0", PlusMinus: "-7", Dnp: false, EnteredGame: true},
	{Id: 2579260, FirstName: "Trey", LastName: "Burke", PositionAbbrev: "PG", Jersey: "3", Active: "false", IsStarter: "false", Fg: "5/10", Ft: "3/4", Threept: "4/7", Rebounds: "1", Assists: "2", Steals: "0", Fouls: "5", Points: "17", Minutes: "24", Blocks: "1", Turnovers: "0", PlusMinus: "-13", Dnp: false, EnteredGame: true},
	{Id: 2960236, FirstName: "Maxi", LastName: "Kleber", PositionAbbrev: "PF", Jersey: "42", Active: "false", IsStarter: "false", Fg: "1/2", Ft: "2/2", Threept: "1/2", Rebounds: "5", Assists: "0", Steals: "2", Fouls: "0", Points: "5", Minutes: "20", Blocks: "0", Turnovers: "0", PlusMinus: "-8", Dnp: false, EnteredGame: true},
	{Id: 3074797, FirstName: "Wes", LastName: "Iwundu", PositionAbbrev: "SF", Jersey: "25", Active: "true", IsStarter: "false", Fg: "1/1", Ft: "0/0", Threept: "0/0", Rebounds: "0", Assists: "0", Steals: "0", Fouls: "0", Points: "2", Minutes: "2", Blocks: "0", Turnovers: "0", PlusMinus: "-4", Dnp: false, EnteredGame: true},
	{Id: 3934672, FirstName: "Jalen", LastName: "Brunson", PositionAbbrev: "PG", Jersey: "13", Active: "true", IsStarter: "false", Fg: "2/5", Ft: "0/0", Threept: "0/1", Rebounds: "2", Assists: "1", Steals: "0", Fouls: "2", Points: "4", Minutes: "14", Blocks: "0", Turnovers: "3", PlusMinus: "-9", Dnp: false, EnteredGame: true},
	{Id: 4431747, FirstName: "Tyrell", LastName: "Terry", PositionAbbrev: "PG", Jersey: "1", Active: "true", IsStarter: "false", Fg: "1/1", Ft: "0/0", Threept: "0/0", Rebounds: "0", Assists: "1", Steals: "0", Fouls: "0", Points: "2", Minutes: "2", Blocks: "0", Turnovers: "0", PlusMinus: "+1", Dnp: false, EnteredGame: true},
	{Id: 4432811, FirstName: "Josh", LastName: "Green", PositionAbbrev: "SG", Jersey: "8", Active: "true", IsStarter: "false", Fg: "0/0", Ft: "0/0", Threept: "0/0", Rebounds: "0", Assists: "0", Steals: "0", Fouls: "0", Points: "0", Minutes: "2", Blocks: "0", Turnovers: "0", PlusMinus: "-4", Dnp: false, EnteredGame: true},
	{Id: 4376, FirstName: "Boban", LastName: "Marjanovic", PositionAbbrev: "C", Jersey: "51", Active: "false", IsStarter: "false", Fg: "3/4", Ft: "0/0", Threept: "0/0", Rebounds: "1", Assists: "0", Steals: "0", Fouls: "1", Points: "6", Minutes: "7", Blocks: "0", Turnovers: "0", PlusMinus: "-2", Dnp: false, EnteredGame: true},
	{Id: 2991282, FirstName: "Willie", LastName: "Cauley-Stein", PositionAbbrev: "C", Jersey: "33", Active: "true", IsStarter: "false", Fg: "0/3", Ft: "0/0", Threept: "0/0", Rebounds: "2", Assists: "0", Steals: "2", Fouls: "3", Points: "0", Minutes: "14", Blocks: "1", Turnovers: "1", PlusMinus: "-6", Dnp: false, EnteredGame: true},
}
