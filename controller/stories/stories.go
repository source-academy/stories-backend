package stories

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/model"
	storyparams "github.com/source-academy/stories-backend/params/stories"
	storyviews "github.com/source-academy/stories-backend/view/stories"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	stories := model.GetAllStories(db)
	controller.EncodeJSONResponse(w, storyviews.ListFrom(stories))
}

func HandleRead(w http.ResponseWriter, r *http.Request) {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		http.Error(w, "Invalid storyID", http.StatusBadRequest)
		return
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	story := model.GetStoryByID(db, storyID)
	controller.EncodeJSONResponse(w, storyviews.SingleFrom(story))
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var params storyparams.Create
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	storyModel := *params.ToModel()

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	model.CreateStory(db, &storyModel)
	controller.EncodeJSONResponse(w, storyviews.SingleFrom(storyModel))
	w.WriteHeader(http.StatusCreated)
}
