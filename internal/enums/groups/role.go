package groupenums

type Role uint

const (
	RoleStandard Role = iota
	RoleModerator
	RoleAdmin
)

// We cannot name it String() because it will conflict with the String() method
func (role Role) ToString() string {
	switch role {
	case RoleStandard:
		return "user"
	case RoleModerator:
		return "moderator"
	case RoleAdmin:
		return "admin"
	}
	return "unknown"
}

func RoleFromString(role string) (Role, bool) {
	switch role {
	case "user":
		return RoleStandard, true
	case "moderator":
		return RoleModerator, true
	case "admin":
		return RoleAdmin, true
	}
	// We fall back to standard role as default
	return RoleStandard, false
}
