package app

import (
	"board-api-server/api/route"
	"board-api-server/config"
	"board-api-server/internal/logger"
	"board-api-server/internal/server"
	"context"
	"log"
	"sync"
)

type App struct {
	cfg *config.EnvConfig
	srv *server.Gin
}

func NewApplication(ctx context.Context) *App {

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config, err : %v", err)
	}

	if err := logger.SlogInit(cfg.Logger); err != nil {
		log.Fatalf("fail to init slog err : %v", err)
	}

	srv := server.NewGinServer(cfg.Server)

	app := &App{
		cfg: cfg,
		srv: srv,
	}

	app.setupRouter(ctx)

	return app
}

func (a *App) Start(wg *sync.WaitGroup) {
	a.srv.Run(wg)
}

func (a *App) Stop(ctx context.Context) {
	a.srv.Shutdown(ctx)
}

func (a *App) setupRouter(ctx context.Context) {
	router := route.Config{
		Engine: a.srv.GetEngine(),
	}
	router.Setup()
}
