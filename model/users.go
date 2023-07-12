package model

import "time"

type UserDB struct {
	UserID         int
	GithubUsername string
	GithubID       int
	CreatedAt      time.Time
	DeletedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	UserID         int    `json:"userId"`
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}

func MapUserDBToUser(userDB UserDB) User {
	user := User{
		UserID:         userDB.UserID,
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
