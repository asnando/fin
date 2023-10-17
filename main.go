package main

import (
	"fin/actions"
	"fin/common"
	"fin/repositories"
)

func main() {
	prompt := common.Prompt{}

	a := &App{
		Prompt:  prompt,
		Actions: createActions(prompt),
	}

	a.Start()
}

func createActions(prompt common.Prompt) []actions.Action {
	var actionList []actions.Action
	var database = &common.Database{}

	costCenterRepository := &repositories.CostCenterRepository{
		Database: *database,
	}

	paymentMethodRepository := &repositories.PaymentMethodRepository{
		Database: *database,
	}

	transactionRepository := &repositories.TransactionRepository{
		Database: *database,
	}

	createCostCenterAction := &actions.CreateCostCenterAction{
		Prompt:     prompt,
		Repository: *costCenterRepository,
	}

	listCostCenterAction := &actions.ListCostCenterAction{
		Prompt:     prompt,
		Repository: *costCenterRepository,
	}

	createPaymentMethodAction := &actions.CreatePaymentMethodAction{
		Prompt:     prompt,
		Repository: *paymentMethodRepository,
	}

	listPaymentMethodsAction := &actions.ListPaymetMethodsAction{
		Prompt:     prompt,
		Repository: *paymentMethodRepository,
	}

	createTransactionAction := &actions.CreateTransactionAction{
		Prompt:                  prompt,
		TransactionRepository:   *transactionRepository,
		PaymentMethodRepository: *paymentMethodRepository,
		CostCenterRepository:    *costCenterRepository,
	}

	getLastTransactionAction := &actions.GetLastTransactionAction{
		Repository: *transactionRepository,
	}

	getMonthTransactionAction := &actions.GetMonthTransactionsAction{
		Repository: *transactionRepository,
	}

	importCSVAction := &actions.ImportCSVAction{
		TransactionRepository: *transactionRepository,
	}

	exportCSVAction := &actions.ExportCSVAction{
		TransactionRepository: *transactionRepository,
	}

	actionList = append(
		actionList,
		createCostCenterAction,
		createPaymentMethodAction,
		listCostCenterAction,
		listPaymentMethodsAction,
		createTransactionAction,
		getLastTransactionAction,
		getMonthTransactionAction,
		importCSVAction,
		exportCSVAction,
	)

	return actionList
}
