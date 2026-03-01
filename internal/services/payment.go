package services

import (
	"time"

	"installment-loan-engine/internal/dto"
	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/constant"
	"installment-loan-engine/internal/shared/errors"
	"installment-loan-engine/internal/shared/logger"
)

func (s *loanService) PayInstallment(req dto.PayInstallmentRequest) (dto.PayInstallmentResponse, error) {
	loan, err := s.loanRepo.GetByRefNum(req.LoanRefNum)
	if err != nil {
		logger.Errorf("[service.PayInstallment] Error fetching loan for RefNum %s: %v", req.LoanRefNum, err)
		return dto.PayInstallmentResponse{}, errors.ErrNotFound
	}

	if loan.Status == constant.LoanStatusClosed {
		logger.Errorf("[service.PayInstallment] Loan is closed for RefNum %s", req.LoanRefNum)
		return dto.PayInstallmentResponse{}, errors.ErrLoanClosed
	}

	installments, err := s.installmentRepo.GetOutstandingInstallments(loan.ID)
	if err != nil {
		logger.Errorf("[service.PayInstallment] Error fetching outstanding installments for RefNum %s: %v", req.LoanRefNum, err)
		return dto.PayInstallmentResponse{}, errors.ErrGeneral
	}

	if len(installments) == 0 {
		logger.Errorf("[service.PayInstallment] No outstanding installment for RefNum %s", req.LoanRefNum)
		return dto.PayInstallmentResponse{}, errors.ErrNoOutstandingInstallment
	}

	var (
		totalAmountExpected    int64
		paidInstallmentNumbers []int16
		now                    = time.Now()

		transactions []*entity.Transaction
	)

	trxRefNum := "aiueo" //TODO: create function generator
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
		logger.Errorf("[service.PayInstallment] Amount %d is not enough or exceeds outstanding balance for RefNum %s", req.Amount, req.LoanRefNum)
		return dto.PayInstallmentResponse{}, errors.ErrInvalidAmount
	}

	if err = s.transactionRepo.Create(transactions); err != nil {
		logger.Errorf("[service.PayInstallment] Error creating transaction for RefNum %s: %v", req.LoanRefNum, err)
		return dto.PayInstallmentResponse{}, errors.ErrGeneral
	}

	if err = s.doPayProcess(installments, loan, now); err != nil {
		logger.Errorf("[service.PayInstallment] Error processing payment for RefNum %s: %v", req.LoanRefNum, err)

		s.transactionRepo.UpdateStatusByRefNum(trxRefNum, constant.TransactionStatusFailed)
		return dto.PayInstallmentResponse{}, errors.ErrGeneral
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
			logger.Errorf("[service.doPayProcess] Error updating installment status for RefNum %s: %v", loan.LoanRefNum, err)
			return err
		}
	}

	if (loan.TotalRepaymentAmount - paidTotalAmount) == 0 {
		if err := s.loanRepo.UpdateStatusWithTx(tx, loan.ID, constant.LoanStatusClosed); err != nil {
			s.installmentRepo.RollbackTx(tx)
			logger.Errorf("[service.doPayProcess] Error updating loan status for RefNum %s: %v", loan.LoanRefNum, err)
			return err
		}
	}

	if err := s.installmentRepo.CommitTx(tx); err != nil {
		logger.Errorf("[service.doPayProcess] Error committing transaction for RefNum %s: %v", loan.LoanRefNum, err)
		return err
	}

	return nil
}
