package gorm

import (
	"context"

	"gorm.io/gorm"
	"loan-engine.bivala.com/entity"
	"loan-engine.bivala.com/repository"
)

type LoanInvestorPsqlRepo struct {
	Conn *gorm.DB
}

func NewLoanInvestorPsqlRepo(conn *gorm.DB) repository.LoanInvestor {
	return &LoanInvestorPsqlRepo{
		Conn: conn,
	}
}

type Result struct {
	Total int64
}

func (p *LoanInvestorPsqlRepo) GetTotalAlreadyInvested(ctx context.Context, loanID int64) (int64, error) {

	var result Result
	resultQuery := p.Conn.Model(&entity.LoanInvestor{}).Select("sum(amount_invested) as total").Where("loan_id=?", loanID).Group("loan_id").First(&result)
	if resultQuery.Error != nil {
		return 0, resultQuery.Error
	}
	return result.Total, nil
}
func (p *LoanInvestorPsqlRepo) SetNewInvestor(ctx context.Context, data entity.LoanInvestor) error {
	return p.Conn.Create(&data).Error
}
