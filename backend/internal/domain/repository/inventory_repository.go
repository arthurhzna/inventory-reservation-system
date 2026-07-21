package repository

import (
	"context"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"
)

type InventoryRepository interface {
	FindAll(
		ctx context.Context,
	) ([]entity.Inventory, error)

	FindByID(
		ctx context.Context,
		id int64,
	) (*entity.Inventory, error)

	FindByIDForUpdate(
		ctx context.Context,
		id int64,
	) (*entity.Inventory, error)

	FindByItemID(
		ctx context.Context,
		itemID string,
	) (*entity.Inventory, error)

	FindByItemIDForUpdate(
		ctx context.Context,
		itemID string,
	) (*entity.Inventory, error)

	Update(
		ctx context.Context,
		inventory *entity.Inventory,
	) error
}
