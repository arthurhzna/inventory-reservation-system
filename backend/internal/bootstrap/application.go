package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/config"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/database"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/identity"
)

type Application struct {
	HttpServer *HttpServer
}

func NewApplication() *Application {

	cfg := config.InitConfig()

	log := NewLogger(cfg)

	db := NewDatabase(
		cfg,
		log,
	)

	uow := database.NewUnitOfWork(
		db,
	)
	uuidGenerator := identity.NewUUIDGenerator()

	useCase := NewUseCase(
		uow,
		uuidGenerator,
	)

	controller := NewController(
		useCase,
	)

	httpServer := NewHTTPServer(
		cfg,
		log,
		controller,
	)

	return &Application{
		HttpServer: httpServer,
	}
}
