package app

import (
	"bard-api-server/config"
	"bard-api-server/internal/logger"
	"context"
	"log"
)

type App struct {
}

func NewApplication(ctx context.Context) *App {

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config, err : %v", err)
	}

	if err := logger.SlogInit(cfg.Logger); err != nil {
		log.Fatalf("fail to init slog err : %v", err)
	}

	return &App{}
}

func (a *App) Start() {
}

func (a *App) Stop(ctx context.Context) {
}
