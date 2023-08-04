package permissions

import (
	"net/http"
)

type PermissionGroup interface {
	IsAuthorized(*http.Request) bool
}

type AllOf struct {
	Groups []PermissionGroup
}

func (a AllOf) IsAuthorized(r *http.Request) bool {
	for _, group := range a.Groups {
		if !group.IsAuthorized(r) {
			return false
		}
	}
	return true
}

type AnyOf struct {
	Groups []PermissionGroup
}

func (a AnyOf) IsAuthorized(r *http.Request) bool {
	for _, group := range a.Groups {
		if group.IsAuthorized(r) {
			return true
		}
	}
	return false
}
