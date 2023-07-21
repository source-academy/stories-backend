package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/sirupsen/logrus"
)

var (
	// TODO: Move to config object
	endpointURL = os.Getenv("JWKS_ENDPOINT")
)

var jwks *jwk.Set

func getJWKS() jwk.Set {
	// Singleton
	// TODO: Refresh every X minutes
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("Failed to read JWK response body")
		return
	}

	set := jwk.NewSet()
	err = json.Unmarshal(body, &set)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse JWK")
		return
	}

	// Set JWK
	jwks = &set
}

// TODO: Test len, etc format as expected
