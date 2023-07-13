package model

import (
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	GithubUsername string
	GithubID       int
}

type User struct {
	ID             uint   `json:"userId"`
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}

func MapUserDBToUser(userDB UserDB) User {
	user := User{
		ID:             userDB.ID,
		GithubUsername: userDB.GithubUsername,
		GithubID:       userDB.GithubID,
	}
	return user
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
