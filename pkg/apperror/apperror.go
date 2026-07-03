// Package apperror carries an application error across the transport boundary:
// a stable Code + Message the client can rely on, plus the HTTP status to send.
// It keeps the underlying cause via Unwrap so errors.Is/As still reach the
// origin — domain and infrastructure layers keep using sentinel errors; a
// handler translates them into an AppError at the edge.
package apperror

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code           ErrorCode    `json:"code"`
	Message        ErrorMessage `json:"message"`
	HTTPStatusCode int          `json:"-"`
	Err            error        `json:"-"` // underlying cause, if any
}

// Error reports the message, appending the cause when one is attached.
func (e AppError) Error() string {
	if e.Err != nil {
		return string(e.Message) + ": " + e.Err.Error()
	}
	return string(e.Message)
}

// Unwrap exposes the cause so errors.Is / errors.As reach the origin error.
func (e AppError) Unwrap() error { return e.Err }

// WithMessage returns a copy carrying a custom message (same code + status).
func (e AppError) WithMessage(msg string) AppError {
	e.Message = ErrorMessage(msg)
	return e
}

func NewErrorMessage(v string) ErrorMessage { return ErrorMessage(v) }

// IsAppError reports whether err is (or wraps) an AppError, returning it.
func IsAppError(err error) (e AppError, ok bool) {
	ok = errors.As(err, &e)
	return
}

// newAppError builds an AppError from a code + status. Any causes are folded
// into one via errors.Join, so every cause is retained (errors.Is/As reach them
// all) and passing none leaves Err nil. It backs every constructor below.
func newAppError(code ErrorCode, status int, cause ...error) AppError {
	return AppError{
		Code:           code,
		Message:        Message[code],
		HTTPStatusCode: status,
		Err:            errors.Join(cause...),
	}
}

// ── Generic constructors ────────────────────────────────────────────────────
// Each optionally takes an underlying cause: apperror.NewNotFound(err). They
// return the concrete AppError (which satisfies error) so callers and the
// mapper can read Code / HTTPStatusCode without a type assertion.

func NewBadRequest(cause ...error) AppError {
	return newAppError(ErrBadRequestCode, http.StatusBadRequest, cause...)
}

func NewInvalidInput(cause ...error) AppError {
	return newAppError(ErrInvalidInputCode, http.StatusUnprocessableEntity, cause...)
}

func NewUnauthorized(cause ...error) AppError {
	return newAppError(ErrUnauthorizedCode, http.StatusUnauthorized, cause...)
}

func NewForbidden(cause ...error) AppError {
	return newAppError(ErrForbiddenCode, http.StatusForbidden, cause...)
}

func NewNotFound(cause ...error) AppError {
	return newAppError(ErrNotFoundCode, http.StatusNotFound, cause...)
}

func NewConflict(cause ...error) AppError {
	return newAppError(ErrConflictCode, http.StatusConflict, cause...)
}

func NewInternal(cause ...error) AppError {
	return newAppError(ErrInternalCode, http.StatusInternalServerError, cause...)
}
