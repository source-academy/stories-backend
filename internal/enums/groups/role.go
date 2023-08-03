package groupenums

type Role string

// We use database enums for roles, unlike login providers,
// hence the string type instead of int.
//
// As a side effect, we do not need a RoleFromString function
// as strings can be directly accepted into Role types.
const (
	RoleUnknown   Role = ""
	RoleStandard  Role = "member"
	RoleModerator Role = "moderator"
	RoleAdmin     Role = "admin"
)

func (role Role) IsValid() bool {
	return role == RoleStandard || role == RoleModerator || role == RoleAdmin
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
