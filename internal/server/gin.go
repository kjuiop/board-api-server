package server

import (
	"board-api-server/api/middleware"
	"board-api-server/config"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"sync"
)

type Gin struct {
	cfg    config.Server
	srv    *http.Server
	router *gin.Engine
}

func NewGinServer(cfg config.Server) *Gin {

	router := getGinEngine(cfg.Mode)
	router.Use(middleware.RecoveryErrorReport())

	if err := router.SetTrustedProxies(strings.Split(cfg.TrustedProxies, ",")); err != nil {
		log.Fatalf("failed set trust proxies, err : %v", err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	return &Gin{
		srv:    srv,
		cfg:    cfg,
		router: router,
	}
}

func (g *Gin) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	err := g.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		slog.Debug("server close")
	} else {
		slog.Error("run server error", "error", err)
	}
}

func (g *Gin) Shutdown(ctx context.Context) {
	if err := g.srv.Shutdown(ctx); err != nil {
		slog.Error("error during server shutdown", "error", err)
	}
}

func getGinEngine(mode string) *gin.Engine {
	switch mode {
	case "prod":
		return gin.New()
	case "test":
		gin.SetMode(gin.TestMode)
		return gin.Default()
	default:
		return gin.Default()
	}
}

func (g *Gin) GetEngine() *gin.Engine {
	return g.router
}
