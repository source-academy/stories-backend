package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	User  User            `gorm:"foreignKey:UserID;references:ID"`
	Group Group           `gorm:"foreignKey:GroupID;references:ID"`
	Role  groupenums.Role // non null
}

func GetUserGroupByID(db *gorm.DB, userID int, groupID int) (UserGroup, error) {
	var userGroup UserGroup

	err := db.Model(&userGroup).
		Preload("users").
		Preload("groups").
		Where("user.id = ?", userID).
		Where("group.id = ?", groupID).
		First(&userGroup).Error
	if err != nil {
		return userGroup, database.HandleDBError(err, "group")
	}
	return userGroup, nil
}

func CreateUserGroup(db *gorm.DB, userGroup *UserGroup) error {
	err := db.Create(userGroup).Error
	if err != nil {
		return database.HandleDBError(err, "group")
	}
	return nil
}
