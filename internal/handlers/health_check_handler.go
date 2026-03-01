package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"installment-loan-engine/internal/dto"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    "SUCCESS",
		Message: "Billing Service is running.",
	})
}
