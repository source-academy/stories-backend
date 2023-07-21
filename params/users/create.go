package userparams

import (
	"errors"

	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"github.com/source-academy/stories-backend/model"
)

const (
	loginParamGitHub = "github"
	loginParamNUSNET = "luminus" // for legacy reasons
)

type Create struct {
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

// TODO: Add some validation
func (params *Create) Validate() error {
	// Validate login provider is one of the ones supported
	switch params.LoginProvider {
	case loginParamGitHub, loginParamNUSNET:
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
	case loginParamGitHub:
		return userenums.LoginProviderGitHub
	case loginParamNUSNET:
		return userenums.LoginProviderNUSNET
	default:
		// Should never happen as we previously validated the provider
		// in the Validate() function, thus ok to panic
		panic(errors.New("Illegal path - invalid provider"))
	}
}
