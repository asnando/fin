package actions

import (
	"fin/common"
	"fin/helpers"
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

	for _, l := range csvData[1:] {
		tx := models.Transaction{
			Description:   l[0],
			Value:         helpers.ToValue(l[1]),
			CostCenter:    l[2],
			Day:           helpers.ToInt(l[3]),
			Month:         l[4],
			PaymentMethod: l[5],
			Ignore:        helpers.ToBool(l[6]),
			Operation:     l[7],
		}
		ica.TransactionRepository.Save(tx)
	}
}
