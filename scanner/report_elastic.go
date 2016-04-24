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
        "itineraries": {
          "properties": {
            "flights": {
              "properties": {
                "airline": {
                  "type": "string",
                  "index": "not_analyzed"
                },
                "departure_date": {
                  "type": "date",
                  "format": "strict_date_optional_time||epoch_millis"
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
                "to_city": {
                  "type": "string",
                  "index": "not_analyzed"
                },
                "stopover_at_destination_in_hours": {
                  "type": "long"
                }
              }
            }
          }
        },
        "price_usd": {
          "type": "long"
        },
        "provider": {
          "type": "string",
          "index": "not_analyzed"
        },
        "search_date": {
          "type": "date",
          "format": "strict_date_optional_time||epoch_millis"
        },
        "search_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
        "from_city": {
          "type": "string",
          "index": "not_analyzed"
        },
        "to_city": {
          "type": "string",
          "index": "not_analyzed"
        },
         "departure_date": {
          "type": "date",
          "format": "strict_date_optional_time||epoch_millis"
        },
        "departure_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        },
         "return_departure_date": {
          "type": "date",
          "format": "strict_date_optional_time||epoch_millis"
        },
        "return_departure_day_of_the_week": {
          "type": "string",
          "index": "not_analyzed"
        }
      }
    }
  }
}`
)

type ElasticFare struct {
	Itineraries                 []ElasticItinerary `json:"itineraries"`
	SearchDate                  time.Time 	`json:"search_date"`
	SearchDateOfTheWeek         string 	`json:"search_day_of_the_week"`
	DepartureDate               time.Time 	`json:"departure_date"`
	DepartureDayOfTheWeek       string 	`json:"departure_day_of_the_week"`
	ReturnDepartureDate         time.Time 	`json:"return_departure_date"`
	ReturnDepartureDayOfTheWeek string 	`json:"return_departure_day_of_the_week"`
	FromCity                    string 	`json:"from_city"`
	ToCity                      string 	`json:"to_city"`
	PriceInUSD                  int 	`json:"price_usd"`
	Provider                    string      `json:"provider"`
}

type ElasticItinerary struct {
	Flights []ElasticFlight `json:"flights"`
}

type ElasticFlight struct {
	FromCity                     string 	`json:"from_city"`
	ToCity                       string 	`json:"to_city"`
	DepartureDate                time.Time 	`json:"departure_date"`
	Airline                      string 	`json:"airline"`
	FlightNumber                 string 	`json:"flight_number"`
	Plane                        string 	`json:"plane"`
	StopOverAtDestinationInHours int 	`json:"stopover_at_destination_in_hours"`
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

	elasticFares := ConvertToElasticFares(fares)
	reportedCount := 0

	for _, fare := range elasticFares {
		_, err := es.Index(index, "fare", fare)
		if (err != nil) {
			fmt.Printf("An error occurred when indexing fare %s: %s\n", fare, err)
		} else {
			reportedCount++
		}
	}

	fmt.Printf("%d out of %d fares have been successfully added to elastic search\n", reportedCount, len(elasticFares))
}

func ConvertToElasticFares(fares []agents.Fare) []ElasticFare {
	convertedFares := []ElasticFare{}
	for _, fare := range fares {
		for _, priceInfo := range fare.Prices {
			itineraries := []ElasticItinerary{}
			for _, itinerary := range fare.Itineraries {
				flights := []ElasticFlight{}
				for fli, flight := range itinerary.Flights {
					var stopOver int
					if (fli + 1 < len(itinerary.Flights)) {
						nextFlight := itinerary.Flights[fli + 1]
						stopOver = int(nextFlight.DepartureTime.Sub(flight.ArrivalTime).Hours())
					} else {
						stopOver = 0
					}

					origin := GetCityName(flight.FromAirport)
					destination := GetCityName(flight.ToAirport)
					airline := GetAirlineName(flight.Airline)
					flightNumber := flight.Airline + flight.FlightNumber

					flights = append(flights, ElasticFlight{origin, destination, flight.DepartureTime, airline, flightNumber, flight.Plane, stopOver })
				}

				itineraries = append(itineraries, ElasticItinerary{flights})
			}

			convertedFare := ElasticFare {
				Itineraries: itineraries,
				SearchDate: fare.Date,
				SearchDateOfTheWeek: fare.Date.Weekday().String(),
				PriceInUSD: int(priceInfo.Price),
				Provider: priceInfo.Agent,
			}

			convertedFare.DepartureDate = itineraries[0].Flights[0].DepartureDate
			convertedFare.DepartureDayOfTheWeek = convertedFare.DepartureDate.Weekday().String()
			convertedFare.FromCity = itineraries[0].Flights[0].FromCity

			if (len(itineraries) > 1) {
				convertedFare.ReturnDepartureDate = itineraries[1].Flights[0].DepartureDate
				convertedFare.ReturnDepartureDayOfTheWeek = convertedFare.ReturnDepartureDate.Weekday().String()
				convertedFare.ToCity = itineraries[1].Flights[0].FromCity
			}

			convertedFares = append(convertedFares, convertedFare)
		}
	}

	return convertedFares
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