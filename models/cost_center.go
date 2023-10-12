package models

import (
	"database/sql"
	"fin/common"
)

type CostCenter struct {
	Id          int
	Description string
}

func (cc CostCenter) ParseCostCenters(rows *sql.Rows) []CostCenter {
	defer rows.Close()

	var result []CostCenter

	for rows.Next() {
		var id int
		var description string

		err := rows.Scan(&id, &description)

		common.Error{}.CheckErr(err)

		result = append(result, CostCenter{
			Id:          id,
			Description: description,
		})
	}

	return result
}
