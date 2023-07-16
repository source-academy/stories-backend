package database

import (
	"errors"
	"fmt"

	"github.com/source-academy/stories-backend/internal/config"
	"gorm.io/gorm"
)

func CreateAndConnect(db *gorm.DB, dbconf *config.DatabaseConfig) (*gorm.DB, error) {
	if dbconf.DatabaseName == "" {
		// TODO: refactor to use central contexted error. or make sure this is not empty
		return nil, errors.New("Failed to create database: no database name provided.")
	}

	// check if db exists
	fmt.Println("Checking if database", dbconf.DatabaseName, "exists.")
	query := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", dbconf.DatabaseName)
	result := db.Raw(query)
	if result.Error != nil {
		return nil, result.Error
	}

	// if not exists create it
	rec := make(map[string]interface{})
	if result.Find(rec); len(rec) == 0 {
		create_command := fmt.Sprintf("CREATE DATABASE %s;", dbconf.DatabaseName)
		result := db.Exec(create_command)

		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		fmt.Println("Database", dbconf.DatabaseName, "exists. Connecting...")
	}

	// !!!!! somehow this does not work
	// create the database if not exist
	// create_command := `'CREATE DATABASE "` + conf.DatabaseName + `"'`
	// conditional_command := `SELECT ` + psql_command + ` WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '` + conf.DatabaseName + `')\gexec`
	// fmt.Println(conditional_command)
	// result := db.Raw(conditional_command)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	return Connect(dbconf.ToDataSourceName())
}

func Drop(dbserver *gorm.DB, dbconf *config.DatabaseConfig) error {
	drop_command := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbconf.DatabaseName)
	result := dbserver.Exec(drop_command)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
