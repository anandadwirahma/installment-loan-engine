package dto

type CreateLoanRequest struct {
	BorrowerRefNum  string `json:"borrower_ref_num"`
	PrincipalAmount int64  `json:"principal_amount"`
}

type CreateLoanResponse struct {
	LoanRefNum           string `json:"loan_ref_num"`
	PrincipalAmount      int64  `json:"principal_amount"`
	TotalInterestAmount  int64  `json:"total_interest_amount"`
	TotalRepaymentAmount int64  `json:"total_repayment_amount"`
	WeeklyInstallment    int64  `json:"weekly_installment"`
	Status               string `json:"status"`
}
