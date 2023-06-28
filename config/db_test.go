package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDataSourceName(t *testing.T) {
	t.Run("should return correct data source name", func(t *testing.T) {
		dbConfig := DatabaseConfig{
			TimeZone: "Asia/Singapore",
			Host:     "localhost",
			Port:     5432,
			User:     "golang",
			Password: "p@ssw0rd",
			Database: "testing-db",
		}
		dsn := dbConfig.ToDataSourceName()
		assert.Equal(t, fmt.Sprintf(
			"TimeZone=%s host=%s port=%d user=%s password=%s dbname=%s ",
			dbConfig.TimeZone,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Database,
		), dsn)
	})
}
