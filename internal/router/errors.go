package router

import (
	"errors"
	"net/http"

	apierrors "github.com/source-academy/stories-backend/internal/errors"
)

func handleAPIError(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			// No error, response already written
			return
		}

		var clientError apierrors.ClientError
		if errors.As(err, &clientError) {
			// Client error (status 4xx), write error message and status code
			http.Error(w, clientError.Error(), clientError.HTTPStatusCode())
			return
		}

		// 500 Internal Server Error as a catch-all
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
