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

func HandleRead(w http.ResponseWriter, r *http.Request) error {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return err
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	user, err := model.GetUserByID(db, userID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	controller.EncodeJSONResponse(w, userviews.SingleFrom(user))
	return nil
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var params userparams.Create
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := params.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
