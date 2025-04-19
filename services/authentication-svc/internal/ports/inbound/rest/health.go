package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nanicienta/api/pkg/ports/logging"
	"net/http"
)

type HealthRestHandler struct {
	gin *gin.Engine
}

// NewHealthRestHandler creates a new HealthRestHandler
func NewHealthRestHandler(gin *gin.Engine, logger logging.Logger) HealthRestHandler {
	return HealthRestHandler{
		gin: gin,
	}
}

// InitRouter initializes the router for health-related endpoints
func (h *HealthRestHandler) InitRouter() {
	api := h.gin.Group("/api/v1")
	{
		api.GET("/health", h.healthCheck)
	}
}

func (h *HealthRestHandler) healthCheck(c *gin.Context) {
	c.JSON(
		http.StatusOK, map[string]interface{}{
			"status": "ok",
		},
	)
}
