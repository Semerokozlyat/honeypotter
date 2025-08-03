package internal

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Semerokozlyat/honeypotter/internal/config"
	"github.com/Semerokozlyat/honeypotter/internal/repository"
	"github.com/Semerokozlyat/honeypotter/internal/server"
)

type App struct {
	httpServer      *server.HTTPServer
	httpRequestRepo *repository.HTTPRequestRepository
	dbPool          *pgxpool.Pool
}

func NewApp(cfg *config.Config) (*App, error) {

	dbPool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		return nil, fmt.Errorf("create db connections pool: %w", err)
	}
	httpRequestRepo := repository.NewHTTPRequestRepository(dbPool)

	httpSrv := server.NewHTTPServer(&cfg.HTTPServer, httpRequestRepo)

	return &App{
		httpServer:      httpSrv,
		httpRequestRepo: httpRequestRepo,
	}, nil
}

func (app *App) Start() error {
	return app.httpServer.Run()
}

func (app *App) Stop() {
	app.dbPool.Close()
}
