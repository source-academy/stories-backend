package model

import (
	"time"

	"github.com/google/uuid"
)

type UserDB struct {
	UserID         uuid.UUID
	GithubUsername string
	GithubID       int
	CreatedAt      time.Time
	DeletedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	UserID         uuid.UUID `json:"userId"`
	GithubUsername string    `json:"githubUsername"`
	GithubID       int       `json:"githubId"`
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

func GetUserByID(userID uuid.UUID) *User {
	var user User
	DB.First(&user, "user_id = ?", userID)
	return &user
}

func CreateUser(user User) {
	newUser := User{
		UserID:         uuid.New(),
		GithubUsername: user.GithubUsername,
		GithubID:       user.GithubID,
	}
	DB.Create(&newUser)
}
