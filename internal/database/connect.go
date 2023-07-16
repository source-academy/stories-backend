package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
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
