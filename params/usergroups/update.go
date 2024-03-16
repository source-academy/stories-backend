package usergroupparams

import (
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"github.com/source-academy/stories-backend/model"
)

type UpdateRole struct {
	Role groupenums.Role `json:"role"`
}

func (params *UpdateRole) ToModel() *model.UserGroup {
	return &model.UserGroup{
		Role: params.Role,
	}
}
