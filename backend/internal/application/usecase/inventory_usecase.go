package usecase

import (
	"context"
	"time"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/response"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"
	errordomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/error"
	policydomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/policy"
	repositoryinterface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"
	servicedomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/service"
)

type InventoryUseCaseInterface interface {
	Reserve(
		ctx context.Context,
		req *request.ReserveInventoryRequest,
	) (*response.ReserveInventoryResponse, error)

	Confirm(
		ctx context.Context,
		req *request.ConfirmReservationRequest,
	) (*response.ConfirmReservationResponse, error)

	GetStock(
		ctx context.Context,
		itemID string,
	) (*response.InventoryResponse, error)

	ListInventory(
		ctx context.Context,
	) (*response.InventoryListResponse, error)
}

type InventoryUseCase struct {
	uow repositoryinterface.UnitOfWork

	uuidGenerator servicedomain.UUIDGenerator
}

func NewInventoryUseCase(
	uow repositoryinterface.UnitOfWork,
	uuidGenerator servicedomain.UUIDGenerator,
) InventoryUseCaseInterface {
	return &InventoryUseCase{
		uow:           uow,
		uuidGenerator: uuidGenerator,
	}
}

func (u *InventoryUseCase) Reserve(
	ctx context.Context,
	req *request.ReserveInventoryRequest,
) (*response.ReserveInventoryResponse, error) {

	var reservedInventory *entity.Inventory
	var reservation *entity.Reservation

	err := u.uow.WithTransaction(
		ctx,
		func(txUow repositoryinterface.UnitOfWork) error {

			inventory, err := txUow.
				InventoryRepository().
				FindByItemIDForUpdate(
					ctx,
					req.ItemID,
				)

			if err != nil {
				return err
			}

			if inventory == nil {
				return errordomain.ErrInventoryNotFound
			}

			if inventory.AvailableStock() < int(req.Quantity) {
				return errordomain.ErrInsufficientStock
			}

			now := time.Now().UTC()

			inventory.ReservedStock += int(req.Quantity)
			inventory.UpdatedAt = now

			err = txUow.
				InventoryRepository().
				Update(
					ctx,
					inventory,
				)

			if err != nil {
				return err
			}

			reservation = &entity.Reservation{
				ReservationID: u.uuidGenerator.New(),
				InventoryID:   inventory.ID,
				UserID:        req.UserID,
				Quantity:      int(req.Quantity),
				Status:        entity.ReservationStatusActive,
				ExpiresAt:     now.Add(policydomain.ReservationExpirationMinutes * time.Minute),
				CreatedAt:     now,
				UpdatedAt:     now,
			}

			reservedInventory = inventory
			err = txUow.
				ReservationRepository().
				Create(
					ctx,
					reservation,
				)

			if err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &response.ReserveInventoryResponse{
		ReservationID: reservation.ReservationID,
		ItemID:        reservedInventory.ItemID,
		Quantity:      reservation.Quantity,
		ExpiresAt:     reservation.ExpiresAt,
	}, nil
}

func (u *InventoryUseCase) Confirm(
	ctx context.Context,
	req *request.ConfirmReservationRequest,
) (*response.ConfirmReservationResponse, error) {

	var reservation *entity.Reservation

	err := u.uow.WithTransaction(
		ctx,
		func(txUow repositoryinterface.UnitOfWork) error {

			reservation, err := txUow.
				ReservationRepository().
				FindByReservationIDForUpdate(
					ctx,
					req.ReservationID,
				)

			if err != nil {
				return err
			}

			if reservation == nil {
				return errordomain.ErrReservationNotFound
			}

			if reservation.Status == entity.ReservationStatusConfirmed {
				return errordomain.ErrReservationAlreadyConfirmed
			}

			if reservation.Status == entity.ReservationStatusExpired {
				return errordomain.ErrReservationExpired
			}

			now := time.Now().UTC()

			if !reservation.ExpiresAt.After(now) {
				return errordomain.ErrReservationExpired
			}

			inventory, err := txUow.
				InventoryRepository().
				FindByIDForUpdate(
					ctx,
					reservation.InventoryID,
				)

			if err != nil {
				return err
			}

			if inventory == nil {
				return errordomain.ErrInventoryNotFound
			}

			inventory.TotalStock -= reservation.Quantity
			inventory.ReservedStock -= reservation.Quantity
			inventory.UpdatedAt = now

			err = txUow.
				InventoryRepository().
				Update(
					ctx,
					inventory,
				)

			if err != nil {
				return err
			}
			reservation.Status = entity.ReservationStatusConfirmed
			reservation.ConfirmedAt = &now
			reservation.UpdatedAt = now

			err = txUow.
				ReservationRepository().
				Update(
					ctx,
					reservation,
				)

			if err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &response.ConfirmReservationResponse{
		ReservationID: reservation.ReservationID,
		ConfirmedAt:   *reservation.ConfirmedAt,
	}, nil
}

func (u *InventoryUseCase) GetStock(
	ctx context.Context,
	itemID string,
) (*response.InventoryResponse, error) {

	inventory, err := u.uow.
		InventoryRepository().
		FindByItemID(
			ctx,
			itemID,
		)

	if err != nil {
		return nil, err
	}

	if inventory == nil {
		return nil, errordomain.ErrInventoryNotFound
	}

	return &response.InventoryResponse{
		ItemID:         inventory.ItemID,
		ItemName:       inventory.ItemName,
		TotalStock:     inventory.TotalStock,
		ReservedStock:  inventory.ReservedStock,
		AvailableStock: inventory.AvailableStock(),
	}, nil
}

func (u *InventoryUseCase) ListInventory(
	ctx context.Context,
) (*response.InventoryListResponse, error) {

	inventories, err := u.uow.
		InventoryRepository().
		FindAll(
			ctx,
		)

	if err != nil {
		return nil, err
	}

	items := make(
		[]response.InventoryResponse,
		0,
		len(inventories),
	)

	for _, inventory := range inventories {

		items = append(
			items,
			response.InventoryResponse{
				ItemID:         inventory.ItemID,
				ItemName:       inventory.ItemName,
				TotalStock:     inventory.TotalStock,
				ReservedStock:  inventory.ReservedStock,
				AvailableStock: inventory.AvailableStock(),
			},
		)
	}

	return &response.InventoryListResponse{
		Items: items,
	}, nil
}
