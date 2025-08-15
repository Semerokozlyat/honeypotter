package httphandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfrastructureResp struct {
	Status string `json:"status" yaml:"status"`
}

func InfrastructureHandler(ginCtx *gin.Context) {
	log.Printf("request to %s %s from IP: %s, headers: %+v \n",
		ginCtx.Request.Method, ginCtx.Request.URL.EscapedPath(), ginCtx.ClientIP(), ginCtx.Request.Header)
	ginCtx.JSON(http.StatusOK, InfrastructureResp{Status: "alive"})
}
