package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	AuthorID uint
	Author   User
	Title    string
	Content  string
	PinOrder *int // nil if not pinned
}

var (
	preloadAssociations = func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author")
	}
)

func GetAllStories(db *gorm.DB) ([]Story, error) {
	var stories []Story
	err := db.
		Scopes(preloadAssociations).
		// TODO: Abstract out
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
		Scopes(preloadAssociations).
		First(&story, id).
		Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}

func CreateStory(db *gorm.DB, story *Story) error {
	err := db.
		Scopes(preloadAssociations).
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
				Scopes(preloadAssociations).
				Where("id = ?", storyID).
				First(&originalStory).
				Error
			if err != nil {
				return database.HandleDBError(err, "story")
			}

			err = tx.
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
		Scopes(preloadAssociations).
		Where("id = ?", storyID).
		First(&story). // store the value to be returned
		Delete(&story).
		Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}
