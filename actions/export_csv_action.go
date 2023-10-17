package actions

import (
	"fin/common"
	"fin/models"
	"fin/repositories"
)

type ExportCSVAction struct {
	TransactionRepository repositories.TransactionRepository
}

func (eca ExportCSVAction) GetActionName() string {
	return "export-csv"
}

func (eca ExportCSVAction) Execute() {
	txs := eca.TransactionRepository.List()
	data := models.CSVFormat{}.FromTransactionList(txs)

	w := common.CSVWriter{}

	w.Write("transactions-out.csv", data)
}
