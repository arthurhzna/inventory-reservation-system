package response

import (
	errordomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/error"
)

func MapError(err error) error {
	switch err {

	case errordomain.ErrInventoryNotFound:
		return NewNotFoundError(err)

	case errordomain.ErrReservationNotFound:
		return NewNotFoundError(err)

	case errordomain.ErrInsufficientStock:
		return NewConflictError(err)

	case errordomain.ErrReservationExpired:
		return NewConflictError(err)

	case errordomain.ErrReservationAlreadyConfirmed:
		return NewConflictError(err)

	case errordomain.ErrInvalidReservationStatus:
		return NewConflictError(err)
	}

	return err
}
