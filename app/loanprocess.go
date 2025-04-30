package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"loan-engine.bivala.com/entity"
	"loan-engine.bivala.com/presenter"
	"loan-engine.bivala.com/repository"
)

type LoanProcessApp struct {
	LoanRepo         repository.Loan
	LoanInvestorRepo repository.LoanInvestor
}

func NewLoanApp(loanProcessRepo repository.Loan, loanInvestorRepo repository.LoanInvestor) *LoanProcessApp {
	return &LoanProcessApp{
		LoanRepo:         loanProcessRepo,
		LoanInvestorRepo: loanInvestorRepo,
	}
}

func (p *LoanProcessApp) GetDetailLoan(ctx context.Context, id int64) (*entity.Loan, error) {
	ent, err := p.LoanRepo.GetLoanByID(ctx, id)
	// Get total invested amount
	totalInvested, err := p.LoanInvestorRepo.GetTotalAlreadyInvested(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			totalInvested = 0
		} else {

			return nil, err
		}
	}
	ent.AmountInvested = totalInvested
	ent.AmountUninvested = ent.PrincipalAmount - ent.AmountInvested
	if err != nil {
		zap.L().Error("Failed to get the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return nil, err
	}
	return &ent, nil
}
func (p *LoanProcessApp) GetAllLoan(ctx context.Context, status string) ([]*entity.Loan, error) {
	ent, err := p.LoanRepo.GetAllLoan(ctx, status)
	if err != nil {
		zap.L().Error("Failed to get the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return nil, err
	}
	return ent, nil
}
func (p *LoanProcessApp) RegisterNewLoan(ctx context.Context, data presenter.ProposedLoan) error {
	roi := data.PrincipalAmount + (data.PrincipalAmount * data.Rate / 100)
	insertData := entity.Loan{
		UniqueID:        uuid.NewString(),
		BorrowerID:      data.BorrowerID,
		PrincipalAmount: data.PrincipalAmount,
		Rate:            data.Rate,
		ROI:             roi,
		CurrentStatus:   "PROPOSAL",
	}
	fmt.Println(insertData)
	if err := p.LoanRepo.RegisterNewLoan(ctx, insertData); err != nil {
		zap.L().Error("Failed to register the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}
	return nil
}
func (p *LoanProcessApp) ApprovedLoan(ctx context.Context, data presenter.ApprovalLoan) error {
	format := "2006-01-02"
	date, _ := time.Parse(format, data.ApprovalDate)
	loan, err := p.LoanRepo.GetLoanByID(ctx, data.LoanID)
	if err != nil {
		zap.L().Error("Failed to get the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}

	if loan.CurrentStatus != "PROPOSAL" {
		zap.L().Error("Current Status is not proposal",
			zap.String("module", "application loan process"),
			zap.Time("at", time.Now()))
		return errors.New("Loan already approved or rejected")
	}
	approveData := entity.Loan{
		ApprovedBy:    &data.EmployeeID,
		ApprovedDate:  &date,
		ProofImage:    &data.ProofImage,
		CurrentStatus: "APPROVED",
	}
	if err := p.LoanRepo.ApproveLoan(ctx, approveData, data.LoanID); err != nil {
		zap.L().Error("Failed to approve the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}
	return nil
}
func (p *LoanProcessApp) InvestLoan(ctx context.Context, data presenter.InvestLoan) error {
	// Start transaction
	tx, err := p.LoanRepo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			p.LoanRepo.RollbackTx(tx)
		}
	}()

	// Get loan with lock
	loanData, err := p.LoanRepo.GetLoanByIDWithLock(ctx, data.LoanID)
	if err != nil {
		return err
	}
	if loanData.CurrentStatus != "APPROVED" {
		return errors.New("Loan not approved yet")
	}
	// Simulate longer processing time to test locking mechanism
	time.Sleep(10 * time.Second)

	// Get total invested amount
	total, err := p.LoanInvestorRepo.GetTotalAlreadyInvested(ctx, data.LoanID)
	if err != nil {
		if err.Error() == "record not found" {
			total = 0
		} else {

			return err
		}
	}

	// Validate investment amount
	if loanData.PrincipalAmount < total+data.InvestAmount {
		return errors.New("total invest over from principal amount")
	}

	// Prepare investment data
	investData := entity.LoanInvestor{
		LoanID:         data.LoanID,
		InvestorID:     data.InvestorID,
		AmountInvested: data.InvestAmount,
	}

	// Update loan status if fully invested
	if loanData.PrincipalAmount == total+data.InvestAmount {
		timeNow := time.Now()
		loanData.CurrentStatus = "INVESTED"
		loanData.InvestedDate = &timeNow
		if err := p.LoanRepo.InvestLoan(ctx, loanData); err != nil {
			zap.L().Error("Failed to invest the loan",
				zap.String("module", "application loan process"),
				zap.String("error", err.Error()),
				zap.Time("at", time.Now()))
			return err
		}
	}

	// Create investment record
	if err := p.LoanInvestorRepo.SetNewInvestor(ctx, tx, investData); err != nil {
		zap.L().Error("Failed to create set new investor",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}

	// Commit transaction
	if err := p.LoanRepo.CommitTx(tx); err != nil {
		return err
	}

	return nil
}
func (p *LoanProcessApp) DisburseLoan(ctx context.Context, data presenter.DisburseLoan) error {
	format := "2006-01-02"
	date, _ := time.Parse(format, data.DateDisbursement)
	loan, err := p.LoanRepo.GetLoanByID(ctx, data.LoanID)
	fmt.Println(loan, "err", err)
	if err != nil {
		zap.L().Error("Failed to get the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}

	if loan.CurrentStatus != "INVESTED" {
		zap.L().Error("Current Status is not INVESTED",
			zap.String("module", "application loan process"),
			zap.Time("at", time.Now()))
		return errors.New("Loan still not invested")
	}
	approveData := entity.Loan{
		DisbursedBy:           &data.EmployeeID,
		DisbursedDate:         &date,
		SignedAgreementLetter: &data.LoanAgreementFile,
		CurrentStatus:         "DISBURSED",
	}
	if err := p.LoanRepo.DisburseLoan(ctx, approveData, data.LoanID); err != nil {
		zap.L().Error("Failed to approve the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}
	return nil
}
