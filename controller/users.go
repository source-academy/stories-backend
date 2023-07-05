package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(users); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.CreateUser(user)
	w.WriteHeader(http.StatusCreated)
}
