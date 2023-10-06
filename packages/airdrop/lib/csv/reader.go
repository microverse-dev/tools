package csv

import (
	"encoding/csv"
	"os"
)

func ReadCsv() ([]string, error) {
	file, err := os.Open("list.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var wallets []string
	for _, v := range rows {
		wallets = append(wallets, v[0])
	}

	// remove header
	return wallets[1:], nil
}
