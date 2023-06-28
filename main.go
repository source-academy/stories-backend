package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const (
	WELCOME_MESSAGE = "Hello from Source Academy Stories!"
)

func main() {
	// TODO: Abstract router setup logic
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Handle CORS
	options := cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}
	if os.Getenv("GO_ENV") == "development" {
		options.AllowedOrigins = []string{"https://*", "http://*"}
	}
	r.Use(cors.Handler(options))

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, WELCOME_MESSAGE)
	})

	// Start server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
