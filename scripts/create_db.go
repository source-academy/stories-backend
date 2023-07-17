package main

import (
	"errors"
	"fmt"

	"github.com/source-academy/stories-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDBServer(conf config.DatabaseConfig) (*gorm.DB, error) {
	conf.DatabaseName = ""
	dsn := conf.ToDataSourceName()
	return connect(dsn)
}

func ConnectToDB(conf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := conf.ToDataSourceName()
	return connect(dsn)
}

func connect(dsn string) (*gorm.DB, error) {
	driver := postgres.Open(dsn)

	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbName, err := getConnectedDBName(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(blueNet, "Connected to database", dbName+".")

	return db, nil
}

func Close(d *gorm.DB) {
	db, err := d.DB()
	if err != nil {
		panic(err)
	}

	dbName, err := getConnectedDBName(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(blueNet, "Closing connection with database", dbName+".")

	if err := db.Close(); err != nil {
		panic(err)
	}
}

func Create(db *gorm.DB, dbconf *config.DatabaseConfig) error {
	if dbconf.DatabaseName == "" {
		return errors.New("Failed to create database: no database name provided.")
	}

	// check if db exists
	fmt.Println(yellowChevron, "Checking if database", dbconf.DatabaseName, "exists.")
	result := db.Raw("SELECT * FROM pg_database WHERE datname = ?", dbconf.DatabaseName)
	if result.Error != nil {
		return result.Error
	}

	// if not exists create it
	rec := make(map[string]interface{})
	if result.Find(rec); len(rec) == 0 {
		fmt.Println(yellowChevron, "Database", dbconf.DatabaseName, "does not exist. Creating...")

		create_command := fmt.Sprintf("CREATE DATABASE %s", dbconf.DatabaseName)
		result := db.Exec(create_command)

		if result.Error != nil {
			return result.Error
		}
	}

	fmt.Println(yellowChevron, "Database", dbconf.DatabaseName, "exists.")

	return nil
}

func Drop(dbserver *gorm.DB, dbconf *config.DatabaseConfig) error {
	drop_command := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbconf.DatabaseName)
	result := dbserver.Exec(drop_command)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getConnectedDBName(db *gorm.DB) (string, error) {
	var dbName string
	result := db.Raw("SELECT current_database();").Scan(&dbName)
	if result.Error != nil {
		return "", result.Error
	}
	return dbName, nil
}