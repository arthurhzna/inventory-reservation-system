package worker

import (
	"context"
	"time"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/usecase"
)

type ReservationExpiryWorker interface {
	Run(ctx context.Context)
}

type reservationExpiryWorker struct {
	inventoryUseCase usecase.InventoryUseCaseInterface
}

func NewReservationExpiryWorker(
	inventoryUseCase usecase.InventoryUseCaseInterface,
) ReservationExpiryWorker {
	return &reservationExpiryWorker{
		inventoryUseCase: inventoryUseCase,
	}
}

func (w *reservationExpiryWorker) Run(
	ctx context.Context,
) {

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {

		case <-ctx.Done():
			return

		case <-ticker.C:

			if err := w.inventoryUseCase.
				ExpireReservation(
					ctx,
				); err != nil {
			}
		}
	}
}
