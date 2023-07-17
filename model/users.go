package model

import (
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

func GetUserByID(db *gorm.DB, id int) User {
	var user User
	db.First(&user, id)
	return user
}

func CreateUser(db *gorm.DB, user *User) {
	db.Create(user)
}
