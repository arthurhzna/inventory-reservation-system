package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/usecase"

	repositorydomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"
	servicedomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/service"
)

type UseCase struct {
	InventoryUseCase usecase.InventoryUseCaseInterface
}

func NewUseCase(
	uow repositorydomain.UnitOfWork,
	uuidGenerator servicedomain.UUIDGenerator,
) *UseCase {

	return &UseCase{
		InventoryUseCase: usecase.NewInventoryUseCase(
			uow,
			uuidGenerator,
		),
	}
}
