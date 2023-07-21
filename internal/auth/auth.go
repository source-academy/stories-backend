package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// TODO: Move to config object
	endpointURL = os.Getenv("JWKS_ENDPOINT")
)

var jwks *map[string]interface{}

func GetJWKS() map[string]interface{} {
	// Singleton
	if jwks == nil {
		setJwkFromEndpoint()
	}
	return *jwks
}

func setJwkFromEndpoint() {
	// Get JWK from endpoint
	fmt.Println(endpointURL)
	resp, err := http.Get(endpointURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to get JWK from endpoint")
		return
	}
	defer resp.Body.Close()

	// Parse JWK
	var decoded map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse JWK")
		return
	}

	// Set JWK
	jwks = &decoded
}

// TODO: Test len, etc format as expected
