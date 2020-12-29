package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("/lines", gameline.LinesHandler)
	log.Fatal(http.ListenAndServe(ap, nil))
}
