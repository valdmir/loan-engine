package repository

import (
	"context"

	"gorm.io/gorm"
	"loan-engine.bivala.com/entity"
)

type Loan interface {
	RegisterNewLoan(ctx context.Context, data entity.Loan) error
	GetAllLoan(ctx context.Context, status string) ([]*entity.Loan, error)
	ApproveLoan(ctx context.Context, data entity.Loan, loanID int64) error
	InvestLoan(ctx context.Context, data entity.Loan) error
	DisburseLoan(ctx context.Context, data entity.Loan, loanID int64) error
	GetLoanByID(ctx context.Context, loanID int64) (entity.Loan, error)
	GetLoanByIDWithLock(ctx context.Context, loanID int64) (entity.Loan, error)
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
}
type LoanInvestor interface {
	GetTotalAlreadyInvested(ctx context.Context, loanID int64) (int64, error)
	SetNewInvestor(ctx context.Context, tx *gorm.DB, data entity.LoanInvestor) error
}
