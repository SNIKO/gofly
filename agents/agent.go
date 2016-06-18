package agents

import (
	"time"
	"strings"
	"bytes"
	"fmt"
	"github.com/sniko/gofly/references/airlines"
	"github.com/sniko/gofly/references/airports"
	"text/tabwriter"
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
	Planes           []string
	ReservationClass string
	CabinClass       string
}

type PriceInfo struct {
	Price    float64
	Currency string
	Agent    string
	Link 	 string
}

type Agent interface {
	Name() string
	Search(directions []FlightDirection) (Fares, error)
}

type FaresByPrice []Fare

func (fare Fare) PrettyString() string {
	var result bytes.Buffer
	departureAirports := []string{}
	departureDates := []string{}

	for _, itinerary := range fare.Itineraries {
		departureAirports = append(departureAirports, itinerary.Flights[0].FromAirport)
		departureDates = append(departureDates, itinerary.Flights[0].DepartureTime.Format("02 Jan"))
	}

	route := strings.Join(departureAirports, "-")
	dates := strings.Join(departureDates, " - ")
	result.WriteString(fmt.Sprintf("%s\t%s\tPrice: %s\n", route, dates, fare.Prices))

	flightsWriter := bytes.NewBufferString("")
	tabWriter := new(tabwriter.Writer)
	tabWriter.Init(flightsWriter, 2, 8, 1, '\t', 0)
	write := func(columns... string) {
		fmt.Fprintln(tabWriter, "\t", strings.Join(columns, "\t"))
	}

	for i, itinerary := range fare.Itineraries {
		write(fmt.Sprintf("Flight %d", i), "", "", "", "", "", "", "")

		for j, flight := range itinerary.Flights {
			airline := getAirlineName(flight.Airline)
			destination := getCityName(flight.ToAirport)
			planes := strings.Join(flight.Planes, ", ")

			stopover := ""
			if j < len(itinerary.Flights) - 1 {
				thisFlight := itinerary.Flights[j]
				nextFlight := itinerary.Flights[j + 1]
				stopover = nextFlight.DepartureTime.Sub(thisFlight.ArrivalTime).String()
			}

			write("", flight.Airline, flight.FlightNumber, airline, planes, destination, flight.DepartureTime.String(), stopover)
		}
	}

	tabWriter.Flush()
	result.WriteString(flightsWriter.String())

	return result.String()
}

func getAirlineName(iataCode string) string  {
	airline, err := airlines.GetByIATACode(iataCode)
	if (err != nil) {
		return iataCode
	} else {
		return airline.Name
	}
}

func getCityName(airportCode string) string  {
	airport, err := airports.GetByIATACode(airportCode)
	if (err != nil) {
		return airportCode
	} else {
		return airport.City
	}
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

func (price PriceInfo) String() string {
	return fmt.Sprintf("%s %.2f %s", price.Currency, price.Price, price.Agent)
}

func (prices Prices) String() string {
	p := []string{}

	for _, price := range prices {
		p = append(p, price.String())
	}

	return strings.Join(p, ", ")
}

func (prices *Prices) IsLessThan(price PriceInfo) bool {
	for _, p := range *prices {
		if (p.Currency == price.Currency) {
			return p.Price < price.Price
		}
	}

	return false
}

func (d FlightDirection) String() string {
	return fmt.Sprintf("%s -> %s %s", d.From, d.To, d.Date.Format("02 Jan"))
}

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

func (f *Flight) CompleteFlightNumber() string {
	return fmt.Sprintf("%s%s", f.Airline, f.FlightNumber)
}