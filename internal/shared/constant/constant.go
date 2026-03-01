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
)

type TransactionStatus string

const (
	TransactionStatusInit    TransactionStatus = "INIT"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)
