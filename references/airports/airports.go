package airports

import (
	"encoding/csv"
	"strings"
	"errors"
	"fmt"
)

type Airport struct {
	Name     string
	City     string
	Country  string
	IATACode string
	ICAOCode string
	Timezone string
}

var airports map[string]Airport

func GetByIATACode(code string) (Airport, error) {
	var airport Airport

	if (airports == nil) {
		a, err := LoadAirports()
		if (err != nil) {
			return airport, err
		}

		airports = a
	}

	airport, found := airports[code]
	if (found) {
		return airport, nil
	} else {
		return airport, errors.New(fmt.Sprintf("There are no airport with IATA code '%s'", code))
	}
}

func LoadAirports() (map[string]Airport, error) {
	reader := csv.NewReader(strings.NewReader(csvdb))
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if (err != nil) {
		return nil, err
	}

	airports := map[string]Airport{}
	for _, record := range records {
		airport := Airport{
			Name: record[1],
			City: record[2],
			Country: record[3],
			IATACode: record[4],
			ICAOCode: record[5],
			Timezone: record[11],
		}

		airports[airport.IATACode] = airport
	}

	return airports, nil
}