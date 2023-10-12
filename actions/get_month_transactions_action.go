package actions

import (
	"fin/repositories"
	"fmt"
)

type GetMonthTransactionsAction struct {
	Repository repositories.TransactionRepository
}

func (gmta GetMonthTransactionsAction) GetActionName() string {
	return "get-month-transactions"
}

func (gmta GetMonthTransactionsAction) Execute() {
	txs := gmta.Repository.GetCurrentMonthTransactions()
	for _, tx := range txs {
		fmt.Println(tx.ToString())
	}
}
