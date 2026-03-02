package services

import (
	"context"
	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/repositories"
	"installment-loan-engine/internal/shared/config"
)

type LoanService interface {
	CreateLoan(ctx context.Context, req dto.CreateLoanRequest) (dto.CreateLoanResponse, error)
	GetInstallment(ctx context.Context, req dto.GetInstallmentRequest) (dto.GetInstallmentResponse, error)
	GetOutstanding(ctx context.Context, req dto.GetOutstandingRequest) (dto.GetOutstandingResponse, error)
	CheckDelinquent(ctx context.Context, req dto.CheckDelinquentRequest) (dto.CheckDelinquentResponse, error)
	PayInstallment(ctx context.Context, req dto.PayInstallmentRequest) (dto.PayInstallmentResponse, error)
}

type loanService struct {
	loanRepo        repositories.LoanRepository
	installmentRepo repositories.InstallmentRepository
	transactionRepo repositories.TransactionRepository
	cacheRepo       repositories.CacheRepository
	cfg             config.Config
}

func NewLoanService(
	loanRepo repositories.LoanRepository,
	installmentRepo repositories.InstallmentRepository,
	transactionRepo repositories.TransactionRepository,
	cacheRepo repositories.CacheRepository,
	cfg config.Config,
) LoanService {
	return &loanService{
		loanRepo:        loanRepo,
		installmentRepo: installmentRepo,
		transactionRepo: transactionRepo,
		cacheRepo:       cacheRepo,
		cfg:             cfg,
	}
}
