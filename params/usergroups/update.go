package usergroupparams

import (
	"fmt"

	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"github.com/source-academy/stories-backend/model"
)

type UpdateRole struct {
	Role groupenums.Role `json:"role"`
}

func (params *UpdateRole) ToModel(userID uint) *model.UserGroup {
	return &model.UserGroup{
		UserID: userID,
		Role:   params.Role,
	}
}

func (params *UpdateRole) Validate() error {
	// Extra params won't do anything, e.g. authorID can't be changed.
	// TODO: Error on extra params?
	if !params.Role.IsValid() {
		return fmt.Errorf("Invalid role %s.", params.Role)
	}
	return nil
}
