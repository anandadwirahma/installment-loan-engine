package repositories

import (
	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/constant"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(entity *entity.Transaction) error
	UpdateStatus(id int64, status constant.TransactionStatus) error
}

type transactionRepository struct {
	gorm *gorm.DB
}

func NewTransactionRepository(gorm *gorm.DB) TransactionRepository {
	return &transactionRepository{gorm: gorm}
}

func (r *transactionRepository) Create(entity *entity.Transaction) error {
	return r.gorm.Create(entity).Error
}

func (r *transactionRepository) UpdateStatus(id int64, status constant.TransactionStatus) error {
	return r.gorm.Model(&entity.Transaction{}).
		Where("id = ?", id).
		Update("status", status).Error
}
