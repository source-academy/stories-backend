package router

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/config"
	envutils "github.com/source-academy/stories-backend/internal/utils/env"
)

func Setup(config *config.Config, injectMiddleWares []func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	for _, injectMiddleware := range injectMiddleWares {
		r.Use(injectMiddleware)
	}
	// Handle CORS
	options := cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}
	if config.Environment == envutils.ENV_DEVELOPMENT {
		options.AllowedOrigins = []string{"https://*", "http://*"}
	}
	r.Use(cors.Handler(options))

	// Define routes
	r.Get("/", controller.HandleHealthCheck)

	r.Route("/stories", func(r chi.Router) {
		r.Get("/", controller.GetStories)
		r.Get("/{storyID}", controller.GetStory)
		r.Post("/", controller.CreateStory)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", controller.GetUsers)
		r.Get("/{userID}", controller.GetUser)
		r.Post("/", controller.CreateUser)
	})

	return r
}
