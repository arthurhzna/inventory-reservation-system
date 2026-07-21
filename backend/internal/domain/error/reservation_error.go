package error

import "errors"

var (
	ErrReservationNotFound = errors.New(
		"reservation not found",
	)

	ErrReservationExpired = errors.New(
		"reservation has expired",
	)

	ErrReservationAlreadyConfirmed = errors.New(
		"reservation has already been confirmed",
	)

	ErrInvalidReservationStatus = errors.New(
		"invalid reservation status",
	)
)
