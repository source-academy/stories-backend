package model

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type Story struct {
	StoryID      int       `json:"storyId"`
	AuthorID     int       `json:"authorId"`
	StoryContent string    `json:"storyContent"`
	CreatedAt    time.Time `json:"createdAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
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
