package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/services"
	"installment-loan-engine/internal/shared/constant"
	"installment-loan-engine/internal/shared/errors"
	"installment-loan-engine/internal/shared/helper"
	"installment-loan-engine/internal/shared/logger"
)

type LoanHandler struct {
	service services.LoanService
}

func NewLoanHandler(service services.LoanService) *LoanHandler {
	return &LoanHandler{
		service: service,
	}
}

func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var req dto.CreateLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorf("[handler.CreateLoan] Invalid request payload: %v", err)
		e := errors.ErrBadRequest
		c.JSON(e.HttpCode, dto.APIResponse{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	loan, err := h.service.CreateLoan(c.Request.Context(), req)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage,
		Data:    loan,
	})
}

func (h *LoanHandler) GetInstallment(c *gin.Context) {
	req := dto.GetInstallmentRequest{LoanRefNum: c.Param("loan_ref_num")}
	if req.LoanRefNum == "" {
		logger.Errorf("[handler.GetInstallment] LoanRefNum is required")
		e := errors.ErrBadRequest
		c.JSON(e.HttpCode, dto.APIResponse{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	installment, err := h.service.GetInstallment(c.Request.Context(), req)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage,
		Data:    installment,
	})
}

func (h *LoanHandler) GetOutstanding(c *gin.Context) {
	req := dto.GetOutstandingRequest{LoanRefNum: c.Param("loan_ref_num")}
	if req.LoanRefNum == "" {
		logger.Errorf("[handler.GetOutstanding] LoanRefNum is required")
		e := errors.ErrBadRequest
		c.JSON(e.HttpCode, dto.APIResponse{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	outstanding, err := h.service.GetOutstanding(c.Request.Context(), req)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage,
		Data:    outstanding,
	})
}

func (h *LoanHandler) CheckDelinquent(c *gin.Context) {
	req := dto.CheckDelinquentRequest{LoanRefNum: c.Param("loan_ref_num")}
	if req.LoanRefNum == "" {
		logger.Errorf("[handler.CheckDelinquent] LoanRefNum is required")
		e := errors.ErrBadRequest
		c.JSON(e.HttpCode, dto.APIResponse{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	delinquent, err := h.service.CheckDelinquent(c.Request.Context(), req)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage,
		Data:    delinquent,
	})
}

func (h *LoanHandler) PayInstallment(c *gin.Context) {
	var req dto.PayInstallmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorf("[handler.PayInstallment] Invalid request payload: %v", err)
		e := errors.ErrBadRequest
		c.JSON(e.HttpCode, dto.APIResponse{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	payment, err := h.service.PayInstallment(c.Request.Context(), req)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage,
		Data:    payment,
	})
}
