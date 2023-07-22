package apierrors

import (
	"net/http"
)

type ClientUnauthorizedError struct {
	Message string
}

func (e ClientUnauthorizedError) Error() string {
	return e.Message
}

func (e ClientUnauthorizedError) HTTPStatusCode() int {
	return http.StatusUnauthorized
}
