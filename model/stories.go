package model

import "time"

type Story struct {
	StoryID      int       `json:"story_id"`
	UserID       int       `json:"user_id"`
	StoryContent string    `json:"story_content"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var stories = []Story{
	{StoryID: 1, UserID: 1, StoryContent: "Story 1 Content", CreatedAt: time.Now(), DeletedAt: time.Now(), UpdatedAt: time.Now()},
	{StoryID: 2, UserID: 2, StoryContent: "Story 2 Content", CreatedAt: time.Now(), DeletedAt: time.Now(), UpdatedAt: time.Now()},
}

func GetAllStories() []Story {
	return stories
}

func CreateStory(story Story) {
	stories = append(stories, story)
}
