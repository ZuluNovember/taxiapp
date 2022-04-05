package file

import (
	"encoding/csv"
	"os"
)

//ImportCsv takes a path string and returns all the records in the csv file.
func ImportCsv(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	_, err = csvReader.Read() //Skip column headers
	if err != nil {
		return nil, err
	}

	recs, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return recs, nil
}
