package model

import "time"

type User struct {
	UserID         int       `json:"user_id"`
	GithubUsername string    `json:"github_username"`
	GithubID       int       `json:"github_id"`
	CreatedAt      time.Time `json:"created_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func GetAllUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func GetUserByID(userID int) *User {
	var user User
	if DB.First(&user, userID).Error != nil {
		return nil
	}
	return &user
}

func CreateUser(user User) {
	DB.Create(&user)
}
