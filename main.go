package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const port = 5000

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	server := NewGreetingServer()
	log.Infof("Starting greeting-server on port: %d", port)
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
