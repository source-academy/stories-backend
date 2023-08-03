package auth

import (
	"context"
	"errors"
	"net/http"
)

type contextKey string

const (
	userContextKey contextKey = "auth_context"
)

// Injects the session into the request context
func injectUserIDToContext(r *http.Request, userID int) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, userID)
	return r.WithContext(ctx)
}

func GetUserIDFrom(r *http.Request) (*int, error) {
	userID, ok := r.Context().Value(userContextKey).(int)
	if !ok {
		return nil, errors.New("Could not get user ID from request context")
	}
	return &userID, nil
}
