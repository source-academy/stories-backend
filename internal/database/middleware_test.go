package database

import (
	"context"
	"net/http"
	"testing"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestMakeMiddlewareFrom(t *testing.T) {
	t.Run("should return a middleware function", func(t *testing.T) {
		// Set up
		conf, _ := config.LoadFromEnvironment()
		db, _ := Connect(conf.Database)
		defer Close(db)

		middleware := MakeMiddlewareFrom(db)
		assert.NotNil(t, middleware, "Expected middleware to not be nil")
		mockMiddleware := func(next http.Handler) http.Handler { return next }
		assert.IsType(t, mockMiddleware, middleware, "Expected middleware to be a function")
	})

}

func TestGetDBFrom(t *testing.T) {
	t.Run("should return an error when database is not in context", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		_, err := GetDBFrom(req)
		assert.NotNil(t, err, "Expected error to not be nil")
	})

	t.Run("should return a database when database is in context", func(t *testing.T) {
		// Set up
		conf, _ := config.LoadFromEnvironment()
		db, _ := Connect(conf.Database)
		defer Close(db)

		// Create request
		r, _ := http.NewRequest("GET", "/", nil)
		// Inject database into context. Unlike the actual code,
		// We intentionally do not duplicate the db session to
		// allow for comparison later.
		ctx := context.WithValue(r.Context(), dbKey, db)
		// Get database from context
		database, err := GetDBFrom(r.WithContext(ctx))

		assert.Nil(t, err, "Expected error to be nil")
		assert.Equal(t, db, database, "Expected database to be the same as the one in context")
	})
}
