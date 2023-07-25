package database

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"gorm.io/gorm"
)

var (
	errUniqueIndexViolation = pgconn.PgError{
		Code: "23505",
	}
)

// Checks if err is a *pgconn.PgError with the given errcode.
// If yes, returns the error as *pgconn.PgError and true.
// Otherwise, returns _ (may be nil or a valid pointer) and false.
func isPGError(err error, errcode pgconn.PgError) (*pgconn.PgError, bool) {
	// fmt.Println(errcode.Name())
	var pgerr *pgconn.PgError
	ok := errors.As(err, &pgerr)
	return pgerr, ok && pgerr.Code == errcode.Code
}

func HandleDBError(err error, fromModel string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apierrors.ClientNotFoundError{
			Message: fmt.Sprintf("Cannot find requested %s.", fromModel),
		}
	}
	if pgErr, ok := isPGError(err, errUniqueIndexViolation); ok {
		return apierrors.ClientConflictError{
			Message: fmt.Sprintf("%s %s", fromModel, pgErr.Detail),
		}
	}
	// TODO: Handle more types of errors
	return err
}
