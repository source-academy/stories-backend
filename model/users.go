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

func GetAllUsers() []userviews.View {
	var users []userviews.View
	DB.Find(&users)
	return users
}

func GetUserByID(userID int) *userviews.View {
	var user userviews.View
	DB.First(&user, userID)
	return &user
}

func CreateUser(user userviews.View) {
	DB.Create(&user)
}
