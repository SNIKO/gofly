package agents

import (
	"time"
	"strings"
	"bytes"
	"fmt"
	"github.com/sniko/gofly/references/airlines"
	"github.com/sniko/gofly/references/airports"
)

type FlightDirection struct {
	From string
	To   string
	Date time.Time
}

type Fare struct {
	Itineraries []Itinerary
	Date        time.Time
	Prices      Prices
}

type Fares []Fare

type Prices []PriceInfo

type Itinerary struct {
	Flights []Flight
}

type Flight struct {
	DepartureTime    time.Time
	ArrivalTime      time.Time
	FromAirport      string
	ToAirport        string
	Airline          string
	FlightNumber     string
	OperatedBy       string
	Plane            string
	ReservationClass string
	CabinClass       string
}

type PriceInfo struct {
	Price    float64
	Currency string
	Agent    string
}

func (fare Fare) PrettyString() string {
	var result bytes.Buffer
	departureAirports := []string{}
	departureDates := []string{}

	for _, itinerary := range fare.Itineraries {
		departureAirports = append(departureAirports, itinerary.Flights[0].FromAirport)
		departureDates = append(departureDates, itinerary.Flights[0].DepartureTime.Format("02 Jan"))
	}

	route := strings.Join(departureAirports, "-")
	dates := strings.Join(departureDates, ", ")

	result.WriteString(fmt.Sprintf("%s\t%s\tPrice: %s\n", route, dates, fare.Prices))

	for i, itinerary := range fare.Itineraries {
		result.WriteString(fmt.Sprintf("\tFlight %d\n", i))

		for j, flight := range itinerary.Flights {
			var airlineName string
			var cityName string

			airline, err := airlines.GetByIATACode(flight.Airline)
			if (err != nil) {
				airlineName = err.Error()
			} else {
				airlineName = airline.Name
			}

			airport, err := airports.GetByIATACode(flight.ToAirport)
			if (err != nil) {
				cityName = flight.ToAirport
			} else {
				cityName = airport.City
			}

			if j < len(itinerary.Flights) - 1 {
				thisFlight := itinerary.Flights[j]
				nextFlight := itinerary.Flights[j+1]
				stopOver := nextFlight.DepartureTime.Sub(thisFlight.ArrivalTime)

				result.WriteString(fmt.Sprintf("\t\t%s %4s %15s %s, %s, stopover %s\n", flight.Airline, flight.FlightNumber, cityName, flight.DepartureTime, airlineName, stopOver))
			} else {
				result.WriteString(fmt.Sprintf("\t\t%s %4s %15s %s %s\n", flight.Airline, flight.FlightNumber, cityName, flight.DepartureTime, airlineName))
			}
		}
	}

	return result.String()
}

func (price PriceInfo) String() string {
	return fmt.Sprintf("%s %.2f %s", price.Currency, price.Price, price.Agent)
}

func (d FlightDirection) String() string {
	return fmt.Sprintf("%s -> %s %s", d.From, d.To, d.Date.Format("02 Jan"))
}

type Agent interface {
	Search(directions []FlightDirection) (Fares, error)
}

type FaresByPrice []Fare

func (a FaresByPrice) Len() int {
	return len(a)
}

func (a FaresByPrice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a FaresByPrice) Less(i, j int) bool {
	for ii, _ := range a[i].Prices {
		for jj, _ := range a[j].Prices {
			if (a[i].Prices[ii].Currency == a[j].Prices[jj].Currency) {
				return a[i].Prices[ii].Price < a[j].Prices[jj].Price
			}
		}
	}

	return false
}

func (fares *Fares) Filter(predicate func (Fare) bool) Fares {
	result := Fares {}

	if (fares != nil && predicate != nil) {
		for _, r := range *fares {
			if (predicate(r)) {
				result = append(result, r)
			}
		}
	}

	return result
}

func (prices *Prices) IsLessThan(price PriceInfo) bool {
	for _, p := range *prices {
		if (p.Currency == price.Currency) {
			return p.Price < price.Price
		}
	}

	return false
}