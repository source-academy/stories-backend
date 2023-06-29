package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/config"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/utils/constants"
)

const (
	STARTUP_MESSAGE = "Starting server..."
	WELCOME_MESSAGE = "Hello from Source Academy Stories!"
)

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialze Sentry configuration
	// TODO: Migrate logic to routing middleware
	//       or internal logger package.
	if conf.Environment == constants.ENV_PRODUCTION {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: conf.SentryDSN,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
		})
		if err != nil {
			log.Fatalln("sentry.Init:", err)

		}
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	// Notify that server is starting
	sentry.CaptureMessage(STARTUP_MESSAGE)

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
