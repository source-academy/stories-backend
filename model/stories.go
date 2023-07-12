package model

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type StoryDB struct {
	StoryID      int
	AuthorID     int
	StoryContent string
	CreatedAt    time.Time
	DeletedAt    time.Time
	UpdatedAt    time.Time
}

type Story struct {
	StoryID      int    `json:"storyId"`
	AuthorID     int    `json:"authorId"`
	StoryContent string `json:"storyContent"`
}

func MapStoryDBToStory(storyDB StoryDB) Story {
	story := Story{
		StoryID:      storyDB.StoryID,
		AuthorID:     storyDB.AuthorID,
		StoryContent: storyDB.StoryContent,
	}
	return story
}

func GetAllStories() []Story {
	var stories []Story
	DB.Find(&stories)
	return stories
}

func GetStoryByID(storyID int) *Story {
	var story Story
	DB.First(&story, storyID)
	return &story
}

func CreateStory(story Story) {
	DB.Create(&story)
}
