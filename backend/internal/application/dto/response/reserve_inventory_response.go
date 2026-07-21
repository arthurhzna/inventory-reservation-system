package response

import "time"

type ReserveInventoryResponse struct {
	ReservationID string    `json:"reservation_id"`
	ItemID        string    `json:"item_id"`
	Quantity      int       `json:"quantity"`
	ExpiresAt     time.Time `json:"expires_at"`
}
