package database

import (
	"github.com/source-academy/stories-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(conf *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := conf.ToDataSourceName()
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
	return db.Close()
}
