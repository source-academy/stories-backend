package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/model"
)

func GetStories(w http.ResponseWriter, r *http.Request) {
	stories := model.GetAllStories()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&stories); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func GetStory(w http.ResponseWriter, r *http.Request) {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		http.Error(w, "Invalid storyID", http.StatusBadRequest)
		return
	}
	story := model.GetStoryByID(storyID)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(story); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	var story model.Story
	if err := json.NewDecoder(r.Body).Decode(&story); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.CreateStory(story)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&story); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
}
