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

func TestScoreboardClient_GameIDs(t *testing.T) {
	// Given
	tdPath := filepath.Join("testdata", "scoreboard-20201225.json")
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
	c := espn.NewScoreboardClient(ts.Client(), url)
	ids, err := c.GameIDs()
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Then
	expected := []int{401266794, 401266797, 401266798, 401266799, 401266795}
	if !reflect.DeepEqual(ids, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, ids)
	}
}
