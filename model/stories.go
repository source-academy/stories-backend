package model

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type Story struct {
	StoryID      int       `json:"story_id"`
	UserID       int       `json:"user_id"`
	StoryContent string    `json:"story_content"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetAllStories() []Story {
	var stories []Story
	DB.Find(&stories)
	return stories
}

func GetStoryByID(storyID int) *Story {
	var story Story
	result := DB.First(&story, storyID)
	if result.Error != nil {
		return nil
	}

	return &story
}

func CreateStory(story Story) {
	DB.Create(&story)
}
