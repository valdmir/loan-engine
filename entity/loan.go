package entity

import (
	"time"
)

type Loan struct {
	ID                    int64      `json:"id" gorm:"primaryKey,<-"`
	UniqueID              string     `json:"unique_id" db:"unique_id"`
	BorrowerID            int64      `json:"borrower_id" db:"borrower_id"`
	PrincipalAmount       int64      `json:"principal_amount" db:"principal_amount"`
	Rate                  int64      `json:"rate" db:"rate"`
	ROI                   int64      `json:"roi" db:"roi"`
	CurrentStatus         string     `json:"current_status" db:"current_status"`
	ProofImage            *[]byte    `json:"proof_image" db:"proof_image"`
	ApprovedBy            *int64     `json:"approved_by"`
	ApprovedDate          *time.Time `json:"approved_date"`
	InvestedDate          *time.Time `json:"invested_date"`
	DisbursedBy           *int64     `json:"disbursed_by" db:"disbursed_by"`
	DisbursedDate         *string    `json:"disbursed_date" db:"disbursed_date"`
	SignedAgreementLetter *string    `json:"signed_agreement_letter" db:"signed_agreement_letter"`
	CreatedAt             time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at" db:"updated_at"`
}
type LoanInvestor struct {
	LoanID         int64 `json:"loan_id"`
	InvestorID     int64 `json:"investor_id"`
	AmountInvested int64 `json:"amount_invested"`
	// EmailAgreement string `json:"email_agreement"`
}
