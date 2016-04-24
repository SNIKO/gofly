package main
import (
	"time"
)

type Config struct {
	Trips []Trip
	ElasticSearch *ElasticSearchConfig
}

type ElasticSearchConfig struct {
	Host  string
	Port  int
	Index string
}

type Trip struct {
	FromAirport []string
	ToAirport   []string
	MinDate     ConfigDate
	MaxDate     ConfigDate
	MinDuration int
	MaxDuration int
}

type ConfigDate struct {
	time time.Time
}

func (d *ConfigDate) UnmarshalJSON(b []byte) error {
	if (b[0] == '"' && b[len(b) - 1] == '"') {
		b = b[1 : len(b)-1]
	}

	date, err := time.Parse("02.01.2006", string(b))
	if (err != nil) {
		return err
	}

	d.time = date
	return nil
}