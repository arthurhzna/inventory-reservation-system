package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"
	errordomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/error"
	repositoryinterface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"
)

func TestInventoryUseCaseReserve(t *testing.T) {
	t.Run("success reserves stock and creates reservation", func(t *testing.T) {
		ctx := context.Background()
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			ItemName:      "Keyboard",
			TotalStock:    10,
			ReservedStock: 2,
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byItemID[inventory.ItemID] = inventory
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{value: "reservation-001"})

		before := time.Now().UTC()
		resp, err := uc.Reserve(ctx, &request.ReserveInventoryRequest{
			UserID:   "user-001",
			ItemID:   "SKU-001",
			Quantity: 3,
		})
		after := time.Now().UTC()

		if err != nil {
			t.Fatalf("Reserve() error = %v", err)
		}
		if resp.ReservationID != "reservation-001" {
			t.Fatalf("ReservationID = %q, want %q", resp.ReservationID, "reservation-001")
		}
		if resp.ItemID != "SKU-001" {
			t.Fatalf("ItemID = %q, want %q", resp.ItemID, "SKU-001")
		}
		if resp.Quantity != 3 {
			t.Fatalf("Quantity = %d, want %d", resp.Quantity, 3)
		}
		if resp.ExpiresAt.Before(before.Add(5*time.Minute)) || resp.ExpiresAt.After(after.Add(5*time.Minute)) {
			t.Fatalf("ExpiresAt = %v, want around now + 5 minutes", resp.ExpiresAt)
		}
		if inventory.ReservedStock != 5 {
			t.Fatalf("ReservedStock = %d, want %d", inventory.ReservedStock, 5)
		}
		if uow.inventoryRepo.updateCalls != 1 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 1)
		}
		if uow.reservationRepo.createCalls != 1 {
			t.Fatalf("reservation create calls = %d, want %d", uow.reservationRepo.createCalls, 1)
		}

		created := uow.reservationRepo.created[0]
		if created.ReservationID != "reservation-001" ||
			created.InventoryID != inventory.ID ||
			created.UserID != "user-001" ||
			created.Quantity != 3 ||
			created.Status != entity.ReservationStatusActive {
			t.Fatalf("created reservation = %+v", created)
		}
	})

	t.Run("returns not found when inventory does not exist", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{value: "reservation-001"})

		resp, err := uc.Reserve(ctx, &request.ReserveInventoryRequest{
			UserID:   "user-001",
			ItemID:   "missing",
			Quantity: 1,
		})

		if !errors.Is(err, errordomain.ErrInventoryNotFound) {
			t.Fatalf("Reserve() error = %v, want %v", err, errordomain.ErrInventoryNotFound)
		}
		if resp != nil {
			t.Fatalf("Reserve() response = %+v, want nil", resp)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
		if uow.reservationRepo.createCalls != 0 {
			t.Fatalf("reservation create calls = %d, want %d", uow.reservationRepo.createCalls, 0)
		}
	})

	t.Run("returns insufficient stock when available stock is too low", func(t *testing.T) {
		ctx := context.Background()
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    4,
			ReservedStock: 2,
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byItemID[inventory.ItemID] = inventory
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{value: "reservation-001"})

		resp, err := uc.Reserve(ctx, &request.ReserveInventoryRequest{
			UserID:   "user-001",
			ItemID:   "SKU-001",
			Quantity: 3,
		})

		if !errors.Is(err, errordomain.ErrInsufficientStock) {
			t.Fatalf("Reserve() error = %v, want %v", err, errordomain.ErrInsufficientStock)
		}
		if resp != nil {
			t.Fatalf("Reserve() response = %+v, want nil", resp)
		}
		if inventory.ReservedStock != 2 {
			t.Fatalf("ReservedStock = %d, want unchanged %d", inventory.ReservedStock, 2)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
		if uow.reservationRepo.createCalls != 0 {
			t.Fatalf("reservation create calls = %d, want %d", uow.reservationRepo.createCalls, 0)
		}
	})
}

