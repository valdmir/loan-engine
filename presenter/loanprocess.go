package presenter

import (
	"time"
)

type ProposedLoan struct {
	BorrowerID      int64 `json:"borrower_id" validate:"required"`
	PrincipalAmount int64 `json:"principal_amount" validate:"required"`
	Rate            int64 `json:"rate" validate:"required"`
	ROI             int64 `json:"roi" validate:"required"`
}
type ApprovalLoan struct {
	LoanID       int64  `json:"loan_id"`
	ProofImage   []byte `json:"proof_image" validate:"required" form:"proof_image"`
	EmployeeID   int64  `json:"employee_id" validate:"required" form:"employee_id"`
	ApprovalDate string `json:"approval_date" validate:"required" form:"approval_date"`
}
type InvestLoan struct {
	InvestorID   int64 `json:"investor_id"  validate:"required"`
	InvestAmount int64 `json:"invest_amount"  validate:"required"`
	LoanID       int64 `json:"loan_id"`
}
type DisburseLoan struct {
	LoanAgreementFile string    `json:"loan_agreement_file"`
	DateDisbursement  time.Time `json:"date_disbursement"`
	EmployeeID        string    `json:"employee_id"`
}
