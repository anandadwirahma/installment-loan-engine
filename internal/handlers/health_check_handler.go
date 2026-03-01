package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/shared/constant"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    constant.SuccessCode,
		Message: "Billing Service is running.",
	})
}
