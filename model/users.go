package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName      string // FIXME: Use nullable string
	Username      string
	LoginProvider userenums.LoginProvider
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return users, database.HandleDBError(err, "user")
	}
	return users, nil
}

func GetUserByID(db *gorm.DB, id int) (User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return user, database.HandleDBError(err, "user")
	}
	return user, err
}

func CreateUser(db *gorm.DB, user *User) error {
	// TODO: If user already exists, but is soft-deleted, undelete the user
	err := db.Create(user).Error
	if err != nil {
		return database.HandleDBError(err, "user")
	}
	return nil
}

func CreateUsers(db *gorm.DB, users *[]*User) (int64, error) {
	tx := db.Create(users)
	rowCount := tx.RowsAffected
	if err := tx.Error; err != nil {
		return rowCount, database.HandleDBError(err, "user")
	}
	return rowCount, nil
}

func DeleteUser(db *gorm.DB, userID int) (User, error) {
	var user User
	err := db.
		Model(&user).
		Where("id = ?", userID).
		First(&user). // store the value to be returned
		Delete(&user).
		Error
	if err != nil {
		return user, database.HandleDBError(err, "user")
	}
	return user, nil
}
