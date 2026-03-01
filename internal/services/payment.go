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
		now                    = time.Now()

		transactions []*entity.Transaction
	)

	trxRefNum := "aiueo"
	for _, i := range installments {
		totalAmountExpected += i.TotalAmount
		paidInstallmentNumbers = append(paidInstallmentNumbers, i.InstallmentNumber)

		// compose transaction data
		transactions = append(transactions, &entity.Transaction{
			TrxRefNum:     trxRefNum,
			LoanID:        i.LoanID,
			InstallmentID: i.ID,
			Amount:        i.TotalAmount,
			Status:        constant.TransactionStatusPending,
		})
	}

	if totalAmountExpected != req.Amount {
		return dto.PayInstallmentResponse{}, fmt.Errorf("Amount %d is not enough or exceeds outstanding balance", req.Amount)
	}

	if err = s.transactionRepo.Create(transactions); err != nil {
		return dto.PayInstallmentResponse{}, fmt.Errorf("Failed init db trx")
	}

	if err = s.doPayProcess(installments, loan, now); err != nil {
		s.transactionRepo.UpdateStatusByRefNum(trxRefNum, constant.TransactionStatusFailed)
		return dto.PayInstallmentResponse{}, fmt.Errorf("Payment process error")
	}

	s.transactionRepo.UpdateStatusByRefNum(trxRefNum, constant.TransactionStatusSuccess)

	return dto.PayInstallmentResponse{
		TrxRefNum:            trxRefNum,
		LoanRefNum:           loan.LoanRefNum,
		PaidAmount:           req.Amount,
		PaidInstallments:     paidInstallmentNumbers,
		RemainingOutstanding: loan.TotalRepaymentAmount - req.Amount,
		PaidAt:               now.Format(time.RFC3339),
	}, nil
}

func (s *loanService) doPayProcess(installments []entity.Installment, loan entity.Loan, paidAt time.Time) error {
	tx := s.installmentRepo.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			s.installmentRepo.RollbackTx(tx)
		}
	}()

	var paidTotalAmount int64
	for _, i := range installments {
		paidTotalAmount += i.TotalAmount
		if err := s.installmentRepo.UpdateStatusWithTx(tx, i.ID, constant.InstallmentStatusPaid, paidAt); err != nil {
			s.installmentRepo.RollbackTx(tx)
			return err
		}
	}

	if (loan.TotalRepaymentAmount - paidTotalAmount) == 0 {
		if err := s.loanRepo.UpdateStatusWithTx(tx, loan.ID, constant.LoanStatusClosed); err != nil {
			s.installmentRepo.RollbackTx(tx)
			return err
		}
	}

	if err := s.installmentRepo.CommitTx(tx); err != nil {
		return err
	}

	return nil
}
