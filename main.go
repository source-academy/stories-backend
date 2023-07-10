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
	"github.com/source-academy/stories-backend/internal/utils/constants"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		logrus.Errorln(err)
	}

	// Connect to the database
	// db, err := database.Connect(conf.Database)
	DB, err = database.Connect(conf.Database)
	controller.DB = DB
	model.DB = DB
	if err != nil {
		logrus.Errorln(err)
	}
	defer database.Close(DB)

	var injectMiddlewares []func(http.Handler) http.Handler
	// Initialze Sentry configuration
	if conf.Environment == constants.ENV_PRODUCTION {
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
