package read

import (
	"encoding/csv"
	"os"
)

func ReadCsv(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}
