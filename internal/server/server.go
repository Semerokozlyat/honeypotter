package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Semerokozlyat/honeypotter/internal/config"
	"github.com/Semerokozlyat/honeypotter/internal/handler/httphandler"
)

type HTTPServer struct {
	httpServer *http.Server

	httpRequestRepo HTTPRequestsRepo
}

type HTTPRequestsRepo interface {
	CreateHTTPRequest(ctx context.Context) error
}

func NewHTTPServer(cfg *config.HTTPServer, httpRequestRepo HTTPRequestsRepo) *HTTPServer {
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	apiV1 := ginEngine.Group("/v1")
	apiV1.GET("/status", httphandler.StatusHandler)
	apiV1.GET("/infrastructure", httphandler.InfrastructureHandler)

	return &HTTPServer{
		httpServer: &http.Server{
			Addr:           cfg.Address,
			Handler:        ginEngine,
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 10 * 1024 * 1024 * 1024,
		},
		httpRequestRepo: httpRequestRepo,
	}
}

func (srv *HTTPServer) Run() error {
	log.Println("running HTTP server")
	return srv.httpServer.ListenAndServe()
}
