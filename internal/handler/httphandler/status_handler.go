package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusHandler(ginCtx *gin.Context) {
	ginCtx.Status(http.StatusOK)
}
