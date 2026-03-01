package repositories

import (
	"time"

	"gorm.io/gorm"

	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/constant"
)

type InstallmentRepository interface {
	BeginTx() *gorm.DB
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
	CreateWithTx(tx *gorm.DB, entity []*entity.Installment) error
	GetOutstandingAmount(loanId int64) (int64, error)
	GetOutstandingInstallments(loanId int64) ([]entity.Installment, error)
	GetOverdueInstallment(loanId int64) ([]entity.Installment, error)
	UpdateStatusWithTx(tx *gorm.DB, id int64, status constant.InstallmentStatus, paidAt time.Time) error
}

type installmentRepository struct {
	gorm *gorm.DB
}

func NewInstallmentRepository(gorm *gorm.DB) InstallmentRepository {
	return &installmentRepository{gorm: gorm}
}

func (r *installmentRepository) BeginTx() *gorm.DB {
	return r.gorm.Begin()
}

func (r *installmentRepository) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *installmentRepository) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *installmentRepository) CreateWithTx(tx *gorm.DB, entity []*entity.Installment) error {
	return tx.Create(entity).Error
}

func (r *installmentRepository) GetOutstandingAmount(loanId int64) (int64, error) {
	var outstanding int64

	if err := r.gorm.Model(&entity.Installment{}).
		Where("loan_id = ?", loanId).
		Where("status = ?", constant.InstallmentStatusPending).
		Select("SUM(total_amount)").
		Scan(&outstanding).Error; err != nil {
		return 0, err
	}

	return outstanding, nil
}

func (r *installmentRepository) GetOutstandingInstallments(loanId int64) ([]entity.Installment, error) {
	var (
		installments []entity.Installment
	)

	if err := r.gorm.Model(&entity.Installment{}).
		Where("loan_id = ?", loanId).
		Where("status = ?", constant.InstallmentStatusPending).
		Where("due_date <= CURRENT_DATE + INTERVAL '7 days'").
		Order("installment_number ASC").
		Find(&installments).Error; err != nil {
		return nil, err
	}

	return installments, nil
}

func (r *installmentRepository) GetOverdueInstallment(loanId int64) ([]entity.Installment, error) {
	var (
		installments []entity.Installment
		now          = time.Now()
	)

	if err := r.gorm.Model(&entity.Installment{}).
		Where("loan_id = ?", loanId).
		Where("status = ?", constant.InstallmentStatusPending).
		Where("due_date < ?", now).
		Find(&installments).Error; err != nil {
		return nil, err
	}

	return installments, nil
}

func (r *installmentRepository) UpdateStatusWithTx(tx *gorm.DB, id int64, status constant.InstallmentStatus, paidAt time.Time) error {
	return tx.Model(&entity.Loan{}).
		Where("id = ?", id).
		Updates(&entity.Installment{Status: status, PaidAt: &paidAt}).
		Error
}
