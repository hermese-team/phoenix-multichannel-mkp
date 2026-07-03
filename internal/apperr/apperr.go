// Package apperr is the single place that knows how this application's sentinel
// errors translate into a transport-facing apperror.AppError (code + HTTP
// status). Domain and infrastructure keep returning internal.Err*; the HTTP
// layer calls From to resolve whatever bubbled up.
package apperr

import (
	"errors"

	"github.com/okdev/marketplace-sync/internal"
	"github.com/okdev/marketplace-sync/pkg/apperror"
)

// From resolves any error into an AppError:
//   - an error that already is (or wraps) an AppError is returned unchanged,
//   - a known internal sentinel maps to its matching code + status,
//   - anything else falls back to a 500 (the cause is preserved via Unwrap).
func From(err error) apperror.AppError {
	if ae, ok := apperror.IsAppError(err); ok {
		return ae
	}
	switch {
	case errors.Is(err, internal.ErrNotFound):
		return apperror.NewNotFound(err)
	case errors.Is(err, internal.ErrConflict):
		return apperror.NewConflict(err)
	case errors.Is(err, internal.ErrUnauthorized):
		return apperror.NewUnauthorized(err)
	case errors.Is(err, internal.ErrInvalid):
		return apperror.NewInvalidInput(err)
	default:
		return apperror.NewInternal(err)
	}
}
