package response

type InventoryResponse struct {
	ItemID         string `json:"item_id"`
	ItemName       string `json:"item_name"`
	TotalStock     int    `json:"total_stock"`
	ReservedStock  int    `json:"reserved_stock"`
	AvailableStock int    `json:"available_stock"`
}
