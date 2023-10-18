package csv

import (
	"encoding/csv"
	"os"

	"github.com/jszwec/csvutil"
)

type Data struct {
	WalletAddress string `csv:"address"`
	Amount        int    `csv:"amount"`
}

func ReadCsv(path string) ([]Data, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []Data
	if err := csvutil.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func WriteCsv(path string, value [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewWriter(file)
	err = r.WriteAll(value)

	return err
}
