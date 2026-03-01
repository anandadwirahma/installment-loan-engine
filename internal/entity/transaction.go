package entity

import (
	"installment-loan-engine/internal/shared/constant"
	"time"
)

type Transaction struct {
	ID            int64                      `gorm:"primaryKey;autoIncrement"`
	TrxRefNum     string                     `gorm:"type:varchar(50);not null;uniqueIndex"`
	LoanID        int64                      `gorm:"type:bigint;not null;index"`
	InstallmentID int64                      `gorm:"type:bigint;not null;index"`
	Amount        int64                      `gorm:"type:bigint;not null"`
	Status        constant.TransactionStatus `gorm:"type:varchar(10);not null"`
	CreatedAt     time.Time                  `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt     time.Time                  `gorm:"type:timestamp;autoUpdateTime"`
}

func (Transaction) TableName() string {
	return "transactions"
}
