package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/erichaase/fantasy/internal/gameline"
)

func main() {
	lh := flag.Bool("localhost", false, "Configure HTTP server to listen to localhost for local development")
	port := flag.Int("port", 80, "The port that the HTTP server listens to for incoming requests")
	flag.Parse()

	addr := ""
	if *lh {
		addr = "localhost"
	}
	ap := fmt.Sprintf("%s:%d", addr, *port)
	log.Printf("Listening on '%s'", ap)

	http.HandleFunc("/lines", handler)
	log.Fatal(http.ListenAndServe(ap, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	q, _ := url.ParseQuery(r.URL.RawQuery)
	d := q.Get("date")

	err := gameline.WriteLinesResponse(w, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
