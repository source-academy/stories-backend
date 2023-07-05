package controller

import (
	"encoding/json"
	"net/http"

	"fmt"

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

func GetUser(w http.ResponseWriter, r *http.Request) {
	// userIDStr := r.URL.Query().Get("userID") // Assuming you want to retrieve it from the query parameter
	// fmt.Println("userID:", userIDStr)
	// userID, err := strconv.Atoi(userIDStr)
	// if err != nil {
	//  // Handle the error if the userID is not a valid integer
	//  http.Error(w, "Invalid userID", http.StatusBadRequest)
	//  return
	// }
	// user := model.GetUserByID(userID)
	user := model.GetUserByID(2)
	fmt.Println("userID:", 2)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(user); err != nil {
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
