package auth

import (
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/source-academy/stories-backend/internal/config"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	envutils "github.com/source-academy/stories-backend/internal/utils/env"
)

func MakeMiddlewareFrom(conf *config.Config) func(http.Handler) http.Handler {
	// Skip auth in development mode
	if conf.Environment == envutils.ENV_DEVELOPMENT {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	keySet := getJWKS(conf.JWKSEndpoint)
	key, ok := keySet.Key(0)
	if !ok {
		// Block all access if JWKS source is down, since we can't verify JWTs
		// TODO: Investigate if 500 is appropriate
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			})
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get JWT from request
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				apierrors.ServeHTTP(w, r, apierrors.ClientUnauthorizedError{
					Message: "Missing Authorization header.",
				})
				return
			}

			// Verify JWT
			toParse := authHeader[len("Bearer "):]
			token, err := jwt.ParseString(toParse, jwt.WithKey(jwa.RS256, key))
			if err != nil {
				apierrors.ServeHTTP(w, r, apierrors.ClientForbiddenError{
					Message: fmt.Sprintf("Failed to verify JWT: %s\n", err),
				})
				return
			}

			fmt.Println(token.Subject())

			// TODO: Get token subject (user information)

			// userData, err := url.ParseQuery(sub)
			// if err != nil {
			// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			// 	return
			// }

			// TODO: If JWT is valid, set user ID in context

			next.ServeHTTP(w, r)
		})
	}
}
