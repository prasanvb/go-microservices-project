package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "3600"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting the broker service on port %s\n", webPort)

	// define http server
	// Implementation allows Multiplexing Request Handlers
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("broker service server error", err)
	}
}
