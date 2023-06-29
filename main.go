package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/config"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/database"
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

	// Connect to the database
	db, err := database.Connect(conf.Database)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close(db)

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

	r.Get("/stories", controller.GetStories)
	r.Post("/stories", controller.CreateStory)

	// Start server
	log.Printf("Starting server on %s port %d", conf.Host, conf.Port)
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
