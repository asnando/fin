package models

import (
	"database/sql"
	"encoding/json"
	"fin/common"
)

type Transaction struct {
	Id            int
	Description   string
	Operation     string
	Value         float32
	CostCenter    string
	Day           int
	Month         string
	PaymentMethod string
	Ignore        bool
}

func (tx Transaction) ParseTransactions(rows *sql.Rows) []Transaction {
	defer rows.Close()

	var result []Transaction

	for rows.Next() {
		var id int
		var description string
		var operation string
		var value float32
		var costCenter string
		var day int
		var month string
		var paymentMethod string
		var ignore bool

		err := rows.Scan(&id, &description, &operation, &value, &costCenter, &day, &month, &paymentMethod, &ignore)

		common.Error{}.CheckErr(err)

		result = append(result, Transaction{
			Id:            id,
			Description:   description,
			Operation:     operation,
			Value:         value,
			CostCenter:    costCenter,
			Day:           day,
			Month:         month,
			PaymentMethod: paymentMethod,
			Ignore:        ignore,
		})
	}

	return result
}

func (tx Transaction) ToString() string {
	p, err := json.Marshal(tx)
	common.Error{}.CheckErr(err)
	return string(p)
}
