package database

import (
	"testing"

	"github.com/source-academy/stories-backend/internal/testutils"

	"github.com/stretchr/testify/assert"
)

const (
	test_env_path  = "../../.env.testing"
	migration_path = "../../migrations"
)

func ignoreError(func() (err error)) {}

func TestConnect(t *testing.T) {
	conf := testutils.GetTestConf(test_env_path)
	err := testutils.Create(conf.Database)
	if err != nil {
		t.Error(err)
	}
	defer testutils.Drop(conf.Database)

	t.Run("should connect to database", func(t *testing.T) {
		db, err := Connect(conf.Database)
		defer ignoreError(func() error { return Close(db) })

		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
	t.Run("should return correct database name", func(t *testing.T) {
		db, _ := Connect(conf.Database)
		defer ignoreError(func() error { return Close(db) })

		// Get currently connected database name
		var dbName string
		db.Raw("SELECT current_database()").Scan(&dbName)
		assert.Equal(t, conf.Database.DatabaseName, dbName)
	})
	// TODO: Populate these with actual tables once schema is finalized
	// t.Run("should show a correct list of tables", func(t *testing.T) {
	// 	db, _ := Connect(conf.Database)
	//  defer ignoreError(func() error { return Close(db) })

	// 	// Get list of tables in the database
	// 	var tables []string
	// 	err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error
	// 	assert.Nil(t, err)

	// 	assert.Equal(t, []string{}, tables)
	// })
}
