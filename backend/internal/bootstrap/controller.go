package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/controller"
)

type Controller struct {
	AppController  *controller.AppController
	UserController *controller.UserController
}

func NewController(
	useCase *UseCase,
) *Controller {
	return &Controller{
		AppController: controller.NewAppController(),

		UserController: controller.NewUserController(
			useCase.UserUseCase,
		),
	}
}
