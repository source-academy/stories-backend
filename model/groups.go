package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name string
}

func GetGroupByID(db *gorm.DB, id int) (Group, error) {
	var group Group
	err := db.First(&group, id).Error
	if err != nil {
		return group, database.HandleDBError(err, "group")
	}
	return group, nil
}

func CreateGroup(db *gorm.DB, group *Group) error {
	err := db.Create(group).Error
	if err != nil {
		return database.HandleDBError(err, "group")
	}
	return nil
}
