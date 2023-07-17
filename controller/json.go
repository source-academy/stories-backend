package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func EncodeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}
