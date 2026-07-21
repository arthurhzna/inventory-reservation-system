package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/controller"
)

type Controller struct {
	AppController       *controller.AppController
	InventoryController *controller.InventoryController
}

func NewController(
	useCase *UseCase,
) *Controller {
	return &Controller{
		AppController: controller.NewAppController(),

		InventoryController: controller.NewInventoryController(
			useCase.InventoryUseCase,
		),
	}
}
