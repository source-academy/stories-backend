package model

import (
	userviews "github.com/source-academy/stories-backend/view/users"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GithubUsername string
	GithubID       int
}

func MapUserDBToUser(userDB User) userviews.View {
	user := userviews.View{
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
