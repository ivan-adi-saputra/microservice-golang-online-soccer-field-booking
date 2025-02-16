package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("databse server failed to execute")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
