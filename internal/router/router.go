package router

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/controller/stories"
	"github.com/source-academy/stories-backend/controller/users"

	// FIXME: Name clash
	usergroupscontroller "github.com/source-academy/stories-backend/controller/usergroups"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/config"
	usergroups "github.com/source-academy/stories-backend/internal/usergroups"
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           60,
	}
	if config.Environment == envutils.ENV_DEVELOPMENT {
		options.AllowedOrigins = []string{"https://*", "http://*"}
	} else if config.AllowedOrigins != nil {
		options.AllowedOrigins = *config.AllowedOrigins
	}
	r.Use(cors.Handler(options))
	r.Use(middleware.NoCache)

	// Define routes
	r.Group(func(r chi.Router) {
		// Public routes
		r.Get("/", controller.HandleHealthCheck)
	})

	r.Group(func(r chi.Router) {
		// Private routes
		r.Use(auth.MakeMiddlewareFrom(config))
		r.Route("/groups/{groupID}", func(r chi.Router) {
			// Group specific routes
			r.Use(usergroups.InjectUserGroupIntoContext)
			r.Route("/stories", func(r chi.Router) {
				r.Get("/", handleAPIError(stories.HandleList))
				r.Get("/draft", handleAPIError(stories.HandleListDraft))
				r.Get("/published", handleAPIError(stories.HandleListPublished))
				r.Get("/{storyID}", handleAPIError(stories.HandleRead))
				r.Put("/{storyID}", handleAPIError(stories.HandleUpdate))
				r.Put("/{storyID}/publish", handleAPIError(stories.HandlePublish))
				r.Delete("/{storyID}", handleAPIError(stories.HandleDelete))
				r.Post("/", handleAPIError(stories.HandleCreate))
			})			

			r.Route("/users", func(r chi.Router) {
				r.Get("/", handleAPIError(users.HandleList))
				r.Get("/{userID}", handleAPIError(users.HandleRead))
				r.Delete("/{userID}", handleAPIError(users.HandleDelete))
				r.Post("/", handleAPIError(users.HandleCreate))
				r.Put("/batch", handleAPIError(usergroupscontroller.HandleBatchCreate))
			})
		})

		r.Get("/user", handleAPIError(users.HandleReadSelf))
	})

	return r
}
