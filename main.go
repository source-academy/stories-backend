package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/config"
	"github.com/source-academy/stories-backend/utils/constants"
)

const (
	WELCOME_MESSAGE = "Hello from Source Academy Stories!"
)

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: Abstract router setup logic
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Handle CORS
	options := cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}
	if conf.Environment == constants.ENV_DEVELOPMENT {
		options.AllowedOrigins = []string{"https://*", "http://*"}
	}
	r.Use(cors.Handler(options))

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, WELCOME_MESSAGE)
	})

	// Start server
	log.Println("Starting server on port 8080")
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
