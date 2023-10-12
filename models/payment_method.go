package models

import (
	"database/sql"
	"fin/common"
)

type PaymentMethod struct {
	Id          int
	Description string
}

func (pm PaymentMethod) ParsePaymentMethods(rows *sql.Rows) []PaymentMethod {
	defer rows.Close()

	var result []PaymentMethod

	for rows.Next() {
		var id int
		var description string

		err := rows.Scan(&id, &description)

		common.Error{}.CheckErr(err)

		result = append(result, PaymentMethod{
			Id:          id,
			Description: description,
		})
	}

	return result
}
