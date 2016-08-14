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
	LocalDateFormat = "Mon, 02 Jan 2006 15:04"
	IndexConfig = `
{
  "mappings": {
    "trip": {
      "properties": {
      	"search_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_return_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_return_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_return_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_arrival_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_arrival_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_return_arrival_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_return_arrival_date_local": {
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
        "trip_via_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_provider": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_main_airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_airline": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_plane": {
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
          "format": "date_time_no_millis",
          "type": "date"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_arrival_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_arrival_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_return_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_return_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "trip_return_arrival_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "trip_return_arrival_date_local": {
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
          "format": "date_time_no_millis",
          "type": "date"
        },
        "flight_departure_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_arrival_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "flight_arrival_date_local": {
          "type": "string",
          "index": "not_analyzed"
        },
        "flight_next_flight_departure_date": {
          "format": "date_time_no_millis",
          "type": "date"
        },
        "flight_next_flight_departure_date_local": {
          "type": "string",
          "index": "not_analyzed"
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

type ElasticTime time.Time

type ElasticTrip struct {
	SearchDate             ElasticTime      `json:"search_date"`
	SearchDateOfTheWeek    string        	`json:"search_day_of_the_week"`
	Date                   ElasticTime      `json:"trip_date"`
	DateLocal              string           `json:"trip_date_local"`
	DayOfTheWeek           string        	`json:"trip_day_of_the_week"`
	ReturnDate             ElasticTime      `json:"trip_return_date"`
	ReturnDateLocal        string           `json:"trip_return_date_local"`
	ReturnDayOfTheWeek     string        	`json:"trip_return_day_of_the_week"`
	ArrivalDate            ElasticTime      `json:"trip_arrival_date"`
	ArrivalDateLocal       string           `json:"trip_arrival_date_local"`
	ReturnArrivalDate      ElasticTime      `json:"trip_return_arrival_date"`
	ReturnArrivalDateLocal string           `json:"trip_return_arrival_date_local"`
	OriginCity             string        	`json:"trip_origin_city"`
	OriginCoordinates      *Location        `json:"trip_origin_coordinates"`
	DestinationCity        string        	`json:"trip_destination_city"`
	DestinationCoordinates *Location        `json:"trip_destination_coordinates"`
	ViaCities 	       []string 	`json:"trip_via_city"`
	Provider               string        	`json:"trip_provider"`
	MainAirline            string        	`json:"trip_main_airline"`
	Airlines               []string 	`json:"trip_airline"`
	Planes                 []string        	`json:"trip_plane"`
	Summary                string        	`json:"trip_summary"`
	PriceInUSD             int              `json:"trip_price_usd"`
	ProviderLink           string        	`json:"trip_provider_link"`
}

type ElasticFlight struct {
	SearchDate                   ElasticTime  	`json:"search_date"`
	SearchDateOfTheWeek          string     	`json:"search_day_of_the_week"`
	TripDate                     ElasticTime  	`json:"trip_date"`
	TripDateLocal                string 		`json:"trip_date_local"`
	TripArrivalDate              ElasticTime  	`json:"trip_arrival_date"`
	TripArrivalDateLocal         string 		`json:"trip_arrival_date_local"`
	TripReturnDate               ElasticTime  	`json:"trip_return_date"`
	TripReturnDateLocal          string 		`json:"trip_return_date_local"`
	TripReturnArrivalDate        ElasticTime  	`json:"trip_return_arrival_date"`
	TripReturnArrivalDateLocal   string 		`json:"trip_return_arrival_date_local"`
	TripOriginCity               string     	`json:"trip_origin_city"`
	TripOriginCoordinates        *Location  	`json:"trip_origin_coordinates"`
	TripDestinationCity          string     	`json:"trip_destination_city"`
	TripDestinationCoordinates   *Location  	`json:"trip_destination_coordinates"`
	TripMainAirline              string     	`json:"trip_main_airline"`
	TripSummary                  string     	`json:"trip_summary"`
	TripPriceInUSD               int        	`json:"trip_price_usd"`
	TripProviderLink             string     	`json:"trip_provider_link"`
	DepartureDate                ElasticTime  	`json:"flight_departure_date"`
	DepartureDateLocal           string 		`json:"flight_departure_date_local"`
	ArrivalDate                  ElasticTime  	`json:"flight_arrival_date"`
	ArrivalDateLocal             string 		`json:"flight_arrival_date_local"`
	NextFlightDepartureDate      ElasticTime  	`json:"flight_next_flight_departure_date"`
	NextFlightDepartureDateLocal string  		`json:"flight_next_flight_departure_date_local"`
	OriginCity                   string     	`json:"flight_origin_city"`
	OriginCoordinates            *Location  	`json:"flight_origin_coordinates"`
	DestinationCity              string     	`json:"flight_destination_city"`
	DestinationCoordinates       *Location  	`json:"flight_destination_coordinates"`
	Airline                      string     	`json:"flight_airline"`
	FlightNumber                 string     	`json:"flight_number"`
	Planes                       []string   	`json:"flight_plane"`
	Key                          string     	`json:"flight_key"`
	StopOver                     int        	`json:"flight_stopover"`
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
	Index(es, index, "trip", elasticFares)
	Index(es, index, "flight", elasticFlights)
	elapsed := time.Since(start)

	fmt.Printf("elastic: Indexing completed. Elapsed time: %s\n", elapsed)
}

func Index(client *elasticsearch.Client, index string, documentType string, documents []interface{}) {
	batch := []elasticsearch.BulkCommand {}

	for _, doc := range documents {
		batch = append(batch, elasticsearch.IndexCommand{Index: index, DocumentType: documentType, Content: doc})

		if (len(batch) >= BatchSize) {
			BulkIndex(client, documentType, batch)
			batch = []elasticsearch.BulkCommand{}
		}
	}

	if (len(batch) > 0) {
		BulkIndex(client, documentType, batch)
	}
}

func BulkIndex(client *elasticsearch.Client, documentType string, batch []elasticsearch.BulkCommand) {
	res, err := client.Bulk(batch)
	if (err != nil) {
		fmt.Printf("%s elastic: an error occurred when indexing a %s batch of size %d: %s\n", time.Now().String(), documentType, len(batch), err)
	} else {
		if (res.Errors) {
			fmt.Printf("%s elastic: an error occurred when indexing a %s batch of size %d\n", time.Now().String(), documentType, len(batch))
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
		DepartureDate:          	ElasticTime(flight.DepartureTime),
		DepartureDateLocal: 		flight.DepartureTime.Format(LocalDateFormat),
		ArrivalDate: 			ElasticTime(flight.ArrivalTime),
		ArrivalDateLocal: 		flight.ArrivalTime.Format(LocalDateFormat),
		SearchDateOfTheWeek:    	trip.SearchDateOfTheWeek,
		Airline:                	GetAirlineName(flight.Airline),
		FlightNumber:           	flight.Airline + flight.FlightNumber,
		Planes:                  	flight.Planes,
		TripSummary:               	trip.Summary,
		TripPriceInUSD:             	trip.PriceInUSD,
		TripProviderLink:		trip.ProviderLink,
		TripDate:        		trip.Date,
		TripDateLocal: 			trip.DateLocal,
		TripReturnDate: 		trip.ReturnDate,
		TripReturnDateLocal: 		trip.ReturnDateLocal,
		TripArrivalDate: 		trip.ArrivalDate,
		TripArrivalDateLocal: 		trip.ArrivalDateLocal,
		TripReturnArrivalDate: 		trip.ReturnArrivalDate,
		TripReturnArrivalDateLocal: 	trip.ReturnArrivalDateLocal,
		TripOriginCity:        		trip.OriginCity,
		TripOriginCoordinates:      	trip.OriginCoordinates,
		TripDestinationCity:        	trip.DestinationCity,
		TripDestinationCoordinates: 	trip.DestinationCoordinates,
		TripMainAirline: 		trip.MainAirline,
	}

	if (nextFlight != nil) {
		stopOver := int(nextFlight.DepartureTime.Sub(flight.ArrivalTime).Hours())

		elasticFlight.StopOver = stopOver
		elasticFlight.NextFlightDepartureDate = ElasticTime(nextFlight.DepartureTime)
		elasticFlight.NextFlightDepartureDateLocal = nextFlight.DepartureTime.Format(LocalDateFormat)
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

	departureDate := trip.Flights[0].DepartureTime
	arrivalDate := trip.Flights[len(trip.Flights) - 1].ArrivalTime

	elasticFare := ElasticTrip{
		SearchDate: 		ElasticTime(fare.Date),
		SearchDateOfTheWeek: 	fare.Date.Weekday().String(),
		PriceInUSD: 		int(priceInfo.Price),
		ProviderLink: 		priceInfo.Link,
		Provider: 		priceInfo.Agent,
		Summary: 		fare.PrettyString(),
		Date:			ElasticTime(departureDate),
		DateLocal: 		departureDate.Format(LocalDateFormat),
		DayOfTheWeek:		departureDate.Weekday().String(),
		ArrivalDate: 		ElasticTime(arrivalDate),
		ArrivalDateLocal: 	arrivalDate.Format(LocalDateFormat),
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
		depDate := returnTrip.Flights[0].DepartureTime
		arvDate :=  returnTrip.Flights[len(returnTrip.Flights) - 1].ArrivalTime

		elasticFare.ReturnDate = ElasticTime(depDate)
		elasticFare.ReturnDateLocal = depDate.Format(LocalDateFormat)
		elasticFare.ReturnDayOfTheWeek = depDate.Weekday().String()
		elasticFare.ReturnArrivalDate = ElasticTime(arvDate)
		elasticFare.ReturnArrivalDateLocal = arvDate.Format(LocalDateFormat)
	}

	// Airlines, planes, via cities
	mainAirline := ""
	planes := map[string]struct{}{}
	via := map[string]int{}
	flightDuration := map[string]float64{}
	for _, itinerary := range fare.Itineraries {
		for _, flight := range itinerary.Flights {
			airline := GetAirlineName(flight.Airline)
			city := GetCityName(flight.ToAirport)
			duration := flight.ArrivalTime.Sub(flight.DepartureTime).Hours()

			for _, plane := range flight.Planes {
				planes[plane] = struct{}{}
			}

			flightDuration[airline] += duration

			if (mainAirline != airline && flightDuration[airline] > flightDuration[mainAirline]) {
				mainAirline = airline
			}

			if (city != elasticFare.OriginCity && city != elasticFare.DestinationCity) {
				via[city] += 1
			}
		}
	}

	for key := range(planes) {
		elasticFare.Planes = append(elasticFare.Planes, key)
	}

	for airline := range(flightDuration) {
		elasticFare.Airlines = append(elasticFare.Airlines, airline)
	}

	for city := range(via) {
		elasticFare.ViaCities = append(elasticFare.ViaCities, city)
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

func GetCityName(airportCode string) string {
	airport, err := airports.GetByIATACode(airportCode)
	if (err != nil) {
		return airportCode
	} else {
		return airport.City
	}
}

func (t ElasticTime) MarshalJSON() ([]byte, error) {
	date := time.Time(t).Format(fmt.Sprintf("\"%s\"", time.RFC3339))
	return []byte(date), nil
}