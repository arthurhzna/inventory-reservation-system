package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/config"
	loggerdomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/logger"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/logging"
)

func NewLogger(
	cfg *config.Config,
) loggerdomain.Logger {

	return logging.NewZeroLogger(
		cfg.Logger.Level,
	)
}
