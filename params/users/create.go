package userparams

import (
	"errors"

	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"github.com/source-academy/stories-backend/model"
)

type Create struct {
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

// TODO: Add some validation
func (params *Create) Validate() error {
	// Validate login provider is one of the ones supported AND allowed
	switch params.LoginProvider {
	case
		// Allowed login providers for now
		// TODO: Allow more login providers
		userenums.LoginProviderNUSNET.ToString():
		break
	default:
		return errors.New("Invalid login provider")
	}

	return nil
}

func (params *Create) ToModel() *model.User {
	return &model.User{
		Username:      params.Username,
		LoginProvider: convertProvider(params.LoginProvider),
	}
}

func convertProvider(provider string) userenums.LoginProvider {
	switch provider {
	case userenums.LoginProviderNUSNET.ToString():
		return userenums.LoginProviderNUSNET
	default:
		// Should never happen as we previously validated the provider
		// in the Validate() function, thus ok to panic
		panic(errors.New("Illegal path - invalid provider"))
	}
}
