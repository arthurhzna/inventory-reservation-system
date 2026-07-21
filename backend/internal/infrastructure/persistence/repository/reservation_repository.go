package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"

	repositoryiface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"

	dbtx "github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/dbtx"
)

type reservationRepository struct {
	db dbtx.DBTX
}

func NewReservationRepository(
	db dbtx.DBTX,
) repositoryiface.ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) Create(
	ctx context.Context,
	reservation *entity.Reservation,
) error {

	query := `
		INSERT INTO reservations (
			reservation_id,
			inventory_id,
			user_id,
			quantity,
			status,
			expires_at,
			confirmed_at,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		)
		RETURNING id
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		reservation.ReservationID,
		reservation.InventoryID,
		reservation.UserID,
		reservation.Quantity,
		reservation.Status,
		reservation.ExpiresAt,
		reservation.ConfirmedAt,
		reservation.CreatedAt,
		reservation.UpdatedAt,
	).Scan(
		&reservation.ID,
	)
}

func (r *reservationRepository) FindByReservationID(
	ctx context.Context,
	reservationID string,
) (*entity.Reservation, error) {

	query := `
		SELECT
			id,
			reservation_id,
			inventory_id,
			user_id,
			quantity,
			status,
			expires_at,
			confirmed_at,
			created_at,
			updated_at
		FROM reservations
		WHERE reservation_id = $1
	`

	var reservation entity.Reservation

	err := r.db.QueryRowxContext(
		ctx,
		query,
		reservationID,
	).Scan(
		&reservation.ID,
		&reservation.ReservationID,
		&reservation.InventoryID,
		&reservation.UserID,
		&reservation.Quantity,
		&reservation.Status,
		&reservation.ExpiresAt,
		&reservation.ConfirmedAt,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *reservationRepository) FindByReservationIDForUpdate(
	ctx context.Context,
	reservationID string,
) (*entity.Reservation, error) {

	query := `
		SELECT
			id,
			reservation_id,
			inventory_id,
			user_id,
			quantity,
			status,
			expires_at,
			confirmed_at,
			created_at,
			updated_at
		FROM reservations
		WHERE reservation_id = $1
		FOR UPDATE
	`

	var reservation entity.Reservation

	err := r.db.QueryRowxContext(
		ctx,
		query,
		reservationID,
	).Scan(
		&reservation.ID,
		&reservation.ReservationID,
		&reservation.InventoryID,
		&reservation.UserID,
		&reservation.Quantity,
		&reservation.Status,
		&reservation.ExpiresAt,
		&reservation.ConfirmedAt,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *reservationRepository) FindExpiredActive(
	ctx context.Context,
	limit int,
) ([]entity.Reservation, error) {

	query := `
		SELECT
			id,
			reservation_id,
			inventory_id,
			user_id,
			quantity,
			status,
			expires_at,
			confirmed_at,
			created_at,
			updated_at
		FROM reservations
		WHERE
			status = 'ACTIVE'
			AND expires_at <= NOW()
		ORDER BY expires_at ASC
		LIMIT $1
	`

	rows, err := r.db.QueryxContext(
		ctx,
		query,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reservations := make(
		[]entity.Reservation,
		0,
	)

	for rows.Next() {

		var reservation entity.Reservation

		err := rows.Scan(
			&reservation.ID,
			&reservation.ReservationID,
			&reservation.InventoryID,
			&reservation.UserID,
			&reservation.Quantity,
			&reservation.Status,
			&reservation.ExpiresAt,
			&reservation.ConfirmedAt,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(
			reservations,
			reservation,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *reservationRepository) Update(
	ctx context.Context,
	reservation *entity.Reservation,
) error {

	query := `
		UPDATE reservations
		SET
			status = $1,
			confirmed_at = $2,
			updated_at = $3
		WHERE id = $4
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		reservation.Status,
		reservation.ConfirmedAt,
		reservation.UpdatedAt,
		reservation.ID,
	)

	return err
}
