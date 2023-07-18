package database

import (
	"testing"

	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	expectedUnchangedError = "Expected error to be unchanged"
	modelName              = "testingModelOneTwoThree"
)

func TestHandleDBError(t *testing.T) {
	t.Run("should return a ClientNotFoundError when gorm.ErrRecordNotFound is passed in", func(t *testing.T) {
		err := HandleDBError(gorm.ErrRecordNotFound, modelName)
		notFoundError, ok := err.(apierrors.ClientNotFoundError)
		assert.True(t, ok, "Expected error to be a ClientNotFoundError")
		assert.Contains(t, notFoundError.Error(), modelName, "Expected error to contain model name")
	})
	t.Run("should return any other errors unchanged", func(t *testing.T) {
		err := HandleDBError(gorm.ErrInvalidDB, modelName)
		assert.Equal(t, gorm.ErrInvalidDB, err, expectedUnchangedError)

		err = HandleDBError(gorm.ErrInvalidTransaction, modelName)
		assert.Equal(t, gorm.ErrInvalidTransaction, err, expectedUnchangedError)

		err = HandleDBError(gorm.ErrNotImplemented, modelName)
		assert.Equal(t, gorm.ErrNotImplemented, err, expectedUnchangedError)

		err = HandleDBError(gorm.ErrMissingWhereClause, modelName)
		assert.Equal(t, gorm.ErrMissingWhereClause, err, expectedUnchangedError)

		err = HandleDBError(gorm.ErrUnsupportedRelation, modelName)
		assert.Equal(t, gorm.ErrUnsupportedRelation, err, expectedUnchangedError)
	})
}
