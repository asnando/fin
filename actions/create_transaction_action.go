package actions

import (
	"fin/common"
	"fin/helpers"
	"fin/models"
	"fin/repositories"
	"fmt"
)

type CreateTransactionAction struct {
	Prompt                  common.Prompt
	TransactionRepository   repositories.TransactionRepository
	CostCenterRepository    repositories.CostCenterRepository
	PaymentMethodRepository repositories.PaymentMethodRepository
}

func (cta CreateTransactionAction) GetActionName() string {
	return "create-transaction"
}

func displayCostCenters(ccs []models.CostCenter) {
	for _, cc := range ccs {
		fmt.Println(cc.Id-1, "->", cc.Description)
	}
}

func displayPaymentMethods(pms []models.PaymentMethod) {
	for _, pm := range pms {
		fmt.Println(pm.Id-1, "->", pm.Description)
	}
}

func (cta CreateTransactionAction) pickCostCenter() string {
	costCenters := cta.CostCenterRepository.List()
	displayCostCenters(costCenters)
	selectedIndex := cta.Prompt.ReadInt("Pick a cost center")
	return costCenters[selectedIndex].Description
}

func (cta CreateTransactionAction) pickPaymentMethod() string {
	paymentMethods := cta.PaymentMethodRepository.List()
	displayPaymentMethods(paymentMethods)
	selectedIndex := cta.Prompt.ReadInt("Pick a payment method")
	return paymentMethods[selectedIndex].Description
}

func (cta CreateTransactionAction) pickMonth() string {
	monthPath := helpers.GetMonthMap()
	monthIndex := cta.Prompt.ReadInt(fmt.Sprint(monthPath))
	return monthPath[monthIndex]
}

func (cta CreateTransactionAction) Execute() {
	description := cta.Prompt.ReadInputWithMessage("Enter description:")
	value := cta.Prompt.ReadFloat("Enter value:")
	costCenter := cta.pickCostCenter()
	day := cta.Prompt.ReadInt("Enter day of month:")
	month := cta.pickMonth()
	paymentMethod := cta.pickPaymentMethod()
	ignore := cta.Prompt.ReadInputForOptions("Should ignore?", []string{"y", "n"}) == "y"

	fmt.Println(description, value, costCenter, day, month, paymentMethod, ignore)

	tx := models.Transaction{
		Description:   description,
		Operation:     models.OUT,
		Value:         value,
		CostCenter:    costCenter,
		Day:           day,
		Month:         month,
		PaymentMethod: paymentMethod,
		Ignore:        ignore,
	}

	cta.TransactionRepository.Save(tx)
}
