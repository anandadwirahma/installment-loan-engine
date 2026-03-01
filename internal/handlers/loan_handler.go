package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/services"
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
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid request payload",
		})
		return
	}

	loan, err := h.service.CreateLoan(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to create loan",
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    "SUCCESS",
		Message: "Loan created successfully.",
		Data:    loan,
	})
}

func (h *LoanHandler) GetInstallment(c *gin.Context) {
	var req dto.GetInstallmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid request payload",
		})
		return
	}

	installment, err := h.service.GetInstallment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to get installment",
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    "SUCCESS",
		Message: "Installment retrieved successfully.",
		Data:    installment,
	})
}

func (h *LoanHandler) GetOutstanding(c *gin.Context) {
	req := dto.GetOutstandingRequest{LoanRefNum: c.Param("loan_ref_num")}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid request payload",
		})
		return
	}

	outstanding, err := h.service.GetOutstanding(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to get outstanding",
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    "SUCCESS",
		Message: "Outstanding retrieved successfully.",
		Data:    outstanding,
	})
}

func (h *LoanHandler) CheckDelinquent(c *gin.Context) {
	req := dto.CheckDelinquentRequest{LoanRefNum: c.Param("loan_ref_num")}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid request payload",
		})
		return
	}

	delinquent, err := h.service.CheckDelinquent(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to check delinquent",
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    "SUCCESS",
		Message: "Delinquent checked successfully.",
		Data:    delinquent,
	})
}

func (h *LoanHandler) PayInstallment(c *gin.Context) {
	var req dto.PayInstallmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid request payload",
		})
		return
	}

	payment, err := h.service.PayInstallment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Failed to pay installment",
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Code:    "SUCCESS",
		Message: "Installment paid successfully.",
		Data:    payment,
	})
}
