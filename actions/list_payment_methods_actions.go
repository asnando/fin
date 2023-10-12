package actions

import (
	"fin/common"
	"fin/repositories"
	"fmt"
)

type ListPaymetMethodsAction struct {
	Prompt     common.Prompt
	Repository repositories.PaymentMethodRepository
}

func (lpma ListPaymetMethodsAction) GetActionName() string {
	return "list-payment-methods"
}

func (lpma ListPaymetMethodsAction) Execute() {
	cc := lpma.Repository.List()
	for _, ccr := range cc {
		fmt.Println(ccr.Id, "->", ccr.Description)
	}
}
