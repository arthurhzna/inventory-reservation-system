package error

import "errors"

var (
	ErrInventoryNotFound = errors.New(
		"inventory not found",
	)

	ErrInsufficientStock = errors.New(
		"insufficient stock",
	)
)
