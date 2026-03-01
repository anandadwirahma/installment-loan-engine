package repositories

import (
	"installment-loan-engine/internal/entity"

	"gorm.io/gorm"
)

type LoanRepository interface {
	BeginTx() *gorm.DB
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
	CreateWithTx(tx *gorm.DB, entity *entity.Loan) error
	GetLoanInstallmentByRefNum(refNum string) (entity.Loan, error)
	GetByRefNum(refNum string) (entity.Loan, error)
	UpdatesWithTx(tx *gorm.DB, id int64, updates map[string]interface{}) error
}

type loanRepository struct {
	gorm *gorm.DB
}

func NewLoanRepository(gorm *gorm.DB) LoanRepository {
	return &loanRepository{gorm: gorm}
}

func (r *loanRepository) BeginTx() *gorm.DB {
	return r.gorm.Begin()
}

func (r *loanRepository) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *loanRepository) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *loanRepository) CreateWithTx(tx *gorm.DB, entity *entity.Loan) error {
	return tx.Create(entity).Error
}

func (r *loanRepository) GetLoanInstallmentByRefNum(refNum string) (entity.Loan, error) {
	var loan entity.Loan

	if err := r.gorm.
		Preload("Installments", func(db *gorm.DB) *gorm.DB {
			return db.Order("installment_number asc")
		}).
		Where("loan_ref_num = ?", refNum).
		First(&loan).Error; err != nil {
		return entity.Loan{}, err
	}

	return loan, nil
}

func (r *loanRepository) GetByRefNum(refNum string) (entity.Loan, error) {
	var loan entity.Loan

	if err := r.gorm.Where("loan_ref_num = ?", refNum).First(&loan).Error; err != nil {
		return entity.Loan{}, err
	}

	return loan, nil
}

func (r *loanRepository) UpdatesWithTx(tx *gorm.DB, id int64, updates map[string]interface{}) error {
	return tx.Model(&entity.Loan{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}
