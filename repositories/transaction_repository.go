package repositories

import (
	"fin/common"
	"fin/helpers"
	"fin/models"
	"fmt"
)

type TransactionRepository struct {
	Database common.Database
}

func (tr TransactionRepository) Save(tx models.Transaction) {
	tr.Database.Insert(
		"transactions",
		"description,operation,value,cost_center,day,month,payment_method,ignore",
		tx.Description,
		tx.Operation,
		tx.Value,
		tx.CostCenter,
		tx.Day,
		tx.Month,
		tx.PaymentMethod,
		tx.Ignore,
	)
}

func (tr TransactionRepository) List() []models.Transaction {
	rows := tr.Database.Query("transactions")
	return models.Transaction{}.ParseTransactions(rows)
}

func (tr TransactionRepository) GetLastTransaction() models.Transaction {
	rows := tr.Database.RawQuery("SELECT * FROM transactions ORDER BY id DESC LIMIT 1")
	return models.Transaction{}.ParseTransactions(rows)[0]
}

func (tr TransactionRepository) GetCurrentMonthTransactions() []models.Transaction {
	rows := tr.Database.RawQuery(fmt.Sprintf("SELECT * FROM transactions WHERE month='%s' ORDER BY id", helpers.GetCurrentMonth()))
	return models.Transaction{}.ParseTransactions(rows)
}
