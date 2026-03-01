package entity

import (
	"installment-loan-engine/internal/shared/constant"
	"time"
)

type Loan struct {
	ID                   int64               `gorm:"primaryKey;autoIncrement;type:bigserial"`
	BorrowerRefNum       string              `gorm:"type:varchar(50);not null;index"`
	LoanRefNum           string              `gorm:"type:varchar(50);not null;uniqueIndex"`
	PrincipalAmount      int64               `gorm:"type:bigint;not null"`
	InterestRate         float64             `gorm:"type:decimal(5,2);not null"`
	InterestType         string              `gorm:"type:varchar(20);not null"`
	TenorWeeks           int16               `gorm:"type:smallint;not null"`
	TotalInterestAmount  int64               `gorm:"type:bigint;not null"`
	TotalRepaymentAmount int64               `gorm:"type:bigint;not null"`
	Status               constant.LoanStatus `gorm:"type:varchar(10);not null"`
	Installments         []Installment
	CreatedAt            time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (Loan) TableName() string {
	return "loans"
}
