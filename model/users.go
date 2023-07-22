package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string
	LoginProvider userenums.LoginProvider
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return users, database.HandleDBError(err, "user")
	}
	return users, nil
}

func GetUserByID(db *gorm.DB, id int) (User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return user, database.HandleDBError(err, "user")
	}
	return user, err
}

func CreateUser(db *gorm.DB, user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return database.HandleDBError(err, "user")
	}
	return nil
}
