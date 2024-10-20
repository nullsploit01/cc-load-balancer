package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	serverName := os.Getenv("SERVER_NAME")

	r := NewRouter(serverName)

	fmt.Printf("starting server: %s\n", serverName)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func NewRouter(serverName string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if serverName == "" {
		serverName = "default-server"
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("{\"message\": \"%s says, sup?\"}", serverName)))
	})

	return r
}
