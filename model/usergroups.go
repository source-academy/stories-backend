package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserGroup struct {
	gorm.Model
	UserID  int
	User    User
	GroupID int
	Group   Group
	Role    groupenums.Role // non null
}

func GetUserGroupByID(db *gorm.DB, userID int, groupID int) (UserGroup, error) {
	var userGroup UserGroup

	err := db.Model(&userGroup).
		Preload(clause.Associations).
		Where(UserGroup{UserID: userID, GroupID: groupID}).
		First(&userGroup).Error

	if err != nil {
		return userGroup, database.HandleDBError(err, "userGroup")
	}

	return userGroup, nil
}

func CreateUserGroup(db *gorm.DB, userGroup *UserGroup) error {
	err := db.Create(userGroup).Error
	if err != nil {
		return database.HandleDBError(err, "userGroup")
	}
	return nil
}
