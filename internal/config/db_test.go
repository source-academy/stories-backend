package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Update tests when missing data causes fallback to default values
func TestToDataSourceName(t *testing.T) {
	t.Run("should return correct data source name", func(t *testing.T) {
		dbConfig := DatabaseConfig{
			TimeZone:     "Asia/Singapore",
			Host:         "localhost",
			Port:         5432,
			User:         "golang",
			Password:     "p@ssw0rd",
			DatabaseName: "testing-db",
		}
		dsn := dbConfig.ToDataSourceName()
		assert.Equal(t, fmt.Sprintf(
			"TimeZone=%s host=%s port=%d user=%s password=%s dbname=%s ",
			dbConfig.TimeZone,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DatabaseName,
		), dsn)
	})
	t.Run("should not show missing values", func(t *testing.T) {
		t.Run("should not show time zone when missing", func(t *testing.T) {
			dbConfig := DatabaseConfig{
				Host:         "localhost",
				Port:         5432,
				User:         "golang",
				Password:     "p@ssw0rd",
				DatabaseName: "testing-db",
			}
			dsn := dbConfig.ToDataSourceName()
			assert.Equal(t, fmt.Sprintf(
				"host=%s port=%d user=%s password=%s dbname=%s ",
				dbConfig.Host,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.Password,
				dbConfig.DatabaseName,
			), dsn)
		})
		t.Run("should not show host when missing", func(t *testing.T) {
			dbConfig := DatabaseConfig{
				TimeZone:     "Asia/Singapore",
				Port:         5432,
				User:         "golang",
				Password:     "p@ssw0rd",
				DatabaseName: "testing-db",
			}
			dsn := dbConfig.ToDataSourceName()
			assert.Equal(t, fmt.Sprintf(
				"TimeZone=%s port=%d user=%s password=%s dbname=%s ",
				dbConfig.TimeZone,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.Password,
				dbConfig.DatabaseName,
			), dsn)
		})
		t.Run("should not show port when missing", func(t *testing.T) {
			dbConfig := DatabaseConfig{
				TimeZone:     "Asia/Singapore",
				Host:         "localhost",
				User:         "golang",
				Password:     "p@ssw0rd",
				DatabaseName: "testing-db",
			}
			dsn := dbConfig.ToDataSourceName()
			assert.Equal(t, fmt.Sprintf(
				"TimeZone=%s host=%s user=%s password=%s dbname=%s ",
				dbConfig.TimeZone,
				dbConfig.Host,
				dbConfig.User,
				dbConfig.Password,
				dbConfig.DatabaseName,
			), dsn)
		})
		t.Run("should not show user when missing", func(t *testing.T) {
			dbConfig := DatabaseConfig{
				TimeZone:     "Asia/Singapore",
				Host:         "localhost",
				Port:         5432,
				Password:     "p@ssw0rd",
				DatabaseName: "testing-db",
			}
			dsn := dbConfig.ToDataSourceName()
			assert.Equal(t, fmt.Sprintf(
				"TimeZone=%s host=%s port=%d password=%s dbname=%s ",
				dbConfig.TimeZone,
				dbConfig.Host,
				dbConfig.Port,
				dbConfig.Password,
				dbConfig.DatabaseName,
			), dsn)
		})
		t.Run("should not show password when missing", func(t *testing.T) {
			dbConfig := DatabaseConfig{
				TimeZone:     "Asia/Singapore",
				Host:         "localhost",
				Port:         5432,
				User:         "golang",
				DatabaseName: "testing-db",
			}
			dsn := dbConfig.ToDataSourceName()
			assert.Equal(t, fmt.Sprintf(
				"TimeZone=%s host=%s port=%d user=%s dbname=%s ",
				dbConfig.TimeZone,
				dbConfig.Host,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.DatabaseName,
			), dsn)
		})
		t.Run("should not show database name when missing", func(t *testing.T) {
			dbConfig := DatabaseConfig{
				TimeZone: "Asia/Singapore",
				Host:     "localhost",
				Port:     5432,
				User:     "golang",
				Password: "p@ssw0rd",
			}
			dsn := dbConfig.ToDataSourceName()
			assert.Equal(t, fmt.Sprintf(
				"TimeZone=%s host=%s port=%d user=%s password=%s ",
				dbConfig.TimeZone,
				dbConfig.Host,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.Password,
			), dsn)
		})
	})
}
