package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GithubUsername string
	GithubID       int
}

func GetAllUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func GetUserByID(userID int) *User {
	var user User
	DB.First(&user, userID)
	return &user
}

func CreateUser(user User) {
	DB.Create(&user)
}
