package actions

import (
	"fin/common"
	"fin/repositories"
)

type CreatePaymentMethodAction struct {
	Prompt     common.Prompt
	Repository repositories.PaymentMethodRepository
}

func (cpma CreatePaymentMethodAction) GetActionName() string {
	return "create-payment-method"
}

func (cpma CreatePaymentMethodAction) Execute() {
	input := cpma.Prompt.ReadInputWithMessage("Enter payment method:")
	cpma.Repository.Save(input)
}
