package main

import (
	"bard-api-server/cmd/app"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	a := app.NewApplication(ctx)
	a.Start()

	<-exitSignal()
	a.Stop(ctx)
	cancel()
}

func exitSignal() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
