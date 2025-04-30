package http

import (
	"fmt"
	"os"

	// "beduwit.com/cmd/handler"
	// "beduwit.com/cmd/middleware"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	// log "github.com/sirupsen/logrus"
)

func NewHandler(loanProcessHandler *LoanProcessHandler, logger *zap.Logger) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	// Add CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		zap.L().Info("A group of walrus emerges from the ocean")
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/loans", loanProcessHandler.GetAllLoan)
	app.Post("/loans", loanProcessHandler.RegisterNewLoan)
	app.Get("/loans/:id", loanProcessHandler.GetDetailLoan)
	app.Post("/loans/:id/approve", loanProcessHandler.ApprovalLoan)
	app.Post("/loans/:id/invest", loanProcessHandler.InvestLoan)
	app.Post("/loans/:id/disburse", loanProcessHandler.DisburseLoan)
	app.Get("/loans/:id/investors", loanProcessHandler.InvestLoan)
	// app.Get("/loans", loanProcessHandler.RegisterNewLoan)
	// app.Get("/categories", categoryHandler.GetCategories)
	// app.Get("/categories/:name/subcategories", categoryHandler.GetSubCategories)
	// app.Post("/public/login", authHandler.Login)
	// app.Post("/public/register", authHandler.Register)
	// app.Post("/public/v1/login", authHandler.Login)
	// app.Post("/public/v1/register", authHandler.Register)
	// app.Post("/public/v1/connections", connectionHandler.CreateNewConnection)
	// app.Post("/public/v1/callback/brankas", openFinanceHandler.BrankasCallback)

	// protectedRoute := app.Group("/", middleware.IsLoggedIn(middleware.Config{}))
	// {
	// 	protectedRoute.Get("/v1/transactions/:id", connectionHandler.GetListTransaction)

	// 	protectedRoute.Get("/v1/transactions/:id/initial", connectionHandler.GetInitialBalance)
	// 	protectedRoute.Get("/v1/connections", connectionHandler.GetListConnection)
	// 	protectedRoute.Post("/v1/connect/step-1", connectionHandler.CreateConnectionViaAPI)
	// 	protectedRoute.Post("/v1/connect/step-2", connectionHandler.ConnectWithOTP)
	// 	protectedRoute.Get("/v1/investments", connectionHandler.GetListInvestment)

	// }
	port := "8080"
	fmt.Println("PORT FROM ENV:" + os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	fmt.Println("PORT Current: " + port)
	app.Listen(":" + port)
}
