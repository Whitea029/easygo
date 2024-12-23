package handlers

import (
	"1232313/api/response"
	"1232313/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logger = logging.GetLogger()

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HandlerGet(c *gin.Context) {
	logger.Info(logging.General, logging.Api, "Health check", nil)
	c.JSON(http.StatusOK, response.GenResp(true, response.Success, "ok", nil))
}
