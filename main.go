package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"installment-loan-engine/internal/handlers"
	"installment-loan-engine/internal/repositories"
	"installment-loan-engine/internal/services"
	"installment-loan-engine/internal/shared/config"
	"installment-loan-engine/internal/shared/database"
	"installment-loan-engine/internal/shared/logger"
)

func main() {
	// Initiate Logger
	logger.Init()

	// Initiate Config
	config.Init()

	// Initiate Database
	database.Init()

	// Initiate Repository
	loanRepo := repositories.NewLoanRepository(database.DB)
	installmentRepo := repositories.NewInstallmentRepository(database.DB)
	transactionRepo := repositories.NewTransactionRepository(database.DB)

	// Initiate Service
	loanService := services.NewLoanService(loanRepo, installmentRepo, transactionRepo, config.AppConfig)

	// Initiate Handler
	healthCheckHandler := handlers.NewHealthCheckHandler()
	loanHandler := handlers.NewLoanHandler(loanService)

	router := gin.Default()

	// Initiate Routes
	router.GET("/health", healthCheckHandler.HealthCheck)
	router.POST("/api/v1/loans", loanHandler.CreateLoan)
	router.GET("/api/v1/loans/:loan_ref_num/installments", loanHandler.GetInstallment)
	router.GET("/api/v1/loans/:loan_ref_num/outstanding", loanHandler.GetOutstanding)
	router.GET("/api/v1/loans/:loan_ref_num/delinquent", loanHandler.CheckDelinquent)
	router.POST("/api/v1/loans/payment", loanHandler.PayInstallment)

	router.Run(fmt.Sprintf(":%s", config.AppConfig.AppPort))
}
