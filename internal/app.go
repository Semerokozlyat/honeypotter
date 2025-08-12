package internal

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"

	"github.com/Semerokozlyat/honeypotter/internal/config"
	"github.com/Semerokozlyat/honeypotter/internal/repository"
	"github.com/Semerokozlyat/honeypotter/internal/server"
)

type App struct {
	httpServer      *server.HTTPServer
	httpRequestRepo *repository.HTTPRequestRepository
	packetCapturer  *server.PacketCapturer
	dbPool          *pgxpool.Pool
}

func NewApp(cfg *config.Config) (*App, error) {

	dbPool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		return nil, fmt.Errorf("create db connections pool: %w", err)
	}
	httpRequestRepo := repository.NewHTTPRequestRepository(dbPool)

	httpSrv := server.NewHTTPServer(&cfg.HTTPServer, httpRequestRepo)

	packetCapturer, err := server.NewPacketCapturer(&cfg.PacketCapturer)
	if err != nil {
		return nil, fmt.Errorf("init packet capturer: %w", err)
	}

	return &App{
		httpServer:      httpSrv,
		httpRequestRepo: httpRequestRepo,
		packetCapturer:  packetCapturer,
	}, nil
}

func (app *App) Start(wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := app.httpServer.Run()
		if err != nil {
			log.Fatal("failed to start HTTP server", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := app.packetCapturer.Run()
		if err != nil {
			log.Fatal("failed to start packet capturer", err)
		}
	}()
	return nil
}

func (app *App) Stop() {
	app.dbPool.Close()
}
