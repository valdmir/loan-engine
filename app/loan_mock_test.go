// Code generated by MockGen. DO NOT EDIT.
// Source: repository/loan.go

// Package repository is a generated GoMock package.
package app

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "loan-engine.bivala.com/entity"
)

// MockLoan is a mock of Loan interface.
type MockLoan struct {
	ctrl     *gomock.Controller
	recorder *MockLoanMockRecorder
}

// MockLoanMockRecorder is the mock recorder for MockLoan.
type MockLoanMockRecorder struct {
	mock *MockLoan
}

// NewMockLoan creates a new mock instance.
func NewMockLoan(ctrl *gomock.Controller) *MockLoan {
	mock := &MockLoan{ctrl: ctrl}
	mock.recorder = &MockLoanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoan) EXPECT() *MockLoanMockRecorder {
	return m.recorder
}

// ApproveLoan mocks base method.
func (m *MockLoan) ApproveLoan(ctx context.Context, data entity.Loan, loanID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveLoan", ctx, data, loanID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveLoan indicates an expected call of ApproveLoan.
func (mr *MockLoanMockRecorder) ApproveLoan(ctx, data, loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveLoan", reflect.TypeOf((*MockLoan)(nil).ApproveLoan), ctx, data, loanID)
}

// DisburseLoan mocks base method.
func (m *MockLoan) DisburseLoan(ctx context.Context, data entity.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisburseLoan", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisburseLoan indicates an expected call of DisburseLoan.
func (mr *MockLoanMockRecorder) DisburseLoan(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisburseLoan", reflect.TypeOf((*MockLoan)(nil).DisburseLoan), ctx, data)
}

// GetAllLoan mocks base method.
func (m *MockLoan) GetAllLoan(ctx context.Context, status string) ([]*entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLoan", ctx, status)
	ret0, _ := ret[0].([]*entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLoan indicates an expected call of GetAllLoan.
func (mr *MockLoanMockRecorder) GetAllLoan(ctx, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLoan", reflect.TypeOf((*MockLoan)(nil).GetAllLoan), ctx, status)
}

// GetLoanByID mocks base method.
func (m *MockLoan) GetLoanByID(ctx context.Context, loanID int64) (entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanByID", ctx, loanID)
	ret0, _ := ret[0].(entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanByID indicates an expected call of GetLoanByID.
func (mr *MockLoanMockRecorder) GetLoanByID(ctx, loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanByID", reflect.TypeOf((*MockLoan)(nil).GetLoanByID), ctx, loanID)
}

// InvestLoan mocks base method.
func (m *MockLoan) InvestLoan(ctx context.Context, data entity.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvestLoan", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvestLoan indicates an expected call of InvestLoan.
func (mr *MockLoanMockRecorder) InvestLoan(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvestLoan", reflect.TypeOf((*MockLoan)(nil).InvestLoan), ctx, data)
}

// RegisterNewLoan mocks base method.
func (m *MockLoan) RegisterNewLoan(ctx context.Context, data entity.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNewLoan", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterNewLoan indicates an expected call of RegisterNewLoan.
func (mr *MockLoanMockRecorder) RegisterNewLoan(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewLoan", reflect.TypeOf((*MockLoan)(nil).RegisterNewLoan), ctx, data)
}

// MockLoanInvestor is a mock of LoanInvestor interface.
type MockLoanInvestor struct {
	ctrl     *gomock.Controller
	recorder *MockLoanInvestorMockRecorder
}

// MockLoanInvestorMockRecorder is the mock recorder for MockLoanInvestor.
type MockLoanInvestorMockRecorder struct {
	mock *MockLoanInvestor
}

// NewMockLoanInvestor creates a new mock instance.
func NewMockLoanInvestor(ctrl *gomock.Controller) *MockLoanInvestor {
	mock := &MockLoanInvestor{ctrl: ctrl}
	mock.recorder = &MockLoanInvestorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoanInvestor) EXPECT() *MockLoanInvestorMockRecorder {
	return m.recorder
}

// GetTotalAlreadyInvested mocks base method.
func (m *MockLoanInvestor) GetTotalAlreadyInvested(ctx context.Context, loanID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalAlreadyInvested", ctx, loanID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalAlreadyInvested indicates an expected call of GetTotalAlreadyInvested.
func (mr *MockLoanInvestorMockRecorder) GetTotalAlreadyInvested(ctx, loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalAlreadyInvested", reflect.TypeOf((*MockLoanInvestor)(nil).GetTotalAlreadyInvested), ctx, loanID)
}

// SetNewInvestor mocks base method.
func (m *MockLoanInvestor) SetNewInvestor(ctx context.Context, data entity.LoanInvestor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNewInvestor", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNewInvestor indicates an expected call of SetNewInvestor.
func (mr *MockLoanInvestorMockRecorder) SetNewInvestor(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNewInvestor", reflect.TypeOf((*MockLoanInvestor)(nil).SetNewInvestor), ctx, data)
}