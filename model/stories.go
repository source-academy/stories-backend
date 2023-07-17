package model

import (
	storyviews "github.com/source-academy/stories-backend/view/stories"
	"gorm.io/gorm"
)

var DB *gorm.DB

type StoryDB struct {
	gorm.Model
	AuthorID uint
	Content  string
}

func MapStoryDBToStory(storyDB StoryDB) storyviews.View {
	story := storyviews.View{
		ID:           storyDB.ID,
		AuthorID:     storyDB.AuthorID,
		StoryContent: storyDB.Content,
	}
	return story
}

func GetAllStories() []storyviews.View {
	var stories []storyviews.View
	DB.Find(&stories)
	return stories
}

func GetStoryByID(storyID int) *storyviews.View {
	var story storyviews.View
	DB.First(&story, storyID)
	return &story
}

func CreateStory(story storyviews.View) {
	DB.Create(&story)
}
