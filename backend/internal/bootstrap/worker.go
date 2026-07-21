package bootstrap

import (
	workerapplication "github.com/arthurhzna/inventory-reservation-system/backend/internal/application/worker"
)

type Worker struct {
	ReservationExpiryWorker workerapplication.ReservationExpiryWorker
}

func NewWorker(
	useCase *UseCase,
) *Worker {

	return &Worker{
		ReservationExpiryWorker: workerapplication.NewReservationExpiryWorker(
			useCase.InventoryUseCase,
		),
	}
}
