package auth

import (
	"errors"
	"net/url"

	"github.com/source-academy/stories-backend/internal/database"
	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"github.com/source-academy/stories-backend/model"
	"gorm.io/gorm"
)

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
