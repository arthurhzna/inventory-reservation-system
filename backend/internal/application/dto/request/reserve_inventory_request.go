package request

type ReserveInventoryRequest struct {
	UserID   string `json:"user_id"`
	ItemID   string `json:"item_id"`
	Quantity int64  `json:"quantity"`
}
