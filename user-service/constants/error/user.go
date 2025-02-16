package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordInCorrect    = errors.New("password incorreact")
	ErrUsernameExist        = errors.New("username already exists")
	ErrEmailExist           = errors.New("email already exists")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrEmailExist,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
}
