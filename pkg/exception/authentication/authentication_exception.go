package authentication

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")

	ErrEmailNotFound    = errors.New("email not found")
	ErrPasswordAreWrong = errors.New("password are wrong")
)
