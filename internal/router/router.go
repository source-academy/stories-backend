package router

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/utils/constants"
)

func Setup(config *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Handle CORS
	options := cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}
	if config.Environment == constants.ENV_DEVELOPMENT {
		options.AllowedOrigins = []string{"https://*", "http://*"}
	}
	r.Use(cors.Handler(options))

	// Define routes
	r.Get("/", controller.HandleHealthCheck)

	r.Get("/stories", controller.GetStories)
	r.Post("/stories", controller.CreateStory)

	return r
}
