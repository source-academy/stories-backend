package controller

import (
	"fmt"
	"net/http"
)

const (
	welcomeMessage = "Hello from Source Academy Stories!"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, welcomeMessage)
}
