package app

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"loan-engine.bivala.com/repository"
)

func Test_RegisterNewLoanApp_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := repository.NewMockLoan(ctrl)
	fmt.Println(c)
}
