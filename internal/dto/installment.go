package dto

type GetInstallmentRequest struct {
	LoanRefNum string `json:"loan_ref_num"`
}

type GetInstallmentResponse struct {
	LoanRefNum        string        `json:"loan_ref_num"`
	TotalInstallments int64         `json:"total_installments"`
	Installments      []Installment `json:"installments"`
	Summary           Summary       `json:"summary"`
}

type Installment struct {
	InstallmentNumber int     `json:"installment_number"`
	PrincipalAmount   int64   `json:"principal_amount"`
	InterestAmount    int64   `json:"interest_amount"`
	TotalAmount       int64   `json:"total_amount"`
	DueDate           string  `json:"due_date"`
	Status            string  `json:"status"`
	PaidAt            *string `json:"paid_at"`
}

type Summary struct {
	TotalOutstanding int64 `json:"total_outstanding"`
	IsDelinquent     bool  `json:"is_delinquent"`
}

type GetOutstandingRequest struct {
	LoanRefNum string `json:"loan_ref_num"`
}

type GetOutstandingResponse struct {
	LoanRefNum        string `json:"loan_ref_num"`
	OutstandingAmount int64  `json:"outstanding_amount"`
}

type CheckDelinquentRequest struct {
	LoanRefNum string `json:"loan_ref_num"`
}

type CheckDelinquentResponse struct {
	LoanRefNum   string `json:"loan_ref_num"`
	IsDelinquent bool   `json:"is_delinquent"`
}

type PayInstallmentRequest struct {
	LoanRefNum string `json:"loan_ref_num"`
	Amount     int64  `json:"amount"`
}

type PayInstallmentResponse struct {
	LoanRefNum           string  `json:"loan_ref_num"`
	TrxRefNum            string  `json:"trx_ref_num"`
	PaidAmount           int64   `json:"paid_amount"`
	PaidInstallments     []int16 `json:"paid_installments"`
	RemainingOutstanding int64   `json:"remaining_outstanding"`
	PaidAt               string  `json:"paid_at"`
}
