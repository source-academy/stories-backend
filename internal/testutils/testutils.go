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
	test_conf *config.Config
	once      sync.Once
)

func GetTestConf(test_env_path string) config.Config {
	yellowChevron := color.With(color.Yellow, "❯")
	fmt.Println(yellowChevron, "Getting test conf")
	once.Do(func() { // <-- atomic, does not allow repeating
		fmt.Println(yellowChevron, yellowChevron, yellowChevron, "Creating test conf")
		test_conf, _ = config.LoadFromEnvironment(test_env_path)
	})
	// if test_conf == nil {
	// 	fmt.Println(yellowChevron, yellowChevron, "Creating test conf")
	// 	test_conf, _ = config.LoadFromEnvironment(test_env_path)
	// }
	return *test_conf
}

func SetupDBConnection(t *testing.T, dbConfig *config.DatabaseConfig, migration_path string) (*gorm.DB, func(*testing.T)) {
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

	err = migrateDB(db, migration_path)
	if err != nil {
		t.Error(err)
	}

	return db, func(t *testing.T) {
		close(db)

		// Drop test DB
		Drop(dbConfig)
	}
}
