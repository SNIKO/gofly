package main

import (
	"time"
	"github.com/sniko/gofly/agents"
	"github.com/sniko/gofly/references/airlines"
	"github.com/sniko/gofly/references/airports"
	"github.com/sniko/gofly/api/elasticsearch"
	"fmt"
)

const (
	IndexConfig = `
{
  "mappings": {
    "fare": {
      "properties": {
        "date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "fare_info": {
          "type": "string"
        },
        "from_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "price_usd": {
          "type": "long"
        },
        "provider": {
          "type": "string",
          "index": "not_analyzed"
        },
        "return_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "return_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "search_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "to_city": {
          "type": "string",
          "index": "not_analyzed"
        }
      }
    },
    "flight": {
      "properties": {
        "airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "departure_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "fare_info": {
          "type": "string"
        },
        "flight_key": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_number": {
          "type": "string",
          "index": "not_analyzed"
        },
        "from_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "plane": {
          "type": "string",
          "index": "not_analyzed"
        },
        "price_usd": {
          "type": "long"
        },
        "search_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "stopover_at_destination_in_hours": {
          "type": "long"
        },
        "to_city": {
          "type": "string",
          "index": "not_analyzed"
        }
      }
    }
  }
}`
)

type ElasticFare struct {
	SearchDate          time.Time   `json:"search_date"`
	Date                time.Time   `json:"date"`
	ReturnDate          time.Time   `json:"return_date"`
	SearchDateOfTheWeek string      `json:"search_day_of_the_week"`
	DayOfTheWeek        string      `json:"day_of_the_week"`
	ReturnDayOfTheWeek  string      `json:"return_day_of_the_week"`
	FromCity            string      `json:"from_city"`
	ToCity              string      `json:"to_city"`
	Provider            string      `json:"provider"`
	FareInfo            string	`json:"fare_info"`
	PriceInUSD          int        	`json:"price_usd"`
}

type ElasticFlight struct {
	SearchDate                   time.Time  `json:"search_date"`
	DepartureDate                time.Time  `json:"departure_date"`
	SearchDateOfTheWeek          string     `json:"search_day_of_the_week"`
	FromCity                     string     `json:"from_city"`
	ToCity                       string     `json:"to_city"`
	Airline                      string     `json:"airline"`
	FlightNumber                 string     `json:"flight_number"`
	Plane                        string     `json:"plane"`
	FareInfo                     string     `json:"fare_info"`
	Key			     string	`json:"flight_key"`
	StopOverAtDestinationInHours int        `json:"stopover_at_destination_in_hours"`
	PriceInUSD                   int        `json:"price_usd"`
}

func ReportElasticSearch(host string, port int, index string, fares []agents.Fare) {
	es := elasticsearch.NewClient(host, port)

	exists, err := es.IndexExists(index)
	if (err != nil) {
		fmt.Printf("An error occurred when checking the index '%s': %s\n", index, err)
		return
	}

	if (!exists) {
		fmt.Printf("Index '%s' doesn't exist. Creating it...\n", index)
		res, err := es.CreateIndex(index, IndexConfig)
		if (err != nil) {
			fmt.Printf("An error occurred when creating index '%s': %s\n", index, err)
			return
		} else {
			if (!res.Acknowledged) {
				fmt.Printf("Can't create index '%s'\n", index)
				return
			}
		}
	}

	elasticFares, elasticFlights := GetElasticDocs(fares)
	indexedFares, indexedFlights := 0, 0

	for _, fare := range elasticFares {
		_, err := es.Index(index, "fare", fare)
		if (err != nil) {
			fmt.Printf("An error occurred when indexing fare %s: %s\n", fare, err)
		} else {
			indexedFares++
		}
	}

	fmt.Printf("elastic: %d out of %d fares have been successfully indexed\n", indexedFares, len(elasticFares))

	for _, flight := range elasticFlights {
		_, err := es.Index(index, "flight", flight)
		if (err != nil) {
			fmt.Printf("An error occurred when indexing flight %s: %s\n", flight, err)
		} else {
			indexedFlights++
		}
	}

	fmt.Printf("elastic: %d out of %d flights have been successfully indexed\n", indexedFlights, len(elasticFlights))
}

