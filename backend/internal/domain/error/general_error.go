package error

import "errors"

var (
	ErrInvalidCredential = errors.New(
		"invalid credential",
	)
)
