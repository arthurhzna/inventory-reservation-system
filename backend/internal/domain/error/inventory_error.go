package error

import "errors"

var (
	ErrInventoryNotFound = errors.New(
		"inventory item",
	)

	ErrInsufficientStock = errors.New(
		"insufficient stock",
	)
)
