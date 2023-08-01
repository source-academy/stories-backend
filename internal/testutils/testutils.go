package testutils

import (
	"fmt"
	"sync"
	"testing"

	"github.com/TwiN/go-color"
	"github.com/source-academy/stories-backend/internal/config"
	"gorm.io/gorm"
)

var (
	testConfig *config.Config
	once       sync.Once
)

func GetTestConf(testEnvPath string) config.Config {
	yellowChevron := color.With(color.Yellow, "❯")
	fmt.Println(yellowChevron, "Getting test conf")
	once.Do(func() { // <-- atomic, does not allow repeating
		fmt.Println(yellowChevron, yellowChevron, yellowChevron, "Creating test conf")
		testConfig, _ = config.LoadFromEnvironment(testEnvPath)
	})
	return *testConfig
}

func SetupDBConnection(t *testing.T, dbConfig *config.DatabaseConfig, migrationPath string) (*gorm.DB, func(*testing.T)) {
	Drop(dbConfig)

	// Create test DB
	err := Create(dbConfig)
	if err != nil {
		t.Error(err)
	}

	// Connect to DB
	db, err := connect(dbConfig)
	if err != nil {
		t.Error(err)
	}

	err = migrateDB(db, migrationPath)
	if err != nil {
		t.Error(err)
	}

	return db, func(t *testing.T) {
		close(db)

		// Drop test DB
		Drop(dbConfig)
	}
}
