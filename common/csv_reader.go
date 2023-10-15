package common

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
}

func (csvr CSVReader) Read(file string) [][]string {
	f, err := os.Open(file)

	Error{}.CheckErr(err)

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	Error{}.CheckErr(err)

	return data
}
