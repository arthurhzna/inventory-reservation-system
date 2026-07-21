package repository

import (
	"context"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"
)

type ReservationRepository interface {
	Create(
		ctx context.Context,
		reservation *entity.Reservation,
	) error

	FindByReservationID(
		ctx context.Context,
		reservationID string,
	) (*entity.Reservation, error)

	FindByReservationIDForUpdate(
		ctx context.Context,
		reservationID string,
	) (*entity.Reservation, error)

	FindExpiredActive(
		ctx context.Context,
		limit int,
	) ([]entity.Reservation, error)

	Update(
		ctx context.Context,
		reservation *entity.Reservation,
	) error
}
