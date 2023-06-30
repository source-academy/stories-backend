package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/internal/router"
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
