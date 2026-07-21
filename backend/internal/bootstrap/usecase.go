package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/usecase"

	servicedomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/service"

	repositorydomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"
)

type UseCase struct {
	UserUseCase usecase.UserUseCaseInterface
}

func NewUseCase(
	uow repositorydomain.UnitOfWork,
	uuidGenerator servicedomain.UUIDGenerator,
) *UseCase {

	return &UseCase{
		UserUseCase: usecase.NewUserUseCase(
			uow,
			uuidGenerator,
		),
	}
}
