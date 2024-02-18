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

func (u *User) create(tx *gorm.DB) *gorm.DB {
	// TODO: If user already exists, but is soft-deleted, undelete the user
	return tx.Create(u)
}

func CreateUser(db *gorm.DB, user *User) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		return user.create(tx).Error
	})
	if err != nil {
		return database.HandleDBError(err, "user")
	}
	return nil
}

func CreateUsers(db *gorm.DB, users *[]*User) (int64, error) {
	// TODO: Use users.create() instead
	// Blocked by `RowsAffected` not being accessible.
	tx := db.Create(users)
	rowCount := tx.RowsAffected
	if err := tx.Error; err != nil {
		return rowCount, database.HandleDBError(err, "user")
	}
	return rowCount, nil
}

func (u *User) delete(tx *gorm.DB, userID uint) *gorm.DB {
	return tx.
		Model(u).
		Where("id = ?", userID).
		First(u). // store the value to be returned
		Delete(u)
}

func DeleteUser(db *gorm.DB, userID int) (User, error) {
	var user User
	err := db.Transaction(func(tx *gorm.DB) error {
		return user.delete(tx, uint(userID)).Error
	})
	if err != nil {
		return user, database.HandleDBError(err, "user")
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, userID int, user *User) (User, error) {
	err := db.
		Model(&user).
		Where("id = ?", userID).
		First(&user).
		Updates(&user).
		Error
	if err != nil {
		return *user, database.HandleDBError(err, "user")
	}
	return *user, nil
}
