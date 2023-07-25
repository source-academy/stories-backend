package apierrors

import (
	"net/http"
)

type ClientConflictError struct {
	Message string
}

func (e ClientConflictError) Error() string {
	return e.Message
}

func (e ClientConflictError) HTTPStatusCode() int {
	return http.StatusConflict
}
