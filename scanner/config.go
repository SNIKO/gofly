package main
import (
	"time"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Trips []Trip
	ElasticSearch *ElasticSearchConfig
}

type ElasticSearchConfig struct {
	Host  string
	Port  int
}

type Trip struct {
	FromAirport   []string
	ToAirport     []string
	MinDate       ConfigDate
	MaxDate       ConfigDate
	MinReturnDate ConfigDate
	MaxReturnDate ConfigDate
	MinDuration   int
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

func loadConfig(path string) (*Config, error) {
	configFile, err := ioutil.ReadFile(path)
	if (err != nil) {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if (err != nil) {
		return nil, err
	}

	return &config, err
}