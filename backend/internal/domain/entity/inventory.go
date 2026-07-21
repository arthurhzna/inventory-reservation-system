package entity

import "time"

type Inventory struct {
	ID            int64
	ItemID        string
	ItemName      string
	TotalStock    int
	ReservedStock int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (i Inventory) AvailableStock() int {
	return i.TotalStock - i.ReservedStock
}
