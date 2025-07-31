package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/Semerokozlyat/honeypotter/internal/handler/httphandler"
)

type HTTPServer struct {
	router *gin.Engine

	addr string
}

func NewHTTPServer() *HTTPServer {
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	apiV1 := ginEngine.Group("/v1")
	apiV1.GET("/status", httphandler.StatusHandler)

	return &HTTPServer{
		router: ginEngine,
	}
}

func (srv *HTTPServer) Run() error {
	return http.ListenAndServe(srv.addr, srv.router)
}
