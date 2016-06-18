package agents

import (
	"time"
	"github.com/sniko/gofly/api/momondo"
	"strconv"
	"reflect"
)

type Momondo struct {
}

type ReferenceFiles struct {
	Airlines      []momondo.Airline
	Airports      []momondo.Airport
	Flights       []momondo.Flight
	Legs          []momondo.Leg
	Segments      []momondo.Segment
	Offers        []momondo.Offer
	TicketClasses []momondo.TicketClass
}

func (references *ReferenceFiles) Update(searchResult *momondo.SearchResult) {
	if (searchResult.Airlines != nil) {
		references.Airlines = append(references.Airlines, *searchResult.Airlines...)
	}

	if (searchResult.Airports != nil) {
		references.Airports = append(references.Airports, *searchResult.Airports...)
	}

	if (searchResult.Flights != nil) {
		references.Flights = append(references.Flights, *searchResult.Flights...)
	}

	if (searchResult.Legs != nil) {
		references.Legs = append(references.Legs, *searchResult.Legs...)
	}

	if (searchResult.Segments != nil) {
		references.Segments = append(references.Segments, *searchResult.Segments...)
	}

	if (searchResult.Offers != nil) {
		references.Offers = append(references.Offers, *searchResult.Offers...)
	}

	if (searchResult.TicketClasses != nil) {
		references.TicketClasses = append(references.TicketClasses, *searchResult.TicketClasses...)
	}
}

func (client Momondo) Name() string {
	return "Momondo"
}

func (client Momondo) Search(directions []FlightDirection) (Fares, error) {
	searchId, engineId, err := startSearch(directions)
	if err != nil {
		return nil, err
	}

	searchCompleted := false
	references := ReferenceFiles{}
	allFares := []Fare{}

	for searchCompleted == false {
		time.Sleep(3 * time.Second)

		result, err := momondo.PollSearchResults(searchId, engineId)
		if err != nil {
			return nil, err
		}

		references.Update(result)

		fares, err := parseFares(result, &references)
		if (err != nil) {
			return nil, err
		}

		allFares = append(allFares, fares...)
		searchCompleted = result.Done
	}

	return allFares, nil
}

func startSearch(directions []FlightDirection) (string, int, error) {
	segments := make([]momondo.Direction, len(directions))
	for i, direction := range directions {
		segments[i] = momondo.Direction{direction.From, direction.To, momondo.FlightDate(direction.Date)}
	}

	request := momondo.NewRequest(1, "en-US", "ECO", segments)

	info, err := momondo.StartSearch(*request)
	if err != nil {
		return "", 0, err
	}

	return info.SearchId, info.EngineId, nil
}

func parseFares(searchResult *momondo.SearchResult, references *ReferenceFiles) ([]Fare, error) {
	fares := []Fare{}

	if (searchResult.Suppliers == nil) {
		return fares, nil
	}

	for _, supplier := range *searchResult.Suppliers {
		for _, offerIndex := range supplier.OfferIndexes {
			offer := references.Offers[offerIndex]
			flight := references.Flights[offer.FlightIndex]
			ticketClass := references.TicketClasses[offer.TicketClassIndex]
			itineraries := []Itinerary{}

			for _, segmentIndex := range flight.SegmentIndexes {
				segment := references.Segments[segmentIndex]
				flights := []Flight{}

				for _, legIndex := range segment.LegIndexes {
					leg := references.Legs[legIndex]
					origin := references.Airports[leg.OriginIndex]
					destination := references.Airports[leg.DestinationIndex]
					airline := references.Airlines[leg.AirlineIndex]

					flight := Flight{leg.DepartureDate, leg.ArrivalDate, origin.IATACode, destination.IATACode, airline.IATACode, strconv.Itoa(leg.FlightNumber), "", []string {}, ticketClass.Code, ticketClass.Name }
					flights = append(flights, flight)
				}

				itineraries = append(itineraries, Itinerary{flights})
			}

			price := PriceInfo{offer.TotalPrice, offer.Currency, supplier.DisplayName, offer.DeepLink}

			var existingFare *Fare

			for i, _ := range fares {
				if (reflect.DeepEqual(fares[i].Itineraries, itineraries)) {
					existingFare = &fares[i]
					break
				}
			}

			if (existingFare != nil) {
				existingFare.Prices = append(existingFare.Prices, price)
			} else {
				fares = append(fares, Fare{itineraries, time.Now(), []PriceInfo{price}})
			}
		}
	}

	return fares, nil
}