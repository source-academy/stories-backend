package auth

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	envutils "github.com/source-academy/stories-backend/internal/utils/env"
	"github.com/source-academy/stories-backend/model"
	"gorm.io/gorm"
)

const (
	invalidTokenSubjectMessage = "Invalid user."
	usernameKey                = "username"
	loginProviderKey           = "provider"
)

// MakeMiddlewareFrom returns a middleware that verifies JWTs from the Authorization
// header of incoming requests. If the JWT is valid, the user ID is set in the
// request context.
//
// It must be called after the DB middleware, since it depends on the DB connection.
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

			// No error due to precondition of DB middleware being called first.
			// Just to be safe, we check for error anyway.
			db, err := database.GetDBFrom(r)
			if err != nil {
				// Will be caught by apierrors as 500 Internal Server Error
				apierrors.ServeHTTP(w, r, errors.New("Failed to get DB connection."))
				return
			}

			user, err := validateAndGetUser(token.Subject(), db)
			if err != nil {
				// Intentionally override any status code with 403
				apierrors.ServeHTTP(w, r, apierrors.ClientForbiddenError{
					Message: err.Error(),
				})
				return
			}

			// If JWT is valid, set user ID in context
			r = injectUserIDToContext(r, int(user.ID))

			next.ServeHTTP(w, r)
		})
	}
}

func validateAndGetUser(queryString string, db *gorm.DB) (*model.User, error) {
	// Validate valid query string
	userData, err := url.ParseQuery(queryString)
	if err != nil {
		return nil, errors.New(invalidTokenSubjectMessage)
	}

	// Validate required fields
	requiredFields := []string{usernameKey, loginProviderKey}
	for _, field := range requiredFields {
		if !userData.Has(field) {
			return nil, errors.New(invalidTokenSubjectMessage)
		}
	}

	// Validate login provider
	provider, ok := userenums.LoginProviderFromString(userData.Get(loginProviderKey))
	if !ok {
		// Invalid/unsupported login provider
		return nil, errors.New(invalidTokenSubjectMessage)
	}

	// Validate user
	user := model.User{
		Username:      userData.Get(usernameKey),
		LoginProvider: provider,
	}
	var dbUser model.User
	err = db.Where(&user).First(&dbUser).Error
	if err != nil {
		return nil, database.HandleDBError(err, "user")
	}

	return &dbUser, nil
}
