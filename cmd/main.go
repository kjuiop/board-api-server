package main

import (
	"board-api-server/cmd/app"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	BUILD_TIME  = "no flag of BUILD_TIME"
	GIT_HASH    = "no flag of GIT_HASH"
	APP_VERSION = "no flag of APP_VERSION"
)

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	a := app.NewApplication(ctx)
	go a.Start(&wg)

	slog.Debug("boarder api server app start", "git_hash", GIT_HASH, "build_time", BUILD_TIME, "app_version", APP_VERSION)

	<-exitSignal()
	a.Stop(ctx)
	cancel()
	slog.Debug("board api server app gracefully stopped")
}

func exitSignal() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
