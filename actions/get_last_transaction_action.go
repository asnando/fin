package actions

import (
	"fin/repositories"
	"fmt"
)

type GetLastTransactionAction struct {
	Repository repositories.TransactionRepository
}

func (glta GetLastTransactionAction) Execute() {
	tx := glta.Repository.GetLastTransaction()
	fmt.Println(tx.ToString())
}

func (glta GetLastTransactionAction) GetActionName() string {
	return "get-last-transaction"
}
