package error

import "errors"

var (
	ErrEmailAlreadyExist = errors.New(
		"email",
	)
	ErrUserNotFound = errors.New(
		"user",
	)
	ErrEmailNotFound = errors.New(
		"email",
	)
)
