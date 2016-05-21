package agents

import (
	"github.com/sniko/gofly/api/onetwotrip"
	"errors"
	"fmt"
	"time"
)

type OneTwoTrip struct {
}

func (client OneTwoTrip) Name() string {
	return "OneTwoTrip"
}

func (c OneTwoTrip) Search(directions []FlightDirection) (Fares, error) {
	query := getSearchQuery(directions)

	result, err := onetwotrip.Search(query)
	if (err != nil) {
		return nil, err
	}

	fares, err := GetFares(result)
	if (err != nil) {
		return nil, err
	}

	return fares, nil
}

func getSearchQuery(directions []FlightDirection) onetwotrip.SearchQuery {
	flights := []onetwotrip.Flight{}

	for _, dir := range directions {
		flight := onetwotrip.Flight{dir.From, dir.To, dir.Date}
		flights = append(flights, flight)
	}

	return onetwotrip.SearchQuery{flights}
}

func GetFares(result *onetwotrip.SearchResult) (Fares, error) {
	fares := []Fare{}

	for _, fare := range result.Fares {
		itineraries, err := getItineraries(result, fare)
		if (err != nil) {
			return nil, err
		}

		price := PriceInfo{fare.Price.TotalPrice(), fare.Price.Currency, "OneTwoTrip", ""}

		fare := Fare{itineraries, time.Now(), Prices{price}}
		fares = append(fares, fare)
	}

	return fares, nil
}

func getItineraries(result *onetwotrip.SearchResult, fare onetwotrip.Fare)  ([]Itinerary, error) {
	itineraries := []Itinerary{}

	for _, dir := range fare.Directions {
		flights, err := getFlights(result, dir)
		if (err != nil) {
			return nil, err
		}

		itinerary := Itinerary{flights}
		itineraries = append(itineraries, itinerary)
	}

	return itineraries, nil
}

func getFlights(result *onetwotrip.SearchResult, dir onetwotrip.Direction) ([]Flight, error) {
	flights := []Flight{}

	for _, tripRef := range dir.Trips {
		flight, err := getFlight(result, &tripRef)
		if (err != nil) {
			return nil, err
		}

		flights = append(flights, *flight)
	}

	return flights, nil
}

func getFlight(result *onetwotrip.SearchResult, tripRef *onetwotrip.TripRef) (*Flight, error) {
	if (tripRef.Id >= len(result.Trips)) {
		msg := fmt.Sprintf("Failed to parse onetwotrip response: Trip with id '%s' not found in the response.", tripRef.Id)
		return nil, errors.New(msg)
	}

	trip := result.Trips[tripRef.Id]

	departureTime, err := trip.DepartureTime()
	if (err != nil) {
		return nil, err
	}

	arrivalTime, err := trip.ArrivalTime()
	if (err != nil) {
		return nil, err
	}

	plane := result.Planes[trip.Plane]

	flight := Flight{
		*departureTime,
		*arrivalTime,
		trip.From,
		trip.To,
		trip.AirlineCode,
		trip.FlightNumber,
		trip.OperatedBy,
		plane,
		tripRef.ReservationClass,
		tripRef.CabinClass}

	return &flight, nil
}