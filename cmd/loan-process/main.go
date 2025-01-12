package main

import (
	"os"
	"time"

	"github.com/subosito/gotenv"
	"go.uber.org/zap"
	"loan-engine.bivala.com/app"
	"loan-engine.bivala.com/cmd/loan-process/http"
	"loan-engine.bivala.com/db"
	"loan-engine.bivala.com/repository/gorm"
)

func main() {
	err := gotenv.Load()
	logger, _ := zap.NewProduction()

	if err != nil {
		zap.L().Fatal("Failed to start Backend Loan Process",
			zap.String("module", "loan-process"),
			zap.String("error", err.Error()),
			zap.Time("at", time.Now()),
		)
	}
	dbURL := os.Getenv("DATABASE_URL")

	// Infra setup
	gormInfra := db.NewPsqlGormInfra(dbURL)

	loanProcessRepo := gorm.NewLoanPsqlRepo(gormInfra)
	loanInvestorRepo := gorm.NewLoanInvestorPsqlRepo(gormInfra)
	loanProcessApp := app.NewLoanApp(loanProcessRepo, loanInvestorRepo)

	connectionHandler := http.NewLoanProcessHandler(loanProcessApp)
	http.NewHandler(connectionHandler, logger)
}
