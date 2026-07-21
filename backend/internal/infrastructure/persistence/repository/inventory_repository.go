package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"

	repositoryiface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"

	dbtx "github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/dbtx"
)

type inventoryRepository struct {
	db dbtx.DBTX
}

func NewInventoryRepository(
	db dbtx.DBTX,
) repositoryiface.InventoryRepository {
	return &inventoryRepository{
		db: db,
	}
}

func (r *inventoryRepository) FindAll(
	ctx context.Context,
) ([]entity.Inventory, error) {

	query := `
		SELECT
			id,
			item_id,
			item_name,
			total_stock,
			reserved_stock,
			created_at,
			updated_at
		FROM inventories
		ORDER BY item_name ASC
	`

	rows, err := r.db.QueryxContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	inventories := make(
		[]entity.Inventory,
		0,
	)

	for rows.Next() {

		var inventory entity.Inventory

		err := rows.Scan(
			&inventory.ID,
			&inventory.ItemID,
			&inventory.ItemName,
			&inventory.TotalStock,
			&inventory.ReservedStock,
			&inventory.CreatedAt,
			&inventory.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		inventories = append(
			inventories,
			inventory,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inventories, nil
}

func (r *inventoryRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entity.Inventory, error) {

	query := `
		SELECT
			id,
			item_id,
			item_name,
			total_stock,
			reserved_stock,
			created_at,
			updated_at
		FROM inventories
		WHERE id = $1
	`

	var inventory entity.Inventory

	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&inventory.ID,
		&inventory.ItemID,
		&inventory.ItemName,
		&inventory.TotalStock,
		&inventory.ReservedStock,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) FindByIDForUpdate(
	ctx context.Context,
	id int64,
) (*entity.Inventory, error) {

	query := `
		SELECT
			id,
			item_id,
			item_name,
			total_stock,
			reserved_stock,
			created_at,
			updated_at
		FROM inventories
		WHERE id = $1
		FOR UPDATE
	`

	var inventory entity.Inventory

	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&inventory.ID,
		&inventory.ItemID,
		&inventory.ItemName,
		&inventory.TotalStock,
		&inventory.ReservedStock,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) FindByItemID(
	ctx context.Context,
	itemID string,
) (*entity.Inventory, error) {

	query := `
		SELECT
			id,
			item_id,
			item_name,
			total_stock,
			reserved_stock,
			created_at,
			updated_at
		FROM inventories
		WHERE item_id = $1
	`

	var inventory entity.Inventory

	err := r.db.QueryRowxContext(
		ctx,
		query,
		itemID,
	).Scan(
		&inventory.ID,
		&inventory.ItemID,
		&inventory.ItemName,
		&inventory.TotalStock,
		&inventory.ReservedStock,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) FindByItemIDForUpdate(
	ctx context.Context,
	itemID string,
) (*entity.Inventory, error) {

	query := `
		SELECT
			id,
			item_id,
			item_name,
			total_stock,
			reserved_stock,
			created_at,
			updated_at
		FROM inventories
		WHERE item_id = $1
		FOR UPDATE
	`

	var inventory entity.Inventory

	err := r.db.QueryRowxContext(
		ctx,
		query,
		itemID,
	).Scan(
		&inventory.ID,
		&inventory.ItemID,
		&inventory.ItemName,
		&inventory.TotalStock,
		&inventory.ReservedStock,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) Update(
	ctx context.Context,
	inventory *entity.Inventory,
) error {

	query := `
		UPDATE inventories
		SET
			total_stock = $1,
			reserved_stock = $2,
			updated_at = $3
		WHERE id = $4
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		inventory.TotalStock,
		inventory.ReservedStock,
		inventory.UpdatedAt,
		inventory.ID,
	)

	return err
}