func TestInventoryUseCaseConfirm(t *testing.T) {
	t.Run("success confirms reservation and deducts stock", func(t *testing.T) {
		ctx := context.Background()
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    10,
			ReservedStock: 4,
		}
		reservation := &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   inventory.ID,
			UserID:        "user-001",
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(time.Minute),
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byID[inventory.ID] = inventory
		uow.reservationRepo.byReservationID[reservation.ReservationID] = reservation
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "reservation-001",
		})

		if err != nil {
			t.Fatalf("Confirm() error = %v", err)
		}
		if resp.ReservationID != "reservation-001" {
			t.Fatalf("ReservationID = %q, want %q", resp.ReservationID, "reservation-001")
		}
		if resp.ConfirmedAt.IsZero() {
			t.Fatal("ConfirmedAt is zero")
		}
		if inventory.TotalStock != 7 {
			t.Fatalf("TotalStock = %d, want %d", inventory.TotalStock, 7)
		}
		if inventory.ReservedStock != 1 {
			t.Fatalf("ReservedStock = %d, want %d", inventory.ReservedStock, 1)
		}
		if reservation.Status != entity.ReservationStatusConfirmed {
			t.Fatalf("Status = %q, want %q", reservation.Status, entity.ReservationStatusConfirmed)
		}
		if reservation.ConfirmedAt == nil {
			t.Fatal("ConfirmedAt = nil, want value")
		}
		if uow.inventoryRepo.updateCalls != 1 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 1)
		}
		if uow.reservationRepo.updateCalls != 1 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 1)
		}
	})

	t.Run("returns not found when reservation does not exist", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "missing",
		})

		if !errors.Is(err, errordomain.ErrReservationNotFound) {
			t.Fatalf("Confirm() error = %v, want %v", err, errordomain.ErrReservationNotFound)
		}
		if resp != nil {
			t.Fatalf("Confirm() response = %+v, want nil", resp)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
	})

	t.Run("returns already confirmed", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uow.reservationRepo.byReservationID["reservation-001"] = &entity.Reservation{
			ReservationID: "reservation-001",
			Status:        entity.ReservationStatusConfirmed,
			ExpiresAt:     time.Now().UTC().Add(time.Minute),
		}
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "reservation-001",
		})

		if !errors.Is(err, errordomain.ErrReservationAlreadyConfirmed) {
			t.Fatalf("Confirm() error = %v, want %v", err, errordomain.ErrReservationAlreadyConfirmed)
		}
		if resp != nil {
			t.Fatalf("Confirm() response = %+v, want nil", resp)
		}
	})

	t.Run("returns expired when reservation status is expired", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uow.reservationRepo.byReservationID["reservation-001"] = &entity.Reservation{
			ReservationID: "reservation-001",
			Status:        entity.ReservationStatusExpired,
			ExpiresAt:     time.Now().UTC().Add(time.Minute),
		}
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "reservation-001",
		})

		if !errors.Is(err, errordomain.ErrReservationExpired) {
			t.Fatalf("Confirm() error = %v, want %v", err, errordomain.ErrReservationExpired)
		}
		if resp != nil {
			t.Fatalf("Confirm() response = %+v, want nil", resp)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
		if uow.reservationRepo.updateCalls != 0 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 0)
		}
	})

	t.Run("returns expired when expiration time has passed", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uow.reservationRepo.byReservationID["reservation-001"] = &entity.Reservation{
			ReservationID: "reservation-001",
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(-time.Minute),
		}
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "reservation-001",
		})

		if !errors.Is(err, errordomain.ErrReservationExpired) {
			t.Fatalf("Confirm() error = %v, want %v", err, errordomain.ErrReservationExpired)
		}
		if resp != nil {
			t.Fatalf("Confirm() response = %+v, want nil", resp)
		}
	})

	t.Run("returns inventory not found when reservation inventory does not exist", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uow.reservationRepo.byReservationID["reservation-001"] = &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   99,
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(time.Minute),
		}
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "reservation-001",
		})

		if !errors.Is(err, errordomain.ErrInventoryNotFound) {
			t.Fatalf("Confirm() error = %v, want %v", err, errordomain.ErrInventoryNotFound)
		}
		if resp != nil {
			t.Fatalf("Confirm() response = %+v, want nil", resp)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
		if uow.reservationRepo.updateCalls != 0 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 0)
		}
	})

	t.Run("returns repository error when inventory update fails", func(t *testing.T) {
		ctx := context.Background()
		updateErr := errors.New("update inventory failed")
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    10,
			ReservedStock: 4,
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byID[inventory.ID] = inventory
		uow.inventoryRepo.updateErr = updateErr
		uow.reservationRepo.byReservationID["reservation-001"] = &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   inventory.ID,
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(time.Minute),
		}
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.Confirm(ctx, &request.ConfirmReservationRequest{
			ReservationID: "reservation-001",
		})

		if !errors.Is(err, updateErr) {
			t.Fatalf("Confirm() error = %v, want %v", err, updateErr)
		}
		if resp != nil {
			t.Fatalf("Confirm() response = %+v, want nil", resp)
		}
		if uow.inventoryRepo.updateCalls != 1 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 1)
		}
		if uow.reservationRepo.updateCalls != 0 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 0)
		}
	})
}

