package model

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Story struct {
	gorm.Model
	AuthorID     uint
	StoryContent string
}

func GetAllStories(db *gorm.DB) []Story {
	var stories []Story
	db.Find(&stories)
	return stories
}

func GetStoryByID(db *gorm.DB, id int) Story {
	var story Story
	db.First(&story, id)
	return story
}

func CreateStory(db *gorm.DB, story *Story) {
	db.Create(story)
}
