package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/sirupsen/logrus"
)

var jwks *jwk.Set

func getJWKS(endpointURL string) jwk.Set {
	// Singleton
	// TODO: Refresh every X minutes
	if jwks == nil {
		setJwkFromEndpoint(endpointURL)
	}
	return *jwks
}

func setJwkFromEndpoint(endpointURL string) {
	// Get JWK from endpoint
	logrus.Debugf("Using %s as JWKS source\n", endpointURL)
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
