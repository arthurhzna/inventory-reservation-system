package worker

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func Start() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	rootCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run HTTP server",
		Run: func(cmd *cobra.Command, _ []string) {
			runHttpWorker(ctx)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
