package userparams

import (
	"errors"
	"fmt"

	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"github.com/source-academy/stories-backend/model"
)

type Create struct {
	Name          string `json:"name"`
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

type BatchCreate struct {
	Users []Create `json:"users"`
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
		return fmt.Errorf("Invalid login provider %s.", params.LoginProvider)
	}

	return nil
}

func (params *Create) ToModel() *model.User {
	provider, ok := userenums.LoginProviderFromString(params.LoginProvider)
	if !ok {
		// Should never happen as we previously validated the provider
		// in the Validate() function, thus ok to panic
		panic(errors.New("Illegal path - invalid provider"))
	}
	return &model.User{
		FullName:      params.Name,
		Username:      params.Username,
		LoginProvider: provider,
	}
}
