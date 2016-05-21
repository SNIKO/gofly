package main

import (
	"fmt"
	"github.com/sniko/gofly/api/yahoo"
	"github.com/sniko/gofly/agents"
	"sort"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bufio"
	"os"
	"time"
	"sync"
)

type TripInfo struct {
	Route []agents.FlightDirection
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

func getRates(trip TripInfo) <-chan agents.Fares {
	wg := sync.WaitGroup{}
	out := make(chan agents.Fares)
	providers := []agents.Agent{agents.Momondo{}, agents.OneTwoTrip{}}

	for _, provider := range providers {
		wg.Add(1)

		go func(agent agents.Agent) {
			defer wg.Done()
			fmt.Printf("%s: Searching flights for route '%s'\n", agent.Name(), trip)
			rates, err := agent.Search(trip.Route)
			if (err != nil) {
				fmt.Printf("%s: An error occurred when loading fares for route '%s': %s\n", agent.Name(), trip, err)
				return
			}

			fmt.Printf("%s: %d fares for route '%s' have been loaded\n", agent.Name(), len(rates), trip)
			out <- rates
		}(provider)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	config, err := loadConfig()
	if (err != nil) {
		fmt.Printf("Failed to load config: %s", err)
		return
	}

	trips := getTripCombinations(config.Trips)

	fares := agents.Fares{}

	for _, trip := range trips {
		for tripFares := range getRates(trip) {
			fares = append(fares, tripFares...)
		}
	}

	currencies := yahoo.CurrencyPairs{}
	for i, _ := range fares {
		currency := fares[i].Prices[0].Currency

		if (currency != "USD" && !currencies.Contains(currency, "USD")) {
			currencies = append(currencies, yahoo.CurrencyPair{currency, "USD"})
		}
	}

	rates, err := yahoo.GetExchangeRates(currencies)
	if (err != nil) {
		fmt.Printf("error: %s\n", err)
		return
	}

	for i, _ := range fares {
		convertedPrices := []agents.PriceInfo{}

		for _, priceInfo := range fares[i].Prices {
			if (priceInfo.Currency == "USD") {
				convertedPrices = append(convertedPrices, priceInfo)
				continue
			}

			rate := rates.Find(priceInfo.Currency, "USD")

			if (rate != nil) {
				usdPrice := priceInfo.Price * rate.Rate
				usdPriceInfo := agents.PriceInfo{usdPrice, "USD", priceInfo.Agent, priceInfo.Link}

				convertedPrices = append(convertedPrices, priceInfo)
				convertedPrices = append(convertedPrices, usdPriceInfo)
			}
		}

		fares[i].Prices = convertedPrices
	}

	sort.Sort(agents.FaresByPrice(fares))

	fmt.Printf("%d fares have been loaded\n", len(fares))

	fileName := fmt.Sprintf("%s (on %s).txt", trips[0].ShortString(), time.Now().Format("02 Jan 2006"));
	f, err := os.Create(fmt.Sprintf("fares/%s", fileName))
	if (err != nil) {
		fmt.Printf("An error occurred when creating an output file: %s", err)
	}

	writer := bufio.NewWriter(f)

	for i, _ := range fares {
		writer.WriteString(fares[i].PrettyString())
		writer.WriteString("\n")
	}

	writer.Flush()

	ReportElasticSearch(config.ElasticSearch.Host, config.ElasticSearch.Port, config.ElasticSearch.Index, fares)
}

func loadConfig() (*Config, error) {
	configFile, err := ioutil.ReadFile("config.json")
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

func getTripCombinations(tripsConfigs []Trip) []TripInfo {
	trips := []TripInfo{}

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

	return trips
}