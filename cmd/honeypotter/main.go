package main

import (
	"github.com/Semerokozlyat/honeypotter/internal"
	"github.com/Semerokozlyat/honeypotter/internal/config"
)

func main() {
	cfg := loadConfig()
	app, err := internal.NewApp(cfg)
	if err != nil {
		panic("create app instance: " + err.Error())
	}
	defer func() {
		app.Stop()
	}()

	if err = app.Start(); err != nil {
		panic("start app: " + err.Error())
	}
}

func loadConfig() (cfg *config.Config) {
	cfg = &config.Config{
		HTTPServer: config.HTTPServer{
			Address: "127.0.0.1:9090",
		},
		Database: config.Database{
			URL: "",
		},
	}
	return cfg
}
