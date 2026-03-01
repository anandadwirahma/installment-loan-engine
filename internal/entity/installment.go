package entity

import (
	"time"

	"installment-loan-engine/internal/shared/constant"
)

type Installment struct {
	ID                int64                      `gorm:"primaryKey;autoIncrement"`
	LoanID            int64                      `gorm:"type:bigint;not null;index;uniqueIndex:idx_loan_installment"`
	InstallmentNumber int16                      `gorm:"type:smallint;not null;uniqueIndex:idx_loan_installment"`
	DueDate           time.Time                  `gorm:"type:date;not null;index"`
	PrincipalAmount   int64                      `gorm:"type:bigint;not null"`
	InterestAmount    int64                      `gorm:"type:bigint;not null"`
	TotalAmount       int64                      `gorm:"type:bigint;not null"`
	PaidAt            *time.Time                 `gorm:"type:timestamp;default:null"`
	Status            constant.InstallmentStatus `gorm:"type:varchar(10);not null"`
	CreatedAt         time.Time                  `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt         time.Time                  `gorm:"type:timestamp;autoUpdateTime"`
}

func (Installment) TableName() string {
	return "installments"
}
