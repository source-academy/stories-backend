package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/internal/router"
	envutils "github.com/source-academy/stories-backend/internal/utils/env"
)

func main() {
	// Load configuration
	// FIXME: Remove hardcoding when further configuration is added
	conf, err := config.LoadFromEnvironment(".env")
	if err != nil {
		logrus.Errorln(err)
	}

	// Set log level
	if conf.Environment == envutils.ENV_DEVELOPMENT {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Connect to the database
	db, err := database.Connect(conf.Database)
	if err != nil {
		logrus.Errorln(err)
	}
	defer database.Close(db)

	var injectMiddlewares []func(http.Handler) http.Handler

	// Inject DB session into request context
	injectMiddlewares = append(injectMiddlewares, database.MakeMiddlewareFrom(db))

	// Initialze Sentry configuration
	if conf.Environment == envutils.ENV_PRODUCTION {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: conf.SentryDSN,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
		})
		if err != nil {
			logrus.Errorln("sentry.Init:", err)

		}
		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)

		// Setup Sentry middleware. Adapted from:
		// https://gist.github.com/rhcarvalho/66130d1252d4a7b1fbaeacfe3687eaf3
		sentryMiddleware := sentryhttp.New(sentryhttp.Options{
			Repanic: true,
		})
		// Important: Chi has a middleware stack and thus it is important to put the
		// Sentry handler on the appropriate place. If using middleware.Recoverer,
		// the Sentry middleware must come afterwards (and configure it with
		// Repanic: true).
		injectMiddlewares = append(injectMiddlewares, sentryMiddleware.Handle)
	}

	// Setup router
	r := router.Setup(conf, injectMiddlewares)

	// Start server
	logrus.Infof("Starting server on %s port %d", conf.Host, conf.Port)
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		logrus.Errorln(err)
	}
}
