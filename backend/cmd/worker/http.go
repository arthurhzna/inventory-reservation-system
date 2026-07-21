package worker

import (
	"context"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/bootstrap"
)

func runHttpWorker(ctx context.Context) {
	srv := bootstrap.NewApplication()
	go srv.HttpServer.Start()

	<-ctx.Done()
	srv.HttpServer.Shutdown()
}
