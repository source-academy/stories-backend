package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/internal/logger"
	"github.com/source-academy/stories-backend/internal/router"
)

func main() {
	// Setup logger
	l := logger.Setup(os.Stdout)

	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		l.Error.Fatalln(err)
	}

	// Connect to the database
	db, err := database.Connect(conf.Database)
	if err != nil {
		l.Error.Fatalln(err)
	}
	defer database.Close(db)

	// Setup router
	r := router.Setup(conf)

	// Start server
	l.Info.Printf("Starting server on %s port %d", conf.Host, conf.Port)
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		l.Error.Fatalln(err)
	}
}
