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

// TODO: remove if not used. since string can be accepted to a role type
func RoleFromString(role string) (Role, bool) {
	switch role {
	case "member":
		return RoleStandard, true
	case "moderator":
		return RoleModerator, true
	case "admin":
		return RoleAdmin, true
	}
	// We fall back to standard role as default
	return RoleStandard, false
}

// Implements the Scanner interface
// func (role *Role) Scan(value interface{}) error {
// 	str, ok := value.(string)
// 	if !ok {
// 		return errors.New("failed to scan role")
// 	}
// 	*role = Role(str)
// 	return nil
// }

// // Imokements the Valuer interface
// func (role Role) Value() (driver.Value, error) {
// 	switch role {
// 	case RoleStandard:
// 		return "member", nil
// 	case RoleModerator:
// 		return "moderator", nil
// 	case RoleAdmin:
// 		return "admin", nil
// 	}
// 	return "member", errors.New("unknown role, fall back to member.")
// }

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
