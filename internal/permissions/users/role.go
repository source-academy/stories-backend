package userpermissions

import (
	"net/http"

	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"github.com/source-academy/stories-backend/internal/usergroups"
)

type RolePermission struct {
	Permission Permission
	Role       groupenums.Role
}

func (p RolePermission) IsAuthorized(r *http.Request) bool {
	role, err := usergroups.GetRoleFrom(r)
	if err != nil {
		return false
	}
	return groupenums.IsRoleGreaterThan(*role, p.Role)
}


