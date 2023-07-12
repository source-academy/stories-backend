package model

import "time"

type User struct {
	UserID         int       `json:"userId"`
	GithubUsername string    `json:"githubUsername"`
	GithubID       int       `json:"githubId"`
	CreatedAt      time.Time `json:"createdAt"`
	DeletedAt      time.Time `json:"deletedAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
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
