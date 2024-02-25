package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mjmcconnell/feature-manager/internal"
	// modules
	_ "github.com/mjmcconnell/feature-manager/modules/hello"
)

func main() {
	// Handle SIGINT gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	internal.NewObservability(ctx)

	router := internal.NewRouter()
	router.StartServer()

	internal.InitModules(router)
}
