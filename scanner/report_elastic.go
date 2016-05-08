package main

import (
	"time"
	"github.com/sniko/gofly/agents"
	"github.com/sniko/gofly/references/airlines"
	"github.com/sniko/gofly/references/airports"
	"github.com/sniko/gofly/api/elasticsearch"
	"fmt"
	"strconv"
)

const (
	BatchSize = 1000

	IndexConfig = `
{
  "mappings": {
    "trip": {
      "properties": {
      	"search_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "trip_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_return_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "trip_return_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_origin_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_origin_coordinates": {
          "type": "geo_point",
          "lat_lon": true
        },
        "trip_destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_destination_coordinates": {
          "type": "geo_point",
          "lat_lon": true
        },
        "trip_provider": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_main_airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_summary": {
          "type": "string"
        },
        "trip_price_usd": {
          "type": "long"
        },
        "trip_provider_link": {
          "type": "string",
          "index": "not_analyzed"
        }
      }
    },
    "flight": {
      "properties": {
      	"search_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_departure_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "trip_origin_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_origin_coordinates": {
          "type": "geo_point",
          "lat_lon": true
        },
        "trip_destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_destination_coordinates": {
          "type": "geo_point",
          "lat_lon": true
        },
        "trip_main_airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_summary": {
          "type": "string"
        },
        "trip_price_usd": {
          "type": "long"
        },
        "trip_provider_link": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_departure_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "flight_destination_arrival_time": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "flight_destination_departure_time": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "flight_origin_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_origin_coordinates": {
          "type": "geo_point",
          "lat_lon": true
        },
        "flight_destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_destination_coordinates": {
          "type": "geo_point",
          "lat_lon": true
        },
        "flight_airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_number": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_plane": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_key": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_stopover": {
          "type": "long"
        }
      }
    }
  }
}`
)

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type ElasticTrip struct {
	SearchDate             time.Time   `json:"search_date"`
	SearchDateOfTheWeek    string      `json:"search_day_of_the_week"`
	Date                   time.Time   `json:"trip_date"`
	DayOfTheWeek           string      `json:"trip_day_of_the_week"`
	ReturnDate             time.Time   `json:"trip_return_date"`
	ReturnDayOfTheWeek     string      `json:"trip_return_day_of_the_week"`
	OriginCity             string      `json:"trip_origin_city"`
	OriginCoordinates      *Location   `json:"trip_origin_coordinates"`
	DestinationCity        string      `json:"trip_destination_city"`
	DestinationCoordinates *Location   `json:"trip_destination_coordinates"`
	Provider               string      `json:"trip_provider"`
	MainAirline            string      `json:"trip_main_airline"`
	Summary                string      `json:"trip_summary"`
	PriceInUSD             int         `json:"trip_price_usd"`
	ProviderLink           string 	   `json:"trip_provider_link"`
}

