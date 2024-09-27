package app

import "context"

type App struct {
}

func NewApplication(ctx context.Context) *App {
	return &App{}
}

func (a *App) Start() {
}

func (a *App) Stop(ctx context.Context) {
}
