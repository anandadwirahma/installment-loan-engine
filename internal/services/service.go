package services

import (
	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/repositories"
	"installment-loan-engine/internal/shared/config"
)

type LoanService interface {
	CreateLoan(req dto.CreateLoanRequest) (dto.CreateLoanResponse, error)
	GetInstallment(req dto.GetInstallmentRequest) (dto.GetInstallmentResponse, error)
	GetOutstanding(req dto.GetOutstandingRequest) (dto.GetOutstandingResponse, error)
	CheckDelinquent(req dto.CheckDelinquentRequest) (dto.CheckDelinquentResponse, error)
	PayInstallment(req dto.PayInstallmentRequest) (dto.PayInstallmentResponse, error)
}

type loanService struct {
	loanRepo        repositories.LoanRepository
	installmentRepo repositories.InstallmentRepository
	transactionRepo repositories.TransactionRepository
	cfg             config.Config
}

func NewLoanService(
	loanRepo repositories.LoanRepository,
	installmentRepo repositories.InstallmentRepository,
	transactionRepo repositories.TransactionRepository,
	cfg config.Config,
) LoanService {
	return &loanService{
		loanRepo:        loanRepo,
		installmentRepo: installmentRepo,
		transactionRepo: transactionRepo,
		cfg:             cfg,
	}
}
