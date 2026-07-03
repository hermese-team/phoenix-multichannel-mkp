package apperror

// ErrorCode is a stable, machine-readable code returned to API clients.
// ErrorMessage is the human-readable text that goes with it.
type (
	ErrorCode    int
	ErrorMessage string
)

// Codes are grouped so the leading digit hints at the class of failure:
//
//	4xxx — client error (maps to 4xx HTTP)
//	5xxx — server error (maps to 5xx HTTP)
const (
	ErrBadRequestCode   ErrorCode = 4000
	ErrInvalidInputCode ErrorCode = 4001
	ErrUnauthorizedCode ErrorCode = 4010
	ErrForbiddenCode    ErrorCode = 4030
	ErrNotFoundCode     ErrorCode = 4040
	ErrConflictCode     ErrorCode = 4090

	ErrInternalCode ErrorCode = 5000
)

// Message maps each code to its default message. A constructor may override the
// text (e.g. to interpolate a field name) while keeping the same code.
var Message = map[ErrorCode]ErrorMessage{
	ErrBadRequestCode:   "Bad request",
	ErrInvalidInputCode: "Invalid input",
	ErrUnauthorizedCode: "Unauthorized",
	ErrForbiddenCode:    "Forbidden",
	ErrNotFoundCode:     "Not found",
	ErrConflictCode:     "Conflict",

	ErrInternalCode: "Internal server error",
}
