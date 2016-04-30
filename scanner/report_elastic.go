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
	BatchSize = 1000

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
        "origin_city": {
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
        "destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
     	"origin_coordinates": {
          "type": "geo_point"
        },
        "destination_coordinates": {
          "type": "geo_point"
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
        "origin_city": {
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
        "stopover": {
          "type": "long"
        },
        "destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "origin_coordinates": {
          "type": "geo_point"
        },
        "destination_coordinates": {
          "type": "geo_point"
        }
      }
    }
  }
}`
)

type ElasticFare struct {
	SearchDate             time.Time   `json:"search_date"`
	Date                   time.Time   `json:"date"`
	ReturnDate             time.Time   `json:"return_date"`
	SearchDateOfTheWeek    string      `json:"search_day_of_the_week"`
	DayOfTheWeek           string      `json:"day_of_the_week"`
	ReturnDayOfTheWeek     string      `json:"return_day_of_the_week"`
	OriginCity             string      `json:"origin_city"`
	OriginCoordinates      string      `json:"origin_coordinates"`
	DestinationCity        string      `json:"destination_city"`
	DestinationCoordinates string      `json:"destination_coordinates"`
	Provider               string      `json:"provider"`
	FareInfo               string      `json:"fare_info"`
	PriceInUSD             int         `json:"price_usd"`
}

type ElasticFlight struct {
	SearchDate             time.Time  `json:"search_date"`
	DepartureDate          time.Time  `json:"departure_date"`
	SearchDateOfTheWeek    string     `json:"search_day_of_the_week"`
	OriginCity             string     `json:"origin_city"`
	OriginCoordinates      string     `json:"origin_coordinates"`
	DestinationCity        string     `json:"destination_city"`
	DestinationCoordinates string     `json:"destination_coordinates"`
	Airline                string     `json:"airline"`
	FlightNumber           string     `json:"flight_number"`
	Plane                  string     `json:"plane"`
	FareInfo               string     `json:"fare_info"`
	Key                    string     `json:"flight_key"`
	StopOver               int        `json:"stopover"`
	PriceInUSD             int        `json:"price_usd"`
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

	fmt.Printf("elastic: Indexing %d fares and %d flights...\n", len(elasticFares), len(elasticFlights))

	start := time.Now()
	BulkIndex(es, index, "fare", elasticFares)
	BulkIndex(es, index, "flight", elasticFlights)
	elapsed := time.Since(start)

	fmt.Printf("elastic: Indexing completed. Elapsed time: %s\n", elapsed)
}

func BulkIndex(client *elasticsearch.Client, index string, documentType string, documents []interface{}) {
	batch := []elasticsearch.BulkCommand {}

	for _, doc := range documents {
		batch = append(batch, elasticsearch.IndexCommand{Index: index, DocumentType: documentType, Content: doc})

		if (len(batch) >= BatchSize) {
			res, err := client.Bulk(batch)
			if (err != nil) {
				fmt.Printf("%s elastic: an error occurred when indexing a %s batch of size %d: %s\n", time.Now().String(), documentType, len(batch), err)
			} else {
				if (res.Errors) {
					fmt.Printf("%s elastic: an error occurred when indexing a %s batch of size %d\n", time.Now().String(), documentType, len(batch))
				}
			}

			batch = []elasticsearch.BulkCommand{}
		}
	}
}

func GetElasticDocs(fares []agents.Fare) (elasticFares []interface{}, elasticFlights []interface{}) {
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
		Airline:                GetAirlineName(flight.Airline),
		FlightNumber:           flight.Airline + flight.FlightNumber,
		Plane:                  flight.Plane,
		FareInfo:               fare.PrettyString(),
		PriceInUSD:             int(priceInfo.Price),
	}

	if (nextFlight != nil) {
		stopOver := int(nextFlight.DepartureTime.Sub(flight.ArrivalTime).Hours())
		elasticFlight.StopOver = stopOver
	}

	city, coordinates := GetAirportInfo(flight.FromAirport)
	elasticFlight.OriginCity = city
	elasticFlight.OriginCoordinates = coordinates

	city, coordinates = GetAirportInfo(flight.ToAirport)
	elasticFlight.DestinationCity = city
	elasticFlight.DestinationCoordinates = coordinates

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
	}

	city, coordinates := GetAirportInfo(trip.Flights[0].FromAirport)
	elasticFare.OriginCity = city
	elasticFare.OriginCoordinates = coordinates

	var destinationAirport string
	if (returnTrip != nil) {
		destinationAirport = returnTrip.Flights[0].FromAirport
	} else {
		destinationAirport = trip.Flights[len(trip.Flights) - 1].ToAirport
	}

	city, coordinates = GetAirportInfo(destinationAirport)
	elasticFare.DestinationCity = city
	elasticFare.DestinationCoordinates = coordinates

	if (returnTrip != nil) {
		elasticFare.ReturnDate = returnTrip.Flights[0].DepartureTime
		elasticFare.ReturnDayOfTheWeek = elasticFare.ReturnDate.Weekday().String()
	}

	return &elasticFare
}

func GetAirportInfo(airportIataCode string) (cityName string, coordinates string) {
	airport, err := airports.GetByIATACode(airportIataCode)
	if (err != nil) {
		return "", airportIataCode
	} else {
		return airport.City, fmt.Sprintf("%s,%s", airport.Latitude, airport.Longitude)
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