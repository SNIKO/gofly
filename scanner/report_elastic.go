package main

import (
	"time"
	"github.com/sniko/gofly/agents"
	"github.com/sniko/gofly/references/airlines"
	"github.com/sniko/gofly/references/airports"
	"github.com/sniko/gofly/api/elasticsearch"
	"fmt"
	"strconv"
	"errors"
)

const (
	BatchSize = 1000
	LocalDateFormat = "Mon, 02 Jan 2006 15:04"
	IndexName = "gofly"
	IndexTemplateName = "gofly"
	IndexTemplate = `
{
  "template": "gofly-*",
  "mappings": {
    "_default_": {
      "_all": {
        "enabled": false
      },
      "dynamic_templates": [
        {
          "strings_as_keywords": {
            "match_mapping_type": "string",
            "unmatch": "summary",
            "mapping": {
              "type": "keyword"
            }
          }
        },
        {
          "geo_fields": {
            "mapping": {
              "type": "geo_point",
              "lat_lon": true
            },
            "match_mapping_type": "object",
            "match": "geo"
          }
        }
      ]
    }
  }
}`
)

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type ElasticTime time.Time

type ElasticDate struct {
	Time         ElasticTime `json:"time"`
	DayOfTheWeek string 	 `json:"day_of_the_week"`
	LocalTime    string 	 `json:"local_time"`
}

type ElasticAirport struct {
	IATACode string   `json:"iata_code"`
	City     string   `json:"city"`
	Geo      Location `json:"geo"`
}

type ElasticFlight struct {
	DepartureDate ElasticDate    `json:"departure_date"`
	ArrivalDate   ElasticDate    `json:"arrival_date"`
	Origin        ElasticAirport `json:"origin"`
	Destination   ElasticAirport `json:"destination"`
	Airline       string         `json:"airline"`
	FlightNumber  string 	     `json:"flight_number"`
	Planes        []string       `json:"plane"`
	StopOver      int            `json:"stopover"`
	Key           string 	     `json:"key"`
}

type ElasticFare struct {
	SearchDate        ElasticDate     `json:"search_date"`
	DepartureDate     ElasticDate 	  `json:"departure_date"`
	ArrivalDate       ElasticDate 	  `json:"arrival_date"`
	ReturnDate        ElasticDate 	  `json:"return_date"`
	ReturnArrivalDate ElasticDate 	  `json:"return_arrival_date"`
	Origin            ElasticAirport  `json:"origin"`
	Destination       ElasticAirport  `json:"destination"`
	Flight1           *ElasticFlight  `json:"flight_1"`
	Flight2           *ElasticFlight  `json:"flight_2"`
	Flight3           *ElasticFlight  `json:"flight_3"`
	ReturnFlight1     *ElasticFlight  `json:"return_flight_1"`
	ReturnFlight2     *ElasticFlight  `json:"return_flight_2"`
	ReturnFlight3     *ElasticFlight  `json:"return_flight_3"`
	ViaCities         []string        `json:"via_city"`
	Airlines          []string        `json:"airline"`
	Planes            []string        `json:"plane"`
	PriceInUSD        int             `json:"price_usd"`
	Provider          string          `json:"provider"`
	ProviderLink      string          `json:"provider_link"`
	Summary           string          `json:"summary"`
}

func ReportElasticsearch(host string, port int, fares []agents.Fare) {
	es := elasticsearch.NewClient(host, port)

	err := ApplyIndexTemplate(es, IndexTemplateName, IndexTemplate)
	if (err != nil) {
		fmt.Printf("An error occurred when creating template %s: %s", IndexTemplateName, err)
		return
	}

	elasticFares := GetElasticFares(fares)

	fmt.Printf("elastic: Indexing %d fares...\n", len(elasticFares))
	start := time.Now()

	timedIndexName := IndexName + "-" + time.Now().Format("2006.01.02")
	Index(es, timedIndexName, "fare", elasticFares)

	elapsed := time.Since(start)
	fmt.Printf("elastic: Indexing completed. Elapsed time: %s\n", elapsed)
}

func ApplyIndexTemplate(es *elasticsearch.Client, templateName string, template string) error {
	exists, err := es.TemplateExists(templateName)
	if (err != nil) {
		return err
	}

	if (exists) {
		res, err := es.DeleteTemplate(templateName)
		if (err != nil) {
			return err
		}

		if (!res.Acknowledged) {
			return errors.New("Can't delete the existing template")
		}
	}

	res, err := es.CreateTemplate(templateName, template)
	if (err != nil) {
		return err;
	}

	if (!res.Acknowledged) {
		return errors.New("Can't create the template")
	}

	return nil
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
			fmt.Println(res.Items)
			fmt.Printf("%s elastic: an error occurred when indexing a %s batch of size %d\n", time.Now().String(), documentType, len(batch))
		}
	}
}

func GetElasticFares(fares []agents.Fare) (elasticFares []interface{}) {
	for _, fare := range fares {
		for _, priceInfo := range fare.Prices {
			if (priceInfo.Currency != "USD") {
				continue
			}

			elasticTrip := CreateElasticFare(&fare, &priceInfo)
			elasticFares = append(elasticFares, *elasticTrip)
		}
	}

	return elasticFares
}

