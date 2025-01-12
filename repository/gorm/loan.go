package gorm

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"loan-engine.bivala.com/entity"
	"loan-engine.bivala.com/repository"
)

type LoanPsqlRepo struct {
	Conn *gorm.DB
}

func NewLoanPsqlRepo(conn *gorm.DB) repository.Loan {
	return &LoanPsqlRepo{
		Conn: conn,
	}
}
func (p *LoanPsqlRepo) RegisterNewLoan(ctx context.Context, data entity.Loan) error {
	return p.Conn.Create(&data).Error
}
func (p *LoanPsqlRepo) GetAllLoan(ctx context.Context, status string) ([]*entity.Loan, error) {
	return nil, nil
}
func (p *LoanPsqlRepo) ApproveLoan(ctx context.Context, data entity.Loan, loanID int64) error {
	tempLoan := entity.Loan{ID: loanID}
	result := p.Conn.First(&tempLoan)
	if result.Error != nil {
		fmt.Println(result.Error)
		zap.L().Error("Failed to get the loan",
			zap.String("module", "application loan process"),
			zap.String("error", result.Error.Error()),
			zap.Time("at", time.Now()))
		return result.Error
	}
	fmt.Println(tempLoan)
	resultUpdate := p.Conn.Model(&tempLoan).Select("approved_by", "approved_date", "proof_image", "current_status").Updates(data)
	if resultUpdate.Error != nil {
		zap.L().Error("Failed to update the loan",
			zap.String("module", "application loan process"),
			zap.String("error", resultUpdate.Error.Error()),
			zap.Time("at", time.Now()))
		return resultUpdate.Error
	}
	return nil
}
func (p *LoanPsqlRepo) InvestLoan(ctx context.Context, data entity.Loan) error {

	result := p.Conn.Save(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
		zap.L().Error("Failed to invest loan",
			zap.String("module", "application loan process"),
			zap.String("error", result.Error.Error()),
			zap.Time("at", time.Now()))
		return result.Error
	}
	return nil
}
func (p *LoanPsqlRepo) DisburseLoan(ctx context.Context, data entity.Loan) error {
	return nil
}
func (p *LoanPsqlRepo) GetLoanByID(ctx context.Context, loanID int64) (entity.Loan, error) {
	tempLoan := entity.Loan{ID: loanID}
	result := p.Conn.First(&tempLoan)
	if result.Error != nil {
		fmt.Println(result.Error)
		zap.L().Error("Failed to get the loan",
			zap.String("module", "application loan process"),
			zap.String("error", result.Error.Error()),
			zap.Time("at", time.Now()))
		return tempLoan, result.Error
	}
	return tempLoan, nil
}
