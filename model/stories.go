package model

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type StoryDB struct {
	gorm.Model
	AuthorID     uint
	StoryContent string
}

type Story struct {
	ID           uint   `json:"storyId"`
	AuthorID     uint   `json:"authorId"`
	StoryContent string `json:"storyContent"`
}

func MapStoryDBToStory(storyDB StoryDB) Story {
	story := Story{
		ID:           storyDB.ID,
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
