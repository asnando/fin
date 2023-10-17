package actions

import (
	"fin/common"
	"fin/models"
	"fin/repositories"
)

type ImportCSVAction struct {
	TransactionRepository repositories.TransactionRepository
}

func (ica ImportCSVAction) GetActionName() string {
	return "import-csv"
}

func (ica ImportCSVAction) Execute() {
	reader := common.CSVReader{}

	csvData := reader.Read("transactions.csv")

	txs := models.CSVFormat{}.ToTransactionList(csvData[1:])

	for _, tx := range txs {
		ica.TransactionRepository.Save(tx)
	}
}
