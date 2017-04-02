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
				fromDate := tripConfig.MinDate.time
				toDate := tripConfig.MaxDate.time.AddDate(0, 0, 1)

				for date := fromDate; date.Before(toDate); date = date.AddDate(0, 0, 1) {
					flight := agents.FlightDirection{origin, destination, date}

					if tripConfig.MinReturnDate.time.IsZero() == false {
						fromReturnDate := tripConfig.MinReturnDate.time
						toReturnDate := tripConfig.MaxReturnDate.time.AddDate(0, 0, 1)

						returnDate := date.AddDate(0, 0, tripConfig.MinDuration - 1)
						if returnDate.Before(fromReturnDate) {
							returnDate = fromReturnDate
						}

						for ; returnDate.Before(toReturnDate); returnDate = returnDate.AddDate(0, 0, 1) {
							returnFlight := agents.FlightDirection{destination, origin, returnDate}
							trips = append(trips, TripInfo{Route: []agents.FlightDirection{flight, returnFlight}})
						}
					} else {
						trips = append(trips, TripInfo{Route: []agents.FlightDirection{flight}})
					}
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