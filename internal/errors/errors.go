package apierrors

import (
	"errors"
	"net/http"
)

// ClientError is an interface for errors that should be returned to the client.
// They generally start with a 4xx HTTP status code.
type ClientError interface {
	error
	HTTPStatusCode() int
}

func ServeHTTP(w http.ResponseWriter, r *http.Request, err error) {
	var clientError ClientError
	if errors.As(err, &clientError) {
		// Client error (status 4xx), write error message and status code
		http.Error(w, clientError.Error(), clientError.HTTPStatusCode())
		return
	}

	// 500 Internal Server Error as a catch-all
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
