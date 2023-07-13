package model

import (
	"github.com/source-academy/stories-backend/view"
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	GithubUsername string
	GithubID       int
}

func MapUserDBToUser(userDB UserDB) view.User {
	user := view.User{
		ID:             userDB.ID,
		GithubUsername: userDB.GithubUsername,
		GithubID:       userDB.GithubID,
	}
	return user
}

func GetAllUsers() []view.User {
	var users []view.User
	DB.Find(&users)
	return users
}

func GetUserByID(userID int) *view.User {
	var user view.User
	DB.First(&user, userID)
	return &user
}

func CreateUser(user view.User) {
	DB.Create(&user)
}
