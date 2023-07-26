package userpermissions

import (
	"net/http"

	"github.com/source-academy/stories-backend/internal/auth"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
)

type RolePermission struct {
	Permission Permission
	Role       groupenums.Role
}

func (p RolePermission) IsAuthorized(r *http.Request) bool {
	userID, err := auth.GetUserIDFrom(r)
	if err != nil {
		return false
	}
	// TODO: Implement. Blocked by request context missing group info.
	_ = userID
	// role := getRole(userID, groupID)
	// return groupenums.IsRoleGreaterThan(role, p.Role)
	return false
}
