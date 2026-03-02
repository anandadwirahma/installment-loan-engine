package services

import (
	"context"
	"time"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/shared/constant"
	"installment-loan-engine/internal/shared/errors"
	"installment-loan-engine/internal/shared/logger"
)

func (s *loanService) GetInstallment(ctx context.Context, req dto.GetInstallmentRequest) (dto.GetInstallmentResponse, error) {
	loan, err := s.loanRepo.GetLoanInstallmentByRefNum(req.LoanRefNum)
	if err != nil {
		logger.Errorf("[service.GetInstallment] Error fetching loan for RefNum %s: %v", req.LoanRefNum, err)
		return dto.GetInstallmentResponse{}, errors.ErrGeneral
	}

	var (
		totalOverdue int
		installments = []dto.Installment{}

		now = time.Now()
	)

	for _, v := range loan.Installments {
		var paidAt *string

		if v.PaidAt != nil {
			paidAtStr := v.PaidAt.Format(time.RFC3339)
			paidAt = &paidAtStr
		}

		status := v.Status
		if v.Status == constant.InstallmentStatusPending {
			if v.DueDate.Before(now) {
				status = constant.InstallmentStatusOverdue
				totalOverdue++
			}
		}

		installments = append(installments, dto.Installment{
			InstallmentNumber: int(v.InstallmentNumber),
			PrincipalAmount:   v.PrincipalAmount,
			InterestAmount:    v.InterestAmount,
			TotalAmount:       v.TotalAmount,
			DueDate:           v.DueDate.Format(time.RFC3339),
			Status:            string(status),
			PaidAt:            paidAt,
		})
	}

	resp := dto.GetInstallmentResponse{
		LoanRefNum:        loan.LoanRefNum,
		TotalInstallments: int64(len(installments)),
		Installments:      installments,
		Summary: dto.Summary{
			TotalOutstanding: loan.OutstandingAmount,
			IsDelinquent:     totalOverdue >= 2,
		},
	}

	return resp, nil
}

func (s *loanService) GetOutstanding(ctx context.Context, req dto.GetOutstandingRequest) (dto.GetOutstandingResponse, error) {
	loan, err := s.loanRepo.GetByRefNum(req.LoanRefNum)
	if err != nil {
		logger.Errorf("[service.GetOutstanding] Error fetching loan for RefNum %s: %v", req.LoanRefNum, err)
		return dto.GetOutstandingResponse{}, errors.ErrNotFound
	}

	resp := dto.GetOutstandingResponse{
		LoanRefNum:        loan.LoanRefNum,
		OutstandingAmount: loan.OutstandingAmount,
	}

	return resp, nil
}

func (s *loanService) CheckDelinquent(ctx context.Context, req dto.CheckDelinquentRequest) (dto.CheckDelinquentResponse, error) {
	loan, err := s.loanRepo.GetByRefNum(req.LoanRefNum)
	if err != nil {
		logger.Errorf("[service.CheckDelinquent] Error fetching loan for RefNum %s: %v", req.LoanRefNum, err)
		return dto.CheckDelinquentResponse{}, errors.ErrNotFound
	}

	installments, err := s.installmentRepo.GetOverdueInstallment(loan.ID)
	if err != nil {
		logger.Errorf("[service.CheckDelinquent] Error fetching overdue installment for RefNum %s: %v", req.LoanRefNum, err)
		return dto.CheckDelinquentResponse{}, errors.ErrGeneral
	}

	resp := dto.CheckDelinquentResponse{
		LoanRefNum:   loan.LoanRefNum,
		IsDelinquent: len(installments) >= 2,
	}

	return resp, nil
}
