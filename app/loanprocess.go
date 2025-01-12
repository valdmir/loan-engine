package app

import (
	"context"
	"errors"
	"fmt"
	"time"

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
func (p *LoanProcessApp) GetAllLoan(ctx context.Context, status string) ([]*entity.Loan, error) {
	return nil, nil
}
func (p *LoanProcessApp) RegisterNewLoan(ctx context.Context, data presenter.ProposedLoan) error {
	insertData := entity.Loan{
		BorrowerID:      data.BorrowerID,
		PrincipalAmount: data.PrincipalAmount,
		Rate:            data.Rate,
		ROI:             data.ROI,
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
	loanData, err := p.LoanRepo.GetLoanByID(ctx, data.LoanID)
	if err != nil {
		return err
	}
	total, _ := p.LoanInvestorRepo.GetTotalAlreadyInvested(ctx, data.LoanID)
	fmt.Println(total, data.InvestAmount, loanData.PrincipalAmount)
	if loanData.PrincipalAmount < total+data.InvestAmount {
		return errors.New("total invest over from principal amount")
	}
	investData := entity.LoanInvestor{
		LoanID:         data.LoanID,
		InvestorID:     data.InvestorID,
		AmountInvested: data.InvestAmount,
	}
	if err := p.LoanInvestorRepo.SetNewInvestor(ctx, investData); err != nil {
		zap.L().Error("Failed to invest the loan",
			zap.String("module", "application loan process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()))
		return err
	}
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

	return nil
}
func (p *LoanProcessApp) DisburseLoan(ctx context.Context, data presenter.DisburseLoan) error {
	return nil
}
