package bootstrap

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/config"
	loggerdomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/logger"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/database"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/dbtx"
)

func NewDatabase(
	cfg *config.Config,
	log loggerdomain.Logger,
) dbtx.DBTX {

	db := database.NewDatabase(
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Sslmode,
		cfg.Database.MaxIdleConn,
		cfg.Database.MaxOpenConn,
		cfg.Database.MaxConnLifetime,
	)

	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return pool
}
