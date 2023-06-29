package database

import (
	"testing"

	"github.com/source-academy/stories-backend/config"
	"github.com/source-academy/stories-backend/utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	conf := &config.DatabaseConfig{
		TimeZone:     constants.DB_DEFAULT_TIMEZONE,
		Host:         "localhost",
		Port:         constants.DB_DEFAULT_PORT,
		User:         "postgres",
		DatabaseName: constants.DB_DEFAULT_NAME,
	}

	t.Run("should connect to database", func(t *testing.T) {
		db, err := Connect(conf)
		defer Close(db)

		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
	t.Run("should return correct database name", func(t *testing.T) {
		db, _ := Connect(conf)
		defer Close(db)

		// Get currently connected database name
		var dbName string
		db.Raw("SELECT current_database()").Scan(&dbName)
		assert.Equal(t, constants.DB_DEFAULT_NAME, dbName)
	})
}
