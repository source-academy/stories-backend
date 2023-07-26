package database

import (
	"errors"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/config"
	"gorm.io/gorm"
)

func connectAnonDB(conf config.DatabaseConfig) (*gorm.DB, error) {
	conf.DatabaseName = ""
	return Connect(&conf)
}

func Create(conf *config.DatabaseConfig) error {
	if conf.DatabaseName == "" {
		return errors.New("Failed to create database: no database name provided.")
	}

	db, err := connectAnonDB(*conf)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	defer Close(db)

	// check if db exists
	logrus.Infof("Checking if database %s exists.", conf.DatabaseName)
	result := db.Raw("SELECT * FROM pg_database WHERE datname = ?", conf.DatabaseName)
	if result.Error != nil {
		return result.Error
	}

	// if not exists create it
	rec := make(map[string]interface{})
	if result.Find(rec); len(rec) == 0 {
		logrus.Infof("Database %s does not exist. Creating...", conf.DatabaseName)

		create_command := fmt.Sprintf("CREATE DATABASE %s", conf.DatabaseName)
		result := db.Exec(create_command)

		if result.Error != nil {
			return result.Error
		}
		logrus.Infof("Database %s created.", conf.DatabaseName)
		return nil
	}

	logrus.Infof("Database %s exists.", conf.DatabaseName)

	return nil
}

func Drop(conf *config.DatabaseConfig) error {
	if conf.DatabaseName == "" {
		return errors.New("Failed to create database: no database name provided.")
	}

	db, err := connectAnonDB(*conf)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	defer Close(db)

	drop_command := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", conf.DatabaseName)
	result := db.Exec(drop_command)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MigrateDB(d *gorm.DB) error {
	db, err := d.DB()
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	migrations := (migrate.FileMigrationSource{
		Dir: "../migrations",
	})

	_, err = migrate.ExecMax(db, "postgres", migrations, migrate.Up, 0)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
