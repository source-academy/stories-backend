package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/internal/router"
	"github.com/source-academy/stories-backend/internal/utils/constants"
)

const (
	STARTUP_MESSAGE = "Starting server..."
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
		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
		// Notify that server is starting
		sentry.CaptureMessage(STARTUP_MESSAGE)
	}

	// Setup router
	r := router.Setup(conf)

	// Start server
	log.Printf("Starting server on %s port %d", conf.Host, conf.Port)
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
