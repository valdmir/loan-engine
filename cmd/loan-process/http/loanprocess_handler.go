package http

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"loan-engine.bivala.com/app"
	"loan-engine.bivala.com/presenter"
)

type LoanProcessHandler struct {
	LoanProcessApp *app.LoanProcessApp
}

func NewLoanProcessHandler(loanProcessApp *app.LoanProcessApp) *LoanProcessHandler {
	return &LoanProcessHandler{
		LoanProcessApp: loanProcessApp,
	}
}
func (h *LoanProcessHandler) RegisterNewLoan(c *fiber.Ctx) error {
	var postData presenter.ProposedLoan
	err := c.BodyParser(&postData)
	if err != nil {
		zap.L().Error("Error Body Parser")
		return c.Status(400).JSON(err)
	}
	if err := h.LoanProcessApp.RegisterNewLoan(c.Context(), postData); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"msg": err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"msg":       "Proposal New Loan Submitted",
		"unique_id": "testing",
	})
}
func (h *LoanProcessHandler) ApprovalLoan(c *fiber.Ctx) error {
	var postData presenter.ApprovalLoan
	fmt.Println("approval loan")
	err := c.BodyParser(&postData)
	if err != nil {
		zap.L().Error("Error Body Parser")
		return c.Status(400).JSON(err)
	}
	file, err := c.FormFile("proof_image")
	if err != nil {
		// Handle error
		return err
	}
	fileReader, _ := file.Open()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, fileReader); err != nil {
		// return nil, err
	}
	loanID, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	postData.LoanID = loanID
	postData.ProofImage = buf.Bytes()
	if err := h.LoanProcessApp.ApprovedLoan(c.Context(), postData); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"msg": err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"msg": "This Loan get approve",
		// "unique_id": "testing",
	})
}

func (h *LoanProcessHandler) InvestLoan(c *fiber.Ctx) error {
	var postData presenter.InvestLoan
	fmt.Println("invest loan")
	err := c.BodyParser(&postData)
	if err != nil {
		zap.L().Error("Error Body Parser")
		return c.Status(400).JSON(err)
	}

	loanID, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	postData.LoanID = loanID
	if err := h.LoanProcessApp.InvestLoan(c.Context(), postData); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"msg": "This Loan get approve",
		// "unique_id": "testing",
	})
}
