package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	UserID  int             `gorm:"primaryKey"`
	GroupID int             `gorm:"primaryKey"`
	Role    groupenums.Role // non null
}

func GetUserGroupByID(db *gorm.DB, userId int, groupID int) (UserGroup, error) {
	var userGroup UserGroup
	err := db.First(&userGroup, userId, groupID).Error
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
