package users

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/model"
	userparams "github.com/source-academy/stories-backend/params/users"
	userviews "github.com/source-academy/stories-backend/view/users"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	users := model.GetAllUsers(db)
	controller.EncodeJSONResponse(w, userviews.ListFrom(users))
}

func HandleRead(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	user := model.GetUserByID(db, userID)
	controller.EncodeJSONResponse(w, userviews.SingleFrom(user))
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var params userparams.Create
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userModel := *params.ToModel()

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	model.CreateUser(db, &userModel)
	controller.EncodeJSONResponse(w, userviews.SingleFrom(userModel))
	w.WriteHeader(http.StatusCreated)
}
