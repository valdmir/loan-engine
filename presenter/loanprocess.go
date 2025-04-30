package presenter

type ProposedLoan struct {
	BorrowerID      int64 `json:"borrower_id" validate:"required"`
	PrincipalAmount int64 `json:"principal_amount" validate:"required"`
	Rate            int64 `json:"rate" validate:"required,gte=0,lte=100"`
}
type ApprovalLoan struct {
	LoanID       int64  `json:"loan_id" validate:"required"`
	ProofImage   []byte `json:"proof_image" validate:"required" form:"proof_image"`
	EmployeeID   int64  `json:"employee_id" validate:"required,gt=0" form:"employee_id"`
	ApprovalDate string `json:"approval_date" validate:"required,datetime=2006-01-02" form:"approval_date"`
}
type InvestLoan struct {
	InvestorID   int64 `json:"investor_id" validate:"required"`
	InvestAmount int64 `json:"invest_amount" validate:"required"`
	LoanID       int64 `json:"loan_id" validate:"required"`
}
type DisburseLoan struct {
	LoanID            int64  `json:"loan_id" validate:"required"`
	LoanAgreementFile []byte `json:"loan_agreement_file" validate:"required" form:"loan_agreement_file"`
	DateDisbursement  string `json:"date_disbursement" validate:"required" form:"date_disbursement"`
	EmployeeID        int64  `json:"employee_id" validate:"required" form:"employee_id"`
}
