package internal

import "github.com/Semerokozlyat/honeypotter/internal/server"

type App struct {
	httpServer *server.HTTPServer
}

func NewApp() *App {
	httpSrv := server.NewHTTPServer()

	return &App{
		httpServer: httpSrv,
	}
}

func (app *App) Start() error {
	return app.httpServer.Run()
}