func GetElasticDocs(fares []agents.Fare) ([]ElasticFare, []ElasticFlight) {
	elasticFares := []ElasticFare{}
	elasticFlights := []ElasticFlight{}

	for _, fare := range fares {
		for _, priceInfo := range fare.Prices {
			if (priceInfo.Currency != "USD") {
				continue
			}

			for _, itinerary := range fare.Itineraries {
				for i, flight := range itinerary.Flights {
					var nextFlight *agents.Flight
					if (i + 1 < len(itinerary.Flights)) {
						nextFlight = &itinerary.Flights[i + 1]
					}

					elasticFlight := CreateElasticFlight(&fare, &priceInfo, &flight, nextFlight)
					elasticFlights = append(elasticFlights, *elasticFlight)
				}
			}

			elasticFare := CreateElasticFare(&fare, &priceInfo)
			elasticFares = append(elasticFares, *elasticFare)
		}
	}

	return elasticFares, elasticFlights
}

func CreateElasticFlight(fare *agents.Fare, priceInfo *agents.PriceInfo, flight *agents.Flight, nextFlight *agents.Flight) *ElasticFlight {
	// The unique key of the flight: departure date + airline IATA code + flight number
	key := flight.DepartureTime.Format("0102") + flight.Airline + flight.FlightNumber

	elasticFlight := ElasticFlight{
		Key:                    key,
		SearchDate:             fare.Date,
		DepartureDate:          flight.DepartureTime,
		SearchDateOfTheWeek:    fare.Date.Weekday().String(),
		FromCity:               GetCityName(flight.FromAirport),
		ToCity:                 GetCityName(flight.ToAirport),
		Airline:                GetAirlineName(flight.Airline),
		FlightNumber:           flight.Airline + flight.FlightNumber,
		Plane:                  flight.Plane,
		FareInfo:               fare.PrettyString(),
		PriceInUSD:             int(priceInfo.Price),
	}

	if (nextFlight != nil) {
		stopOver := int(nextFlight.DepartureTime.Sub(flight.ArrivalTime).Hours())
		elasticFlight.StopOverAtDestinationInHours = stopOver
	}

	return &elasticFlight
}

func CreateElasticFare(fare *agents.Fare, priceInfo *agents.PriceInfo) *ElasticFare {
	var trip, returnTrip *agents.Itinerary

	trip = &fare.Itineraries[0]
	if (len(fare.Itineraries) > 1) {
		returnTrip = &fare.Itineraries[1]
	}

	elasticFare := ElasticFare{
		SearchDate: 		fare.Date,
		SearchDateOfTheWeek: 	fare.Date.Weekday().String(),
		PriceInUSD: 		int(priceInfo.Price),
		Provider: 		priceInfo.Agent,
		FareInfo: 		fare.PrettyString(),
		Date:			trip.Flights[0].DepartureTime,
		DayOfTheWeek:		trip.Flights[0].DepartureTime.Weekday().String(),
		FromCity:		GetCityName(trip.Flights[0].FromAirport),
	}

	if (returnTrip != nil) {
		elasticFare.ReturnDate = returnTrip.Flights[0].DepartureTime
		elasticFare.ReturnDayOfTheWeek = elasticFare.ReturnDate.Weekday().String()
		elasticFare.ToCity = GetCityName(returnTrip.Flights[0].FromAirport)
	} else {
		elasticFare.ToCity = GetCityName(trip.Flights[len(trip.Flights) - 1].ToAirport)
	}

	return &elasticFare
}

func GetCityName(airportIataCode string) string {
	airport, err := airports.GetByIATACode(airportIataCode)
	if (err != nil) {
		return airportIataCode
	} else {
		return airport.City
	}
}

func GetAirlineName(iataCode string) string  {
	airline, err := airlines.GetByIATACode(iataCode)
	if (err != nil) {
		return iataCode
	} else {
		return airline.Name
	}
}