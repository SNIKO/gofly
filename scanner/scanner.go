package main

import (
	"fmt"
	"github.com/sniko/gofly/agents"
	"strings"
	"sync"
)

type TripInfo struct {
	Route []agents.FlightDirection
}

func main() {
	config, err := loadConfig("config.json")
	if (err != nil) {
		fmt.Printf("Failed to load config: %s", err)
		return
	}

	tripsCh := make(chan TripInfo)
	faresCh := make(chan agents.Fares)
	processedFaresCh := make(chan agents.Fares)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go tripsGenerator(config, tripsCh)
	go fetcher(tripsCh, faresCh)
	go processor(faresCh, processedFaresCh)
	go reporter(processedFaresCh, *config, &wg)

	wg.Wait()
}

func tripsGenerator(config *Config, trips chan<- TripInfo) {
	combinations := getTripCombinations(config.Trips)
	for _, trip := range combinations {
		trips <- trip
	}

	close(trips)
}

func getTripCombinations(tripsConfigs []Trip) (trips []TripInfo) {
	for _, tripConfig := range tripsConfigs {
		for _, origin := range tripConfig.FromAirport {
			for _, destination := range tripConfig.ToAirport {
				date := tripConfig.MinDate.time

				for date.Before(tripConfig.MaxDate.time.AddDate(0, 0, 1)) {
					duration := tripConfig.MinDuration

					for duration <= tripConfig.MaxDuration {
						flight := agents.FlightDirection{origin, destination, date}
						returnFlight := agents.FlightDirection{destination, origin, date.AddDate(0, 0, duration)}

						trips = append(trips, TripInfo{[]agents.FlightDirection{flight, returnFlight}})
						duration++
					}

					date = date.AddDate(0, 0, 1)
				}
			}
		}
	}

	return
}

func (t TripInfo) String() string {
	directions := []string{}
	for _, d := range t.Route {
		directions = append(directions, d.String())
	}

	return strings.Join(directions, ", ")
}

func (t TripInfo) ShortString() string {
	airports := []string{}
	for _, d := range t.Route {
		airports = append(airports, d.From)
	}

	return fmt.Sprintf("%s %s", t.Route[0].Date.Format("02 Jan"), strings.Join(airports, "-"))
}