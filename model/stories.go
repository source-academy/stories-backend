package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Story struct {
	gorm.Model
	AuthorID uint
	Author   User
	GroupID  uint
	Group    Group
	Title    string
	Content  string
	PinOrder *int // nil if not pinned
}

func GetAllStories(db *gorm.DB) ([]Story, error) {
	var stories []Story
	err := db.
		Preload(clause.Associations).
		Order("pin_order ASC NULLS LAST, title ASC, content ASC").
		Find(&stories).
		Error
	if err != nil {
		return stories, database.HandleDBError(err, "story")
	}
	return stories, nil
}

func GetStoryByID(db *gorm.DB, id int) (Story, error) {
	var story Story
	err := db.
		Preload(clause.Associations).
		First(&story, id).
		Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}

func CreateStory(db *gorm.DB, story *Story) error {
	err := db.
		Preload(clause.Associations).
		Create(story).
		// Get associated Author. See
		// https://github.com/go-gorm/gen/issues/618 on why
		// a separate .First() is needed.
		First(story).
		Error
	if err != nil {
		return database.HandleDBError(err, "story")
	}
	return nil
}

func UpdateStory(db *gorm.DB, storyID int, newStory *Story) error {
	// TODO: Possible restore functionality for soft-deleted stories?
	err := db.
		Transaction(func(tx *gorm.DB) error {
			var originalStory Story
			err := tx.
				Preload(clause.Associations).
				Where("id = ?", storyID).
				First(&originalStory).
				Error
			if err != nil {
				return database.HandleDBError(err, "story")
			}

			// Handle nullable fields when null
			if newStory.PinOrder == nil {
				err = tx.
					Preload(clause.Associations).
					Where("id = ?", storyID).
					Model(newStory).
					Update("pin_order", gorm.Expr("NULL")).
					Error
				if err != nil {
					return database.HandleDBError(err, "story")
				}
			}

			// Update remaining fields
			err = tx.
				Preload(clause.Associations).
				Where("id = ?", storyID).
				Updates(newStory).
				First(newStory).
				Error
			if err != nil {
				return database.HandleDBError(err, "story")
			}

			return nil
		})
	if err != nil {
		return database.HandleDBError(err, "story")
	}

	return nil
}

func DeleteStory(db *gorm.DB, storyID int) (Story, error) {
	var story Story
	err := db.
		Preload(clause.Associations).
		Where("id = ?", storyID).
		First(&story). // store the value to be returned
		Delete(&story).
		Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}
