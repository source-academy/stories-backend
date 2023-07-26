package testutils

import (
	"testing"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
)

func SetupDBConnection(t *testing.T, dbConfig *config.DatabaseConfig) (*gorm.DB, func(*testing.T)) {
	err := database.Drop(dbConfig)
	if err != nil {
		t.Error(err)
	}

	// Create test DB
	err = database.Create(dbConfig)
	if err != nil {
		t.Error(err)
	}

	// Connect to DB
	db, err := database.Connect(dbConfig)
	if err != nil {
		t.Error(err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		t.Error(err)
	}

	return db, func(t *testing.T) {
		database.Close(db)

		// Drop test DB
		database.Drop(dbConfig)
	}
}
