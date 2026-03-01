package entity

import (
	"installment-loan-engine/internal/shared/constant"
	"time"
)

type Installment struct {
	ID                int64                      `gorm:"primaryKey"`
	LoanID            int64                      `gorm:"not null;index;uniqueIndex:idx_loan_installment"`
	InstallmentNumber int16                      `gorm:"not null;uniqueIndex:idx_loan_installment"`
	DueDate           time.Time                  `gorm:"type:date;not null;index"`
	PrincipalAmount   int64                      `gorm:"not null"`
	InterestAmount    int64                      `gorm:"not null"`
	TotalAmount       int64                      `gorm:"not null"`
	PaidAt            *time.Time                 `gorm:"default:null"`
	Status            constant.InstallmentStatus `gorm:"type:varchar(10);not null"`
	CreatedAt         time.Time                  `gorm:"autoCreateTime"`
	UpdatedAt         time.Time                  `gorm:"autoUpdateTime"`
}

func (Installment) TableName() string {
	return "installments"
}
