package router

import (
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

		// Error, write error message and status code
		apierrors.ServeHTTP(w, r, err)
	}
}
