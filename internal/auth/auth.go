package auth

import (
	"context"

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
	set, err := jwk.Fetch(context.Background(), endpointURL)
	if err != nil {
		logrus.WithError(err).Error("Failed to fetch JWK from endpoint")
		return
	}

	// Set JWK
	jwks = &set
}

// TODO: Test len, etc format as expected
