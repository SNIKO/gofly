package main

import (
	"github.com/sniko/gofly/agents"
	"fmt"
	"github.com/sniko/gofly/api/yahoo"
	"sort"
	"time"
)

const (
	DesiredCurrency = "USD"
)

type Flights map[string]DepartureTimes

type DepartureTimes map[time.Time]Planes

type Planes map[string]int

var (
	exchangeRatesCache = make(map[string]float64)
	flightsRefFiles = make(Flights)
)

func processor(fares <-chan agents.Fares, processedFares chan<- agents.Fares) {
	defer close(processedFares)

	for fares := range fares {
		fmt.Printf("Processing %d fares.\n", len(fares))

		updateRefFiles(fares)
		updateFromRefFiles(fares)

		convertPrices(fares, DesiredCurrency)

		sort.Sort(agents.FaresByPrice(fares))
		processedFares <- fares
	}

	fmt.Printf("All fares have been processed.\n")
}

func updateRefFiles(fares agents.Fares) {
	for _, fare := range fares {
		for _, itinerary := range fare.Itineraries {
			for _, flight := range itinerary.Flights {
				// Plane
				if (len(flight.Planes) > 0) {
					flightNumber := flight.CompleteFlightNumber()
					date := flight.DepartureTime.UTC()

					addPlanesToRefFiles(flightNumber, date, flight.Planes)
				}
			}
		}
	}
}

func updateFromRefFiles(fares agents.Fares) {
	for fi := range fares {
		fare := &fares[fi]
		for ii := range fare.Itineraries {
			itinerary := &fare.Itineraries[ii]
			for fli := range itinerary.Flights {
				flight := &itinerary.Flights[fli]

				if len(flight.Planes) == 0 {
					number := flight.CompleteFlightNumber()
					date := flight.DepartureTime.UTC()

					planes, ok := getPlanesFromRefFiles(number, date)
					if (ok) {
						flight.Planes = planes;
					}
				}
			}
		}
	}
}

func addPlanesToRefFiles(flightNumber string, departureDate time.Time, planeType []string) {
	dates, ok := flightsRefFiles[flightNumber]
	if (!ok) {
		dates = make(DepartureTimes)
		flightsRefFiles[flightNumber] = dates
	}

	planes, ok := dates[departureDate]
	if (!ok) {
		planes = make(map[string]int)
		dates[departureDate] = planes
	}

	for _, plane := range planeType {
		planes[plane]++
	}
}

func getPlanesFromRefFiles(flightNumber string, date time.Time) (planes []string, found bool) {
	dates, ok := flightsRefFiles[flightNumber]
	if (!ok) {
		return nil, ok
	}

	planesMap, ok := dates[date]
	if (!ok) {
		// No planes history for the desired date have been found.
		// Collecting the planes statistics from all dates.
		planesMap = make(Planes)
		for date := range dates {
			for plane, count := range dates[date] {
				planesMap[plane] += count
			}
		}
	}

	for plane := range planesMap {
		planes = append(planes, plane)
	}

	return planes, true
}

func convertPrices(fares agents.Fares, currency string) {
	currencies := getCurrencies(fares)
	rates := getExchangeRates(currencies, DesiredCurrency)

	for i, _ := range fares {
		convertedPrices := []agents.PriceInfo{}

		for _, priceInfo := range fares[i].Prices {
			if (priceInfo.Currency == DesiredCurrency) {
				convertedPrices = append(convertedPrices, priceInfo)
				continue
			}

			rate := rates[priceInfo.Currency]
			if (rate > 0) {
				usdPrice := priceInfo.Price * rate
				usdPriceInfo := agents.PriceInfo{usdPrice, DesiredCurrency, priceInfo.Agent, priceInfo.Link}

				convertedPrices = append(convertedPrices, priceInfo)
				convertedPrices = append(convertedPrices, usdPriceInfo)
			}
		}

		fares[i].Prices = convertedPrices
	}
}

func getCurrencies(fares agents.Fares) []string {
	cur := map[string]struct{}{}
	for i, _ := range fares {
		for _, price := range fares[i].Prices {
			cur[price.Currency] = struct{}{}
		}
	}

	c := []string{}
	for key := range cur {
		c = append(c, key)
	}

	return c
}

func getExchangeRates(sourceCurrencies []string, desiredCurrency string) map[string]float64 {
	rates := make(map[string]float64)
	missingPairs := yahoo.CurrencyPairs{}

	cacheKey := func(currency string) string {
		return fmt.Sprintf("%s%s", currency, desiredCurrency)
	}

	for _, currency := range sourceCurrencies {
		if (currency == desiredCurrency) {
			continue
		}

		rate := exchangeRatesCache[cacheKey(currency)]
		if (rate > 0) {
			rates[currency] = rate
		} else {
			missingPairs = append(missingPairs, yahoo.CurrencyPair{currency, desiredCurrency})
		}
	}

	if (len(missingPairs) == 0) {
		return rates
	}

	newRates, err := yahoo.GetExchangeRates(missingPairs);
	if (err != nil) {
		fmt.Printf("yahoo: An error occurred when loading rates for pairs '%s': %s", missingPairs, err)
		return rates
	}

	for _, rate := range newRates {
		rates[rate.From()] = rate.Rate
		exchangeRatesCache[cacheKey(rate.From())] = rate.Rate
	}

	return rates
}