package response

import "time"

type ConfirmReservationResponse struct {
	ReservationID string    `json:"reservation_id"`
	ConfirmedAt   time.Time `json:"confirmed_at"`
}
