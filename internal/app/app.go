package app

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/qreaqtor/referrals/internal/config"
	appserver "github.com/qreaqtor/referrals/pkg/appServer"
	httpserver "github.com/qreaqtor/referrals/pkg/httpServer"
)

type appServer interface {
	Start() error
	Wait() []error
}

type App struct {
	server appServer
}

func NewApp(ctx context.Context, cfg *config.Config) *App {
	setupLogger(cfg.Env)

	r := mux.NewRouter()

	return &App{
		server: appserver.NewAppServer(
			ctx,
			httpserver.NewHTTPServer(r),
			cfg.Addr.ToString(),
		),
	}
}

func (a *App) Start() error {
	return a.server.Start()
}

func (a *App) Wait() []error {
	return a.server.Wait()
}
