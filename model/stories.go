package model

import (
	"github.com/source-academy/stories-backend/view"
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

func MapStoryDBToStory(storyDB StoryDB) view.Story {
	story := view.Story{
		StoryID:      storyDB.StoryID,
		AuthorID:     storyDB.AuthorID,
		StoryContent: storyDB.StoryContent,
	}
	return story
}

func GetAllStories() []view.Story {
	var stories []view.Story
	DB.Find(&stories)
	return stories
}

func GetStoryByID(storyID int) *view.Story {
	var story view.Story
	DB.First(&story, storyID)
	return &story
}

func CreateStory(story view.Story) {
	DB.Create(&story)
}
