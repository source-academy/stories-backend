package database

import (
	"errors"
	"fmt"

	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"gorm.io/gorm"
)

func HandleDBError(err error, fromModel string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apierrors.ClientNotFoundError{
			Message: fmt.Sprintf("Cannot find requested %s", fromModel),
		}
	}
	return err
}