func TestInventoryUseCaseGetStock(t *testing.T) {
	t.Run("success returns stock", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byItemID["SKU-001"] = &entity.Inventory{
			ItemID:        "SKU-001",
			ItemName:      "Keyboard",
			TotalStock:    10,
			ReservedStock: 4,
		}
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.GetStock(ctx, "SKU-001")

		if err != nil {
			t.Fatalf("GetStock() error = %v", err)
		}
		if resp.ItemID != "SKU-001" ||
			resp.ItemName != "Keyboard" ||
			resp.TotalStock != 10 ||
			resp.ReservedStock != 4 ||
			resp.AvailableStock != 6 {
			t.Fatalf("GetStock() response = %+v", resp)
		}
	})

	t.Run("returns not found", func(t *testing.T) {
		ctx := context.Background()
		uow := newFakeUnitOfWork()
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		resp, err := uc.GetStock(ctx, "missing")

		if !errors.Is(err, errordomain.ErrInventoryNotFound) {
			t.Fatalf("GetStock() error = %v, want %v", err, errordomain.ErrInventoryNotFound)
		}
		if resp != nil {
			t.Fatalf("GetStock() response = %+v, want nil", resp)
		}
	})
}

func TestInventoryUseCaseListInventory(t *testing.T) {
	ctx := context.Background()
	uow := newFakeUnitOfWork()
	uow.inventoryRepo.all = []entity.Inventory{
		{
			ItemID:        "SKU-001",
			ItemName:      "Keyboard",
			TotalStock:    10,
			ReservedStock: 4,
		},
		{
			ItemID:        "SKU-002",
			ItemName:      "Mouse",
			TotalStock:    6,
			ReservedStock: 1,
		},
	}
	uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

	resp, err := uc.ListInventory(ctx)

	if err != nil {
		t.Fatalf("ListInventory() error = %v", err)
	}
	if len(resp.Items) != 2 {
		t.Fatalf("items length = %d, want %d", len(resp.Items), 2)
	}
	if resp.Items[0].AvailableStock != 6 {
		t.Fatalf("first available stock = %d, want %d", resp.Items[0].AvailableStock, 6)
	}
	if resp.Items[1].AvailableStock != 5 {
		t.Fatalf("second available stock = %d, want %d", resp.Items[1].AvailableStock, 5)
	}
}

