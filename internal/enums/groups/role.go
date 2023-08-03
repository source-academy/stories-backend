package groupenums

type Role string

const (
	RoleUnknown   Role = ""
	RoleStandard  Role = "member"
	RoleModerator Role = "moderator"
	RoleAdmin     Role = "admin"
)

// We cannot name it String() because it will conflict with the String() method
// TODO: remove if not used. might not be needed unless we need to handle the error case
func (role Role) ToString() string {
	switch role {
	case RoleStandard:
		return "member"
	case RoleModerator:
		return "moderator"
	case RoleAdmin:
		return "admin"
	}
	return "unknown"
}

func IsRoleGreaterThan(role1, role2 Role) bool {
	switch role1 {
	case RoleAdmin:
		return true
	case RoleModerator:
		return role2 != RoleAdmin
	case RoleStandard:
		return role2 == RoleStandard
	}
	return false
}
