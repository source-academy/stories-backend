package auth

import (
	"context"
	"errors"
	"net/http"
)

type contextKey string

const (
	authKey contextKey = "auth_context"
)

type Session struct {
	UserID int
}

// Injects the session into the request context
func injectSession(r *http.Request, session Session) *http.Request {
	ctx := context.WithValue(r.Context(), authKey, session)
	return r.WithContext(ctx)
}

func GetSessionFrom(r *http.Request) (*Session, error) {
	session, ok := r.Context().Value(authKey).(*Session)
	if !ok {
		return nil, errors.New("Could not get session from request context")
	}
	return session, nil
}