func TestInventoryUseCaseExpireReservation(t *testing.T) {
	t.Run("success expires active reservations and releases stock", func(t *testing.T) {
		ctx := context.Background()
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    10,
			ReservedStock: 4,
		}
		reservation := &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   inventory.ID,
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(-time.Minute),
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byID[inventory.ID] = inventory
		uow.reservationRepo.expiredActive = []entity.Reservation{*reservation}
		uow.reservationRepo.byReservationID[reservation.ReservationID] = reservation
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		err := uc.ExpireReservation(ctx)

		if err != nil {
			t.Fatalf("ExpireReservation() error = %v", err)
		}
		if inventory.ReservedStock != 1 {
			t.Fatalf("ReservedStock = %d, want %d", inventory.ReservedStock, 1)
		}
		if reservation.Status != entity.ReservationStatusExpired {
			t.Fatalf("Status = %q, want %q", reservation.Status, entity.ReservationStatusExpired)
		}
		if uow.txCalls != 1 {
			t.Fatalf("transaction calls = %d, want %d", uow.txCalls, 1)
		}
		if uow.inventoryRepo.updateCalls != 1 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 1)
		}
		if uow.reservationRepo.updateCalls != 1 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 1)
		}
	})

	t.Run("clamps reserved stock to zero", func(t *testing.T) {
		ctx := context.Background()
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    10,
			ReservedStock: 2,
		}
		reservation := &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   inventory.ID,
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(-time.Minute),
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byID[inventory.ID] = inventory
		uow.reservationRepo.expiredActive = []entity.Reservation{*reservation}
		uow.reservationRepo.byReservationID[reservation.ReservationID] = reservation
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		err := uc.ExpireReservation(ctx)

		if err != nil {
			t.Fatalf("ExpireReservation() error = %v", err)
		}
		if inventory.ReservedStock != 0 {
			t.Fatalf("ReservedStock = %d, want %d", inventory.ReservedStock, 0)
		}
	})

	t.Run("skips when inventory does not exist", func(t *testing.T) {
		ctx := context.Background()
		reservation := &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   99,
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(-time.Minute),
		}
		uow := newFakeUnitOfWork()
		uow.reservationRepo.expiredActive = []entity.Reservation{*reservation}
		uow.reservationRepo.byReservationID[reservation.ReservationID] = reservation
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		err := uc.ExpireReservation(ctx)

		if err != nil {
			t.Fatalf("ExpireReservation() error = %v", err)
		}
		if reservation.Status != entity.ReservationStatusActive {
			t.Fatalf("Status = %q, want unchanged %q", reservation.Status, entity.ReservationStatusActive)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
		if uow.reservationRepo.updateCalls != 0 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 0)
		}
	})

	t.Run("skips when reservation is no longer active", func(t *testing.T) {
		ctx := context.Background()
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    10,
			ReservedStock: 4,
		}
		reservation := &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   inventory.ID,
			Quantity:      3,
			Status:        entity.ReservationStatusConfirmed,
			ExpiresAt:     time.Now().UTC().Add(-time.Minute),
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byID[inventory.ID] = inventory
		uow.reservationRepo.expiredActive = []entity.Reservation{*reservation}
		uow.reservationRepo.byReservationID[reservation.ReservationID] = reservation
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		err := uc.ExpireReservation(ctx)

		if err != nil {
			t.Fatalf("ExpireReservation() error = %v", err)
		}
		if inventory.ReservedStock != 4 {
			t.Fatalf("ReservedStock = %d, want unchanged %d", inventory.ReservedStock, 4)
		}
		if reservation.Status != entity.ReservationStatusConfirmed {
			t.Fatalf("Status = %q, want unchanged %q", reservation.Status, entity.ReservationStatusConfirmed)
		}
		if uow.inventoryRepo.updateCalls != 0 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 0)
		}
		if uow.reservationRepo.updateCalls != 0 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 0)
		}
	})

	t.Run("returns repository error when inventory update fails", func(t *testing.T) {
		ctx := context.Background()
		updateErr := errors.New("update inventory failed")
		inventory := &entity.Inventory{
			ID:            1,
			ItemID:        "SKU-001",
			TotalStock:    10,
			ReservedStock: 4,
		}
		reservation := &entity.Reservation{
			ReservationID: "reservation-001",
			InventoryID:   inventory.ID,
			Quantity:      3,
			Status:        entity.ReservationStatusActive,
			ExpiresAt:     time.Now().UTC().Add(-time.Minute),
		}
		uow := newFakeUnitOfWork()
		uow.inventoryRepo.byID[inventory.ID] = inventory
		uow.inventoryRepo.updateErr = updateErr
		uow.reservationRepo.expiredActive = []entity.Reservation{*reservation}
		uow.reservationRepo.byReservationID[reservation.ReservationID] = reservation
		uc := NewInventoryUseCase(uow, fakeUUIDGenerator{})

		err := uc.ExpireReservation(ctx)

		if !errors.Is(err, updateErr) {
			t.Fatalf("ExpireReservation() error = %v, want %v", err, updateErr)
		}
		if uow.inventoryRepo.updateCalls != 1 {
			t.Fatalf("inventory update calls = %d, want %d", uow.inventoryRepo.updateCalls, 1)
		}
		if uow.reservationRepo.updateCalls != 0 {
			t.Fatalf("reservation update calls = %d, want %d", uow.reservationRepo.updateCalls, 0)
		}
	})
}