type ElasticFlight struct {
	SearchDate                 time.Time  `json:"search_date"`
	SearchDateOfTheWeek        string     `json:"search_day_of_the_week"`
	TripDepartureDate          time.Time  `json:"trip_departure_date"`
	TripOriginCity             string     `json:"trip_origin_city"`
	TripOriginCoordinates      *Location  `json:"trip_origin_coordinates"`
	TripDestinationCity        string     `json:"trip_destination_city"`
	TripDestinationCoordinates *Location  `json:"trip_destination_coordinates"`
	TripMainAirline            string     `json:"trip_main_airline"`
	TripSummary                string     `json:"trip_summary"`
	TripPriceInUSD             int        `json:"trip_price_usd"`
	TripProviderLink           string     `json:"trip_provider_link"`
	DepartureDate              time.Time  `json:"flight_departure_date"`
	DestinationArrivalTime     time.Time  `json:"flight_destination_arrival_time"`
	DestinationDepartureTime   time.Time  `json:"flight_destination_departure_time"`
	OriginCity                 string     `json:"flight_origin_city"`
	OriginCoordinates          *Location  `json:"flight_origin_coordinates"`
	DestinationCity            string     `json:"flight_destination_city"`
	DestinationCoordinates     *Location  `json:"flight_destination_coordinates"`
	Airline                    string     `json:"flight_airline"`
	FlightNumber               string     `json:"flight_number"`
	Plane                      string     `json:"flight_plane"`
	Key                        string     `json:"flight_key"`
	StopOver                   int        `json:"flight_stopover"`
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
	BulkIndex(es, index, "trip", elasticFares)
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

func GetElasticDocs(fares []agents.Fare) (elasticTrips []interface{}, elasticFlights []interface{}) {
	for _, fare := range fares {
		for _, priceInfo := range fare.Prices {
			if (priceInfo.Currency != "USD") {
				continue
			}

			elasticTrip := CreateElasticFare(&fare, &priceInfo)
			elasticTrips = append(elasticTrips, *elasticTrip)

			for _, itinerary := range fare.Itineraries {
				for i, flight := range itinerary.Flights {
					var nextFlight *agents.Flight
					if (i + 1 < len(itinerary.Flights)) {
						nextFlight = &itinerary.Flights[i + 1]
					}

					elasticFlight := CreateElasticFlight(elasticTrip, &flight, nextFlight)
					elasticFlights = append(elasticFlights, *elasticFlight)
				}
			}
		}
	}

	return elasticTrips, elasticFlights
}

func CreateElasticFlight(trip *ElasticTrip, flight *agents.Flight, nextFlight *agents.Flight) *ElasticFlight {
	// The unique key of the flight: departure date + airline IATA code + flight number
	key := flight.DepartureTime.Format("0102") + flight.Airline + flight.FlightNumber

	elasticFlight := ElasticFlight{
		Key:                    	key,
		SearchDate:             	trip.SearchDate,
		DepartureDate:          	flight.DepartureTime,
		DestinationArrivalTime: 	flight.ArrivalTime,
		SearchDateOfTheWeek:    	trip.SearchDateOfTheWeek,
		Airline:                	GetAirlineName(flight.Airline),
		FlightNumber:           	flight.Airline + flight.FlightNumber,
		Plane:                  	flight.Plane,
		TripSummary:               	trip.Summary,
		TripPriceInUSD:             	trip.PriceInUSD,
		TripProviderLink:		trip.ProviderLink,
		TripDepartureDate:        	trip.Date,
		TripOriginCity:        		trip.OriginCity,
		TripOriginCoordinates:      	trip.OriginCoordinates,
		TripDestinationCity:        	trip.DestinationCity,
		TripDestinationCoordinates: 	trip.DestinationCoordinates,
		TripMainAirline: 		trip.MainAirline,
	}

	if (nextFlight != nil) {
		stopOver := int(nextFlight.DepartureTime.Sub(flight.ArrivalTime).Hours())

		elasticFlight.StopOver = stopOver
		elasticFlight.DestinationDepartureTime = nextFlight.DepartureTime
	}

	city, coordinates := GetAirportInfo(flight.FromAirport)
	elasticFlight.OriginCity = city
	elasticFlight.OriginCoordinates = coordinates

	city, coordinates = GetAirportInfo(flight.ToAirport)
	elasticFlight.DestinationCity = city
	elasticFlight.DestinationCoordinates = coordinates

	return &elasticFlight
}

func CreateElasticFare(fare *agents.Fare, priceInfo *agents.PriceInfo) *ElasticTrip {
	var trip, returnTrip *agents.Itinerary

	trip = &fare.Itineraries[0]
	if (len(fare.Itineraries) > 1) {
		returnTrip = &fare.Itineraries[1]
	}

	elasticFare := ElasticTrip{
		SearchDate: 		fare.Date,
		SearchDateOfTheWeek: 	fare.Date.Weekday().String(),
		PriceInUSD: 		int(priceInfo.Price),
		ProviderLink: 		priceInfo.Link,
		Provider: 		priceInfo.Agent,
		Summary: 		fare.PrettyString(),
		Date:			trip.Flights[0].DepartureTime,
		DayOfTheWeek:		trip.Flights[0].DepartureTime.Weekday().String(),
	}

	// Origin
	city, coordinates := GetAirportInfo(trip.Flights[0].FromAirport)
	elasticFare.OriginCity = city
	elasticFare.OriginCoordinates = coordinates

	// Destination
	var destinationAirport string
	if (returnTrip != nil) {
		destinationAirport = returnTrip.Flights[0].FromAirport
	} else {
		destinationAirport = trip.Flights[len(trip.Flights) - 1].ToAirport
	}

	city, coordinates = GetAirportInfo(destinationAirport)
	elasticFare.DestinationCity = city
	elasticFare.DestinationCoordinates = coordinates

	// Return trip info
	if (returnTrip != nil) {
		elasticFare.ReturnDate = returnTrip.Flights[0].DepartureTime
		elasticFare.ReturnDayOfTheWeek = elasticFare.ReturnDate.Weekday().String()
	}

	// Main Airline
	mainAirline := ""
	flightDuration := map[string]float64{}
	for _, itinerary := range fare.Itineraries {
		for _, flight := range itinerary.Flights {
			airline := GetAirlineName(flight.Airline)
			duration := flight.ArrivalTime.Sub(flight.DepartureTime).Hours()

			flightDuration[airline] += duration

			if (mainAirline != airline && flightDuration[airline] > flightDuration[mainAirline]) {
				mainAirline = airline
			}
		}
	}

	elasticFare.MainAirline = mainAirline

	return &elasticFare
}

func GetAirportInfo(airportIataCode string) (cityName string, coordinates *Location) {
	airport, err := airports.GetByIATACode(airportIataCode)
	if (err != nil) {
		return airportIataCode, nil
	} else {
		lat, err := strconv.ParseFloat(airport.Latitude, 64)
		if (err != nil) {
			return airportIataCode, nil
		}

		lon, err := strconv.ParseFloat(airport.Longitude, 64)
		if (err != nil) {
			return airportIataCode, nil
		}

		return airport.City, &Location{lat, lon}
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