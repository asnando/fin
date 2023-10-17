package common

import (
	"encoding/csv"
	"os"
)

type CSVWriter struct {
}

func (csvw CSVWriter) Write(file string, data [][]string) {
	f, err := os.Create(file)

	Error{}.CheckErr(err)

	defer f.Close()

	w := csv.NewWriter(f)

	err = w.WriteAll(data)

	Error{}.CheckErr(err)
}
