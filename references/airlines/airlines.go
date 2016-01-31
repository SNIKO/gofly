package airlines

import (
	"encoding/csv"
	"strings"
	"errors"
	"fmt"
)

type Airline struct {
	Name     string
	Country  string
	IATACode string
	ICAOCode string
}

var airlines map[string]Airline

func GetByIATACode(code string) (Airline, error) {
	var airline Airline

	if (airlines == nil) {
		a, err := LoadAirlines()
		if (err != nil) {
			return airline, err
		}

		airlines = a
	}

	airline, found := airlines[code]
	if (found) {
		return airline, nil
	} else {
		return airline, errors.New(fmt.Sprintf("There are no airline with IATA code '%s'", code))
	}
}

func LoadAirlines() (map[string]Airline, error) {
	reader := csv.NewReader(strings.NewReader(csvdb))
	reader.Comma = ','

	records, err := reader.ReadAll()
	if (err != nil) {
		return nil, err
	}

	airlines := map[string]Airline{}
	for _, record := range records {
		airline := Airline{
			Name: record[1],
			IATACode: record[3],
			ICAOCode: record[4],
			Country: record[6],
		}

		airlines[airline.IATACode] = airline
	}

	return airlines, nil
}