package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func createDB(db *gorm.DB, dbName string) error {
	if dbName == "" {
		return errors.New("Failed to create database: no database name provided.")
	}

	// check if db exists
	fmt.Println(yellowChevron, "Checking if database", dbName, "exists.")
	result := db.Raw("SELECT * FROM pg_database WHERE datname = ?", dbName)
	if result.Error != nil {
		return result.Error
	}

	// if not exists create it
	rec := make(map[string]interface{})
	if result.Find(rec); len(rec) != 0 {
		fmt.Println(yellowChevron, "Database", dbName, "exists.")
		return nil
	}

	fmt.Println(yellowChevron, "Database", dbName, "does not exist. Creating...")
	create_command := fmt.Sprintf("CREATE DATABASE %s", dbName)
	err := db.Exec(create_command).Error
	if err != nil {
		return err
	}

	fmt.Println(greenTick, "Created database:", dbName)
	return nil
}

func dropDB(db *gorm.DB, dbName string) error {
	drop_command := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)
	err := db.Exec(drop_command).Error
	return err
}

func getDBName(db *gorm.DB) (string, error) {
	var dbName string
	err := db.Raw("SELECT current_database();").Scan(&dbName).Error
	return dbName, err
}
