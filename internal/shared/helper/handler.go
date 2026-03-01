package helper

import (
	"github.com/gin-gonic/gin"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/shared/errors"
)

func ErrorHandler(c *gin.Context, err error) {
	e := errors.ErrGeneral

	if errors.IsCustomError(err) {
		e = errors.Unwrap(err)
	}

	c.JSON(e.HttpCode, dto.APIResponse{
		Code:    e.Code,
		Message: e.Message,
	})
}
