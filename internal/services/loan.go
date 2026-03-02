package services

import (
	"context"
	"time"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/constant"
	"installment-loan-engine/internal/shared/errors"
	"installment-loan-engine/internal/shared/helper"
	"installment-loan-engine/internal/shared/logger"
)

func (s *loanService) CreateLoan(ctx context.Context, req dto.CreateLoanRequest) (dto.CreateLoanResponse, error) {
	totalInterest := int64(float64(req.PrincipalAmount) * s.cfg.InterestRate)
	totalRepayment := req.PrincipalAmount + totalInterest

	loanEntity := &entity.Loan{
		LoanRefNum:           helper.GenerateUniqueNumber("LN"),
		BorrowerRefNum:       req.BorrowerRefNum,
		PrincipalAmount:      req.PrincipalAmount,
		InterestRate:         s.cfg.InterestRate,
		InterestType:         s.cfg.InterestType,
		TenorWeeks:           s.cfg.TenorWeeks,
		Status:               constant.LoanStatusActive,
		TotalInterestAmount:  totalInterest,
		TotalRepaymentAmount: totalRepayment,
		OutstandingAmount:    totalRepayment,
	}

	tx := s.loanRepo.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			s.loanRepo.RollbackTx(tx)
		}
	}()

	if err := s.loanRepo.CreateWithTx(tx, loanEntity); err != nil {
		s.loanRepo.RollbackTx(tx)
		logger.Errorf("[service.CreateLoan] Error creating loan: %v", err)
		return dto.CreateLoanResponse{}, errors.ErrGeneral
	}

	installmentEntities := s.generateInstallments(loanEntity)
	if err := s.installmentRepo.CreateWithTx(tx, installmentEntities); err != nil {
		s.loanRepo.RollbackTx(tx)
		logger.Errorf("[service.CreateLoan] Error creating installments: %v", err)
		return dto.CreateLoanResponse{}, errors.ErrGeneral
	}

	if err := s.loanRepo.CommitTx(tx); err != nil {
		logger.Errorf("[service.CreateLoan] Error committing transaction: %v", err)
		return dto.CreateLoanResponse{}, errors.ErrGeneral
	}

	return dto.CreateLoanResponse{
		LoanRefNum:           loanEntity.LoanRefNum,
		PrincipalAmount:      loanEntity.PrincipalAmount,
		TotalInterestAmount:  loanEntity.TotalInterestAmount,
		TotalRepaymentAmount: loanEntity.TotalRepaymentAmount,
		WeeklyInstallment:    loanEntity.TotalRepaymentAmount / int64(loanEntity.TenorWeeks),
		Status:               string(loanEntity.Status),
	}, nil
}

func (s *loanService) generateInstallments(loanEntity *entity.Loan) []*entity.Installment {
	installments := make([]*entity.Installment, 0, loanEntity.TenorWeeks)

	weeklyPrincipal := loanEntity.PrincipalAmount / int64(loanEntity.TenorWeeks)
	weeklyInterest := loanEntity.TotalInterestAmount / int64(loanEntity.TenorWeeks)

	now := time.Now()

	for i := 1; i <= int(loanEntity.TenorWeeks); i++ {
		inst := &entity.Installment{
			LoanID:            loanEntity.ID,
			InstallmentNumber: int16(i),
			DueDate:           now.AddDate(0, 0, 7*i),
			PrincipalAmount:   weeklyPrincipal,
			InterestAmount:    weeklyInterest,
			TotalAmount:       weeklyPrincipal + weeklyInterest,
			Status:            constant.InstallmentStatusPending,
		}

		installments = append(installments, inst)
	}

	return installments
}
