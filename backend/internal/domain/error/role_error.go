package error

import "errors"

var (
	ErrInvalidRole = errors.New(
		"role",
	)
	ErrForbidden = errors.New(
		"forbidden",
	)
)
