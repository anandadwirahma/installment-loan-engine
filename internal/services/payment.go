package services

import (
	"fmt"
	"time"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/constant"
)

func (s *loanService) PayInstallment(req dto.PayInstallmentRequest) (dto.PayInstallmentResponse, error) {
	loan, err := s.loanRepo.GetByRefNum(req.LoanRefNum)
	if err != nil {
		return dto.PayInstallmentResponse{}, err
	}

	if loan.Status == constant.LoanStatusClosed {
		return dto.PayInstallmentResponse{}, fmt.Errorf("Loan is closed")
	}

	installments, err := s.installmentRepo.GetOutstandingInstallments(loan.ID)
	if err != nil {
		return dto.PayInstallmentResponse{}, err
	}

	if len(installments) == 0 {
		return dto.PayInstallmentResponse{}, fmt.Errorf("No outstanding installment")
	}

	var (
		totalAmountExpected    int64
		paidInstallmentNumbers []int16
		now                    time.Time
	)

	for _, i := range installments {
		totalAmountExpected += i.TotalAmount
		paidInstallmentNumbers = append(paidInstallmentNumbers, i.InstallmentNumber)
	}

	if totalAmountExpected != req.Amount {
		return dto.PayInstallmentResponse{}, fmt.Errorf("Amount %d is not enough or exceeds outstanding balance", req.Amount)
	}

	if err = s.doPayProcess(installments, now); err != nil {
		return dto.PayInstallmentResponse{}, fmt.Errorf("Error process payment")
	}

	return dto.PayInstallmentResponse{
		LoanRefNum:           loan.LoanRefNum,
		PaidAmount:           req.Amount,
		PaidInstallments:     paidInstallmentNumbers,
		RemainingOutstanding: 0,
		PaidAt:               now.Format(time.RFC3339),
	}, nil
}

func (s *loanService) doPayProcess(installments []entity.Installment, paidAt time.Time) error {
	tx := s.installmentRepo.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			s.installmentRepo.RollbackTx(tx)
		}
	}()

	for _, i := range installments {
		if err := s.installmentRepo.UpdateStatusWithTx(tx, i.ID, constant.InstallmentStatusPaid, paidAt); err != nil {
			s.installmentRepo.RollbackTx(tx)
			return err
		}
	}

	if err := s.installmentRepo.CommitTx(tx); err != nil {
		return err
	}

	// TODO: Insert record to table transaction

	return nil
}
