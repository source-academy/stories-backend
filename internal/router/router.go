package router

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/controller/stories"
	"github.com/source-academy/stories-backend/controller/users"
	"github.com/source-academy/stories-backend/internal/auth"
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

	r.Group(func(r chi.Router) {
		r.Use(auth.MakeMiddlewareFrom(config))
		r.Route("/stories", func(r chi.Router) {
			r.Get("/", handleAPIError(stories.HandleList))
			r.Get("/{storyID}", handleAPIError(stories.HandleRead))
			r.Put("/{storyID}", handleAPIError(stories.HandleUpdate))
			r.Post("/", handleAPIError(stories.HandleCreate))
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", handleAPIError(users.HandleList))
			r.Get("/{userID}", handleAPIError(users.HandleRead))
			r.Post("/", handleAPIError(users.HandleCreate))
			r.Post("/batch", handleAPIError(users.HandleBatchCreate))
		})
	})

	return r
}
