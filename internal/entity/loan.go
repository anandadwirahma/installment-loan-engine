package entity

import (
	"time"

	"installment-loan-engine/internal/shared/constant"
)

type Loan struct {
	ID                   int64               `gorm:"primaryKey;autoIncrement"`
	BorrowerRefNum       string              `gorm:"type:varchar(50);not null"`
	LoanRefNum           string              `gorm:"type:varchar(50);not null;index;uniqueIndex"`
	PrincipalAmount      int64               `gorm:"type:bigint;not null"`
	InterestRate         float64             `gorm:"type:decimal(5,2);not null"`
	InterestType         string              `gorm:"type:varchar(20);not null"`
	TenorWeeks           int16               `gorm:"type:smallint;not null"`
	TotalInterestAmount  int64               `gorm:"type:bigint;not null"`
	TotalRepaymentAmount int64               `gorm:"type:bigint;not null"`
	OutstandingAmount    int64               `gorm:"type:bigint;not null"`
	Status               constant.LoanStatus `gorm:"type:varchar(10);not null"`
	CreatedAt            time.Time           `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt            time.Time           `gorm:"type:timestamp;autoUpdateTime"`

	Installments []Installment `gorm:"foreignKey:LoanID;references:ID"`
}

func (Loan) TableName() string {
	return "loans"
}
