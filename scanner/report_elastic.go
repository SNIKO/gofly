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
    "trip": {
      "properties": {
        "trip_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "trip_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "fare_info": {
          "type": "string"
        },
        "trip_origin_city": {
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
        "trip_main_airline": {
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
        "search_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
     	"trip_origin_coordinates": {
          "type": "geo_point"
        },
        "trip_destination_coordinates": {
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
        "flight_departure_date": {
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
        "flight_origin_city": {
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
        "flight_destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_origin_coordinates": {
          "type": "geo_point"
        },
        "flight_destination_coordinates": {
          "type": "geo_point"
        },
        "trip_departure_date": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "trip_origin_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_destination_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_origin_coordinates": {
          "type": "geo_point"
        },
        "trip_destination_coordinates": {
          "type": "geo_point"
        },
        "trip_main_airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_origin_arrival_time": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        },
        "flight_destination_departure_time": {
          "format": "strict_date_optional_time||epoch_millis",
          "type": "date"
        }
      }
    }
  }
}`
)

type ElasticTrip struct {
	SearchDate             time.Time   `json:"search_date"`
	Date                   time.Time   `json:"trip_date"`
	ReturnDate             time.Time   `json:"trip_return_date"`
	SearchDateOfTheWeek    string      `json:"search_day_of_the_week"`
	DayOfTheWeek           string      `json:"trip_day_of_the_week"`
	ReturnDayOfTheWeek     string      `json:"trip_return_day_of_the_week"`
	OriginCity             string      `json:"trip_origin_city"`
	OriginCoordinates      string      `json:"trip_origin_coordinates"`
	DestinationCity        string      `json:"trip_destination_city"`
	DestinationCoordinates string      `json:"trip_destination_coordinates"`
	Provider               string      `json:"provider"`
	MainAirline            string 	   `json:"trip_main_airline"`
	FareInfo               string      `json:"fare_info"`
	PriceInUSD             int         `json:"price_usd"`
}

type ElasticFlight struct {
	SearchDate                 time.Time  `json:"search_date"`
	DepartureDate              time.Time  `json:"flight_departure_date"`
	TripDepartureDate          time.Time  `json:"trip_departure_date"`
	OriginArrivalTime          time.Time  `json:"flight_origin_arrival_time"`
	DestinationDepartureTime   time.Time  `json:"flight_destination_departure_time"`
	SearchDateOfTheWeek        string     `json:"search_day_of_the_week"`
	OriginCity                 string     `json:"flight_origin_city"`
	OriginCoordinates          string     `json:"flight_origin_coordinates"`
	DestinationCity            string     `json:"flight_destination_city"`
	DestinationCoordinates     string     `json:"flight_destination_coordinates"`
	TripOriginCity             string     `json:"trip_origin_city"`
	TripDestinationCity        string     `json:"trip_destination_city"`
	TripOriginCoordinates      string     `json:"trip_origin_coordinates"`
	TripDestinationCoordinates string     `json:"trip_destination_coordinates"`
	TripMainAirline            string     `json:"trip_main_airline"`
	Airline                    string     `json:"airline"`
	FlightNumber               string     `json:"flight_number"`
	Plane                      string     `json:"plane"`
	FareInfo                   string     `json:"fare_info"`
	Key                        string     `json:"flight_key"`
	StopOver                   int        `json:"stopover"`
	PriceInUSD                 int        `json:"price_usd"`
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
		OriginArrivalTime: 		flight.ArrivalTime,
		SearchDateOfTheWeek:    	trip.SearchDateOfTheWeek,
		Airline:                	GetAirlineName(flight.Airline),
		FlightNumber:           	flight.Airline + flight.FlightNumber,
		Plane:                  	flight.Plane,
		FareInfo:               	trip.FareInfo,
		PriceInUSD:             	trip.PriceInUSD,
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
		Provider: 		priceInfo.Agent,
		FareInfo: 		fare.PrettyString(),
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