type fakeUUIDGenerator struct {
	value string
}

func (g fakeUUIDGenerator) New() string {
	return g.value
}

type fakeUnitOfWork struct {
	inventoryRepo   *fakeInventoryRepository
	reservationRepo *fakeReservationRepository
	txCalls         int
}

func newFakeUnitOfWork() *fakeUnitOfWork {
	return &fakeUnitOfWork{
		inventoryRepo: &fakeInventoryRepository{
			byID:     map[int64]*entity.Inventory{},
			byItemID: map[string]*entity.Inventory{},
		},
		reservationRepo: &fakeReservationRepository{
			byReservationID: map[string]*entity.Reservation{},
		},
	}
}

func (u *fakeUnitOfWork) WithTransaction(
	ctx context.Context,
	fn func(repositoryinterface.UnitOfWork) error,
) error {
	u.txCalls++
	return fn(u)
}

func (u *fakeUnitOfWork) InventoryRepository() repositoryinterface.InventoryRepository {
	return u.inventoryRepo
}

func (u *fakeUnitOfWork) ReservationRepository() repositoryinterface.ReservationRepository {
	return u.reservationRepo
}

type fakeInventoryRepository struct {
	all         []entity.Inventory
	byID        map[int64]*entity.Inventory
	byItemID    map[string]*entity.Inventory
	updateCalls int
	updateErr   error
}

func (r *fakeInventoryRepository) FindAll(ctx context.Context) ([]entity.Inventory, error) {
	return r.all, nil
}

func (r *fakeInventoryRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entity.Inventory, error) {
	return r.byID[id], nil
}

func (r *fakeInventoryRepository) FindByIDForUpdate(
	ctx context.Context,
	id int64,
) (*entity.Inventory, error) {
	return r.byID[id], nil
}

func (r *fakeInventoryRepository) FindByItemID(
	ctx context.Context,
	itemID string,
) (*entity.Inventory, error) {
	return r.byItemID[itemID], nil
}

func (r *fakeInventoryRepository) FindByItemIDForUpdate(
	ctx context.Context,
	itemID string,
) (*entity.Inventory, error) {
	return r.byItemID[itemID], nil
}

func (r *fakeInventoryRepository) Update(
	ctx context.Context,
	inventory *entity.Inventory,
) error {
	r.updateCalls++
	return r.updateErr
}

type fakeReservationRepository struct {
	byReservationID map[string]*entity.Reservation
	expiredActive   []entity.Reservation
	created         []*entity.Reservation
	createCalls     int
	updateCalls     int
	createErr       error
	updateErr       error
}

func (r *fakeReservationRepository) Create(
	ctx context.Context,
	reservation *entity.Reservation,
) error {
	r.createCalls++
	r.created = append(r.created, reservation)
	return r.createErr
}

func (r *fakeReservationRepository) FindByReservationID(
	ctx context.Context,
	reservationID string,
) (*entity.Reservation, error) {
	return r.byReservationID[reservationID], nil
}

func (r *fakeReservationRepository) FindByReservationIDForUpdate(
	ctx context.Context,
	reservationID string,
) (*entity.Reservation, error) {
	return r.byReservationID[reservationID], nil
}

func (r *fakeReservationRepository) FindExpiredActive(
	ctx context.Context,
	limit int,
) ([]entity.Reservation, error) {
	return r.expiredActive, nil
}

func (r *fakeReservationRepository) Update(
	ctx context.Context,
	reservation *entity.Reservation,
) error {
	r.updateCalls++
	return r.updateErr
}
