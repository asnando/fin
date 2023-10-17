package models

import "fin/helpers"

type CSVFormat struct {
}

func (cf CSVFormat) FromTransactionList(txs []Transaction) [][]string {
	var rows [][]string
	for _, tx := range txs {
		var line []string
		line = append(line,
			tx.Description,
			helpers.ToBR(tx.Value),
			tx.CostCenter,
			helpers.FromInt(tx.Day),
			tx.Month,
			tx.PaymentMethod,
			helpers.FromBool(tx.Ignore),
			tx.Operation,
		)
		rows = append(rows, line)
	}
	return rows
}

func (cf CSVFormat) ToTransactionList(rows [][]string) []Transaction {
	var txs []Transaction

	for _, l := range rows[1:] {
		txs = append(txs, Transaction{
			Description:   l[0],
			Value:         helpers.FromBR(l[1]),
			CostCenter:    l[2],
			Day:           helpers.ToInt(l[3]),
			Month:         l[4],
			PaymentMethod: l[5],
			Ignore:        helpers.ToBool(l[6]),
			Operation:     l[7],
		})
	}
	return txs
}
