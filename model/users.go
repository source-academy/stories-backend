package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GithubUsername string
	GithubID       int
}

func GetAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserByID(db *gorm.DB, id int) (User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return user, database.HandleDBError(err, "user")
	}
	return user, err
}

func CreateUser(db *gorm.DB, user *User) {
	db.Create(user)
}
