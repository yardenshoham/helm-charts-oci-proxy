package main

import (
	"context"
	"github.com/container-registry/ocip/cmd"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := cmd.Root.ExecuteContext(ctx); err != nil {
		cancel()
		os.Exit(1)
	}
}
