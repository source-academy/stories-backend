package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"strconv"

	"github.com/source-academy/stories-backend/model"
	"github.com/source-academy/stories-backend/view"
)

func GetStories(w http.ResponseWriter, r *http.Request) {
	stories := model.GetAllStories()
	EncodeJSONResponse(w, stories)
}

func GetStory(w http.ResponseWriter, r *http.Request) {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		http.Error(w, "Invalid storyID", http.StatusBadRequest)
		return
	}
	story := model.GetStoryByID(storyID)
	EncodeJSONResponse(w, story)
}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	var story view.Story
	if err := json.NewDecoder(r.Body).Decode(&story); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.CreateStory(story)
	EncodeJSONResponse(w, &story)
	w.WriteHeader(http.StatusCreated)
}
