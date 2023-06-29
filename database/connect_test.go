package database

import (
	"testing"

	"github.com/source-academy/stories-backend/config"
	"github.com/source-academy/stories-backend/utils/constants"
	"github.com/stretchr/testify/assert"
)

func ignoreError(err error) {}

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
		defer ignoreError(Close(db))

		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
	t.Run("should return correct database name", func(t *testing.T) {
		db, _ := Connect(conf)
		defer ignoreError(Close(db))

		// Get currently connected database name
		var dbName string
		db.Raw("SELECT current_database()").Scan(&dbName)
		assert.Equal(t, constants.DB_DEFAULT_NAME, dbName)
	})
	// TODO: Populate these with actual tables once schema is finalized
	// t.Run("should show a correct list of tables", func(t *testing.T) {
	// 	db, _ := Connect(conf)
	// 	defer ignoreError(Close(db))

	// 	// Get list of tables in the database
	// 	var tables []string
	// 	err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error
	// 	assert.Nil(t, err)

	// 	assert.Equal(t, []string{}, tables)
	// })
}
