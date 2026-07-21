package worker

import (
	"context"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/bootstrap"
)

func runHttpWorker(
	ctx context.Context,
) {

	app := bootstrap.NewApplication()

	go app.HttpServer.Start()

	go app.Worker.
		ReservationExpiryWorker.
		Run(ctx)

	<-ctx.Done()

	app.HttpServer.Shutdown()
}
