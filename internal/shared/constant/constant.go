package constant

type LoanStatus string

const (
	LoanStatusActive     LoanStatus = "ACTIVE"
	LoanStatusClosed     LoanStatus = "CLOSED"
	LoanStatusDelinquent LoanStatus = "DELINQUENT"
)

type InstallmentStatus string

const (
	InstallmentStatusPending InstallmentStatus = "PENDING"
	InstallmentStatusPaid    InstallmentStatus = "PAID"
	InstallmentStatusOverdue                   = "OVERDUE" // it's used by logic calculation
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)
