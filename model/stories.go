// model/story.go
package model

type Story struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var stories = []Story{
	{ID: 1, Title: "Story 1", Author: "Username1"},
	{ID: 2, Title: "Story 2", Author: "Username2"},
}

func GetAllStories() []Story {
	return stories
}

func CreateStory(story Story) {
	stories = append(stories, story)
}
