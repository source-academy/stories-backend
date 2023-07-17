package main

import (
	"errors"
	"fmt"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDBServer(conf *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := conf.ToEmptyDataSourceName()
	driver := postgres.Open(dsn)

	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Close(d *gorm.DB) error {
	db, err := d.DB()
	if err != nil {
		return err
	}
	var dbName string
	result := d.Raw("SELECT current_database();").Scan(&dbName)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println(yellowChevron, "Closing connection with database", dbName+".")
	return db.Close()
}

func CreateAndConnect(db *gorm.DB, dbconf *config.DatabaseConfig) (*gorm.DB, error) {
	if dbconf.DatabaseName == "" {
		// TODO: refactor to use central contexted error. or make sure this is not empty
		return nil, errors.New("Failed to create database: no database name provided.")
	}

	// check if db exists
	fmt.Println(yellowChevron, "Checking if database", dbconf.DatabaseName, "exists.")
	result := db.Raw("SELECT * FROM pg_database WHERE datname = ?", dbconf.DatabaseName)
	if result.Error != nil {
		return nil, result.Error
	}

	// if not exists create it
	rec := make(map[string]interface{})
	if result.Find(rec); len(rec) == 0 {
		fmt.Println(yellowChevron, "Database", dbconf.DatabaseName, "does not exist. Creating...")

		create_command := fmt.Sprintf("CREATE DATABASE %s", dbconf.DatabaseName)
		result := db.Exec(create_command)

		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		fmt.Println(yellowChevron, "Database", dbconf.DatabaseName, "exists. Connecting...")
	}

	return database.Connect(dbconf)
}

func Drop(dbserver *gorm.DB, dbconf *config.DatabaseConfig) error {
	drop_command := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbconf.DatabaseName)
	result := dbserver.Exec(drop_command)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
