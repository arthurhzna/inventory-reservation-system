package entity

import "time"

const (
	ReservationStatusActive    = "ACTIVE"
	ReservationStatusConfirmed = "CONFIRMED"
	ReservationStatusExpired   = "EXPIRED"
)

type Reservation struct {
	ID            int64
	ReservationID string
	InventoryID   int64
	UserID        string
	Quantity      int
	Status        string
	ExpiresAt     time.Time
	ConfirmedAt   *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
