package database

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// We define our own type to avoid collisions
type contextKey string

const (
	dbKey contextKey = "database_context"
)

func MakeMiddlewareFrom(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a new session from the database for each request
			ctx := context.WithValue(r.Context(), dbKey, db.Session(&gorm.Session{}))
			// Inject the new context into the request
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetDBFrom(r *http.Request) (*gorm.DB, error) {
	db, ok := r.Context().Value(dbKey).(*gorm.DB)
	if !ok {
		return nil, errors.New("Could not get database from request context")
	}
	return db, nil
}
