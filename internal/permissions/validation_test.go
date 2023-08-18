package permissions

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test_alwaysTrue struct {
	PermissionGroup
}

func (t test_alwaysTrue) IsAuthorized(r *http.Request) bool {
	return true
}

type test_alwaysFalse struct {
	PermissionGroup
}

func (t test_alwaysFalse) IsAuthorized(r *http.Request) bool {
	return false
}

func TestValidateAllOf(t *testing.T) {
	t.Run("should return true in the trivial case without any permissions", func(t *testing.T) {
		groups := []PermissionGroup{}
		allOf := AllOf{Groups: groups}
		assert.True(t, allOf.IsAuthorized(nil))
	})
	t.Run("should return true if all groups return true", func(t *testing.T) {
		groups := []PermissionGroup{
			test_alwaysTrue{},
			test_alwaysTrue{},
			test_alwaysTrue{},
		}
		allOf := AllOf{Groups: groups}
		assert.True(t, allOf.IsAuthorized(nil))
	})
	t.Run("should return false if any group returns false", func(t *testing.T) {
		groups := []PermissionGroup{
			test_alwaysTrue{},
			test_alwaysFalse{},
			test_alwaysTrue{},
		}
		allOf := AllOf{Groups: groups}
		assert.False(t, allOf.IsAuthorized(nil))
	})
}

func TestValidateAnyOf(t *testing.T) {
	t.Run("should return true in the trivial case without any permissions", func(t *testing.T) {
		groups := []PermissionGroup{}
		anyOf := AnyOf{Groups: groups}
		assert.True(t, anyOf.IsAuthorized(nil))
	})
	t.Run("should return false if all groups return false", func(t *testing.T) {
		groups := []PermissionGroup{
			test_alwaysFalse{},
			test_alwaysFalse{},
			test_alwaysFalse{},
		}
		anyOf := AnyOf{Groups: groups}
		assert.False(t, anyOf.IsAuthorized(nil))
	})
	t.Run("should return true if any group returns true", func(t *testing.T) {
		groups := []PermissionGroup{
			test_alwaysFalse{},
			test_alwaysTrue{},
			test_alwaysFalse{},
		}
		anyOf := AnyOf{Groups: groups}
		assert.True(t, anyOf.IsAuthorized(nil))
	})

}
