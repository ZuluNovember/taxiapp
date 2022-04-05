package csv

import (
	"strconv"

	"github.com/ZuluNovember/taxiapp/driver/models"
	"github.com/ZuluNovember/taxiapp/driver/pkg/file"
)

// Import data from Coordinates.csv and write to mongo db.
func CsvCoordinatesToMongo() error {
	records, err := file.ImportCsv("service/csv/Coordinates.csv")
	if err != nil {
		return err
	}
	fSlice := make([][]float64, len(records))
	for i, v := range records {
		f, err := stringTofloatSlice(v)
		if err != nil {
			return err
		}
		fSlice[i] = f
	}
	_, err = models.BulkAddDriver(fSlice)
	if err != nil {
		return err
	}

	return nil
}

func stringTofloatSlice(s []string) ([]float64, error) {
	fSlice := make([]float64, len(s))

	for i, v := range s {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		fSlice[i] = f
	}

	return fSlice, nil

}
