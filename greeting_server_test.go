package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETGreeting(t *testing.T) {
	server := NewGreetingServer()

	t.Run("returns a greeting", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/greeting/Chauncy", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Hello, Chauncy"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
