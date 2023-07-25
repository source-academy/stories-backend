package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	AuthorID uint
	Title    string
	Content  string
	PinOrder *int // nil if not pinned
}

func GetAllStories(db *gorm.DB) ([]Story, error) {
	var stories []Story
	err := db.
		// TODO: Abstract out
		Order("pin_order ASC NULLS LAST, title ASC, content ASC").
		Find(&stories).Error
	if err != nil {
		return stories, database.HandleDBError(err, "story")
	}
	return stories, nil
}

func GetStoryByID(db *gorm.DB, id int) (Story, error) {
	var story Story
	err := db.First(&story, id).Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}

func CreateStory(db *gorm.DB, story *Story) error {
	err := db.Create(story).Error
	if err != nil {
		return database.HandleDBError(err, "story")
	}
	return nil
}