func CreateElasticFare(fare *agents.Fare, priceInfo *agents.PriceInfo) *ElasticFare {
	var trip, returnTrip *agents.Itinerary

	trip = &fare.Itineraries[0]
	if (len(fare.Itineraries) > 1) {
		returnTrip = &fare.Itineraries[1]
	}

	departureDate := trip.Flights[0].DepartureTime
	origin := trip.Flights[0].FromAirport
	arrivalDate := trip.Flights[len(trip.Flights) - 1].ArrivalTime
	destination := trip.Flights[len(trip.Flights) - 1].ToAirport

	elasticFare := ElasticFare{
		SearchDate:    GetElasticDate(fare.Date),
		DepartureDate: GetElasticDate(departureDate),
		ArrivalDate:   GetElasticDate(arrivalDate),
		Origin:        GetElasticAirport(origin),
		Destination:   GetElasticAirport(destination),

		PriceInUSD:    int(priceInfo.Price),
		ProviderLink:  priceInfo.Link,
		Provider:      priceInfo.Agent,
		Summary:       fare.PrettyString(),
	}

	// Flights
	elasticFare.Flight1 = GetElasticFlight(trip.Flights, 1)
	elasticFare.Flight2 = GetElasticFlight(trip.Flights, 2)
	elasticFare.Flight3 = GetElasticFlight(trip.Flights, 3)

	// Return trip info
	if (returnTrip != nil) {
		depDate := returnTrip.Flights[0].DepartureTime
		arvDate := returnTrip.Flights[len(returnTrip.Flights) - 1].ArrivalTime

		elasticFare.ReturnDate = GetElasticDate(depDate)
		elasticFare.ReturnArrivalDate = GetElasticDate(arvDate)

		// Return Flights
		elasticFare.ReturnFlight1 = GetElasticFlight(returnTrip.Flights, 1)
		elasticFare.ReturnFlight2 = GetElasticFlight(returnTrip.Flights, 2)
		elasticFare.ReturnFlight3 = GetElasticFlight(returnTrip.Flights, 3)
	}

	// Airlines, planes, via cities
	planes := map[string]struct{}{}
	airlines := map[string]struct{}{}
	via := map[string]struct{}{}

	for _, itinerary := range fare.Itineraries {
		for _, flight := range itinerary.Flights {
			for _, plane := range flight.Planes {
				planes[plane] = struct{}{}
			}

			airline := GetAirlineName(flight.Airline)
			airlines[airline] = struct{}{}

			city := GetCityName(flight.ToAirport)
			if (city != elasticFare.Origin.City && city != elasticFare.Destination.City) {
				via[city] = struct {}{}
			}
		}
	}

	for key := range(planes) {
		elasticFare.Planes = append(elasticFare.Planes, key)
	}

	for airline := range(airlines) {
		elasticFare.Airlines = append(elasticFare.Airlines, airline)
	}

	for city := range(via) {
		elasticFare.ViaCities = append(elasticFare.ViaCities, city)
	}

	return &elasticFare
}

func CreateElasticFlight(flight *agents.Flight, nextFlight *agents.Flight) *ElasticFlight {
	// The unique key of the flight: departure date + airline IATA code + flight number
	key := flight.DepartureTime.Format("0102") + flight.Airline + flight.FlightNumber

	elasticFlight := ElasticFlight{
		Key:           key,
		DepartureDate: GetElasticDate(flight.DepartureTime),
		ArrivalDate:   GetElasticDate(flight.ArrivalTime),
		Origin:        GetElasticAirport(flight.FromAirport),
		Destination:   GetElasticAirport(flight.ToAirport),
		Airline:       GetAirlineName(flight.Airline),
		FlightNumber:  flight.Airline + flight.FlightNumber,
		Planes:        flight.Planes,
	}

	if (nextFlight != nil) {
		elasticFlight.StopOver = int(nextFlight.DepartureTime.Sub(flight.ArrivalTime).Hours())
	}

	return &elasticFlight
}

func GetElasticFlight(itinerary  []agents.Flight, flightSequenceNumber int) *ElasticFlight {
	if (len(itinerary) >= flightSequenceNumber) {
		if (len(itinerary) > flightSequenceNumber) {
			return CreateElasticFlight(&itinerary[flightSequenceNumber - 1], &itinerary[flightSequenceNumber])
		} else {
			return CreateElasticFlight(&itinerary[flightSequenceNumber - 1], nil)
		}
	} else {
		return nil
	}
}

func GetElasticDate(date time.Time) ElasticDate {
	return ElasticDate{
		Time: 		ElasticTime(date),
		DayOfTheWeek: 	date.Weekday().String(),
		LocalTime: 	date.Format(LocalDateFormat),
	}
}

func GetElasticAirport(airportIataCode string) ElasticAirport {
	elasticAirport := ElasticAirport{
		IATACode: airportIataCode,
	}

	airport, err := airports.GetByIATACode(airportIataCode)
	if (err != nil) {
		return elasticAirport
	}

	elasticAirport.City = airport.City

	lat, err := strconv.ParseFloat(airport.Latitude, 64)
	if (err != nil) {
		return elasticAirport
	}

	lon, err := strconv.ParseFloat(airport.Longitude, 64)
	if (err != nil) {
		return elasticAirport
	}

	elasticAirport.Geo = Location{
		Latitude: lat,
		Longitude: lon,
	}

	return elasticAirport
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