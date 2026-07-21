package response

type InventoryListResponse struct {
	Items []InventoryResponse `json:"items"`
}
