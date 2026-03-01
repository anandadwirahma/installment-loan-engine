package repositories

import (
	"gorm.io/gorm"

	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/constant"
	"installment-loan-engine/internal/shared/logger"
)

type TransactionRepository interface {
	Create(entity []*entity.Transaction) error
	UpdateStatusByRefNum(trxRefNum string, status constant.TransactionStatus) error
}

type transactionRepository struct {
	gorm *gorm.DB
}

func NewTransactionRepository(gorm *gorm.DB) TransactionRepository {
	return &transactionRepository{gorm: gorm}
}

func (r *transactionRepository) Create(entity []*entity.Transaction) error {
	return r.gorm.Create(entity).Error
}

func (r *transactionRepository) UpdateStatusByRefNum(trxRefNum string, status constant.TransactionStatus) error {
	err := r.gorm.Model(&entity.Transaction{}).
		Where("trx_ref_num = ?", trxRefNum).
		Update("status", status).Error
	if err != nil {
		logger.Errorf("[transactionRepository.UpdateStatusByRefNum] Error updating transaction status for RefNum %s: %v", trxRefNum, err)
		return err
	}

	return nil
}
