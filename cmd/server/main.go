package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erichaase/fantasy/internal/gameline"
)

func main() {
	lh := flag.Bool("localhost", false, "Configure HTTP server to listen to localhost for local development")
	flag.Parse()

	addr := ""
	if *lh {
		addr = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ap := fmt.Sprintf("%s:%s", addr, port)
	log.Printf("Listening on 'http://%s'", ap)

	http.HandleFunc("/lines", gameline.LinesHandler)
	log.Fatal(http.ListenAndServe(ap, nil))
}
