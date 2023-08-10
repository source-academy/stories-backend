package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CourseGroup struct {
	gorm.Model
	CourseID uint // assuming that course is unique
	GroupID  uint
	Group    Group
}

func GetGroupByCourseID(db *gorm.DB, courseID int) (Group, error) {
	var courseGroup CourseGroup
	err := db.Preload(clause.Associations).First(&courseGroup, courseID).Error
	if err != nil {
		return courseGroup.Group, database.HandleDBError(err, "courseGroup")
	}
	return courseGroup.Group, nil
}

// Creation will be done by accessing the DB manually now
// func CreateCourseGroup(db *gorm.DB, courseGroup *CourseGroup) error {
// 	err := db.Create(courseGroup).Error
// 	if err != nil {
// 		return database.HandleDBError(err, "courseGroup")
// 	}
// 	return nil
// }
