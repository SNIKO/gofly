package main

import (
	"github.com/sniko/gofly/agents"
	"fmt"
	"github.com/sniko/gofly/api/yahoo"
	"sort"
)

var exchangeRatesCache = make(map[string]float64)

func processor(fares <-chan agents.Fares, processedFares chan<- agents.Fares) {
	defer close(processedFares)

	for fares := range fares {
		fmt.Printf("Processing %d fares.\n", len(fares))

		convertPrices(fares, "USD")
		sort.Sort(agents.FaresByPrice(fares))

		processedFares <- fares
	}

	fmt.Printf("All fares have been processed.\n")
}

func convertPrices(fares agents.Fares, currency string) {
	currencies := getCurrencies(fares)
	rates := getExchangeRates(currencies, "USD")

	for i, _ := range fares {
		convertedPrices := []agents.PriceInfo{}

		for _, priceInfo := range fares[i].Prices {
			if (priceInfo.Currency == "USD") {
				convertedPrices = append(convertedPrices, priceInfo)
				continue
			}

			rate := rates[priceInfo.Currency]
			if (rate > 0) {
				usdPrice := priceInfo.Price * rate
				usdPriceInfo := agents.PriceInfo{usdPrice, "USD", priceInfo.Agent, priceInfo.Link}

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