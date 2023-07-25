package userpermissions

import (
	"net/http"

	"github.com/source-academy/stories-backend/internal/auth"
	// groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
)

type RolePermission struct {
	Permission Permission
	// Role       groupenums.Role
}

func (p RolePermission) IsAuthorized(r *http.Request) bool {
	userID, err := auth.GetUserIDFrom(r)
	if err != nil {
		return false
	}
	// TODO: Implement. Blocked by request context missing group info.
	_ = userID
	// role := getRole(userID, groupID)
	// return isRoleGreaterThan(role, p.Role)
	return false
}

// TODO: Move comparison function to role enum
// func isRoleGreaterThan(role1, role2 groupenums.Role) bool {
// 	switch role1 {
// 	case groupenums.RoleAdmin:
// 		return true
// 	case groupenums.RoleModerator:
// 		return role2 != groupenums.RoleAdmin
// 	case groupenums.RoleStandard:
// 		return role2 == groupenums.RoleStandard
// 	}
// 	return false
// }
