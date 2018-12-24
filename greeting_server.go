package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const defaultGreeting = "Hello"
const defaultName = "World"

type GreetingServer struct {
	greeting string
	router   *http.ServeMux
}

func NewGreetingServer() *GreetingServer {
	greeting := os.Getenv("GREETING")

	if greeting == "" {
		greeting = defaultGreeting
	}

	gs := &GreetingServer{greeting, http.NewServeMux()}

	gs.router.Handle("/greeting/", http.HandlerFunc(gs.greetingHandler))

	return gs
}

func (gs *GreetingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gs.router.ServeHTTP(w, r)
}

func (gs *GreetingServer) greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/greeting/"):]

	if name == "" {
		name = defaultName
	}

	message := fmt.Sprintf("%s, %s", gs.greeting, name)

	w.WriteHeader(http.StatusOK)

	log.WithFields(log.Fields{
		"status": "200",
		"name":   name,
	}).Info("Dispatched greeting")

	fmt.Fprint(w, message)
}
