package onetwotrip

import (
	"time"
	"errors"
	"fmt"
	"strconv"
	"bytes"
	"net/http"
	"encoding/json"
)

type Duration time.Duration

type SearchQuery struct {
	Flights []Flight
}

type Flight struct {
	FromAirport   string
	ToAirport     string
	DepartureDate time.Time
}

type SearchResult struct {
	Fares  []Fare `json:"frs"`
	Trips  []Trip `json:"trps"`
	Planes map[string]string `json:"planes"`
}

type Fare struct {
	Id         int        	`json:"id"`
	Directions []Direction 	`json:"dirs"`
	Price      PriceInfo 	`json:"prcInf"`
}

type Trip struct {
	StartDate   	string `json:"stDt"`
	StartTime 	string `json:"stTm"`
	EndDate   	string `json:"endDate"`
	EndTime 	string `json:"endTm"`
	FlightDuration  Duration `json:"fltTm"`
	JourneyDuration Duration `json:"jrnTm"`
	From            string 	`json:"from"`
	To              string 	`json:"to"`
	AirlineCode     string 	`json:"airCmp"`
	FlightNumber    string 	`json:"fltNm"`
	OperatedBy      string 	`json:"oprdBy"`
	Plane           string 	`json:"plane"`
}

type Direction struct {
	Id    int 	`json:"id"`
	Trips []TripRef `json:"trps"`
}

type TripRef struct {
	Id               int `json:"id"`
	ReservationClass string `json:"cls"`
	CabinClass       string `json:"srvCls"`
}

type PriceInfo struct {
	AdultFare  float64 `json:"adtB"`
	AdultTaxes float64 `json:"adtT"`
	Markup     float64 `json:"markupNew"`
	Currency   string `json:"cur"`
}

func Search(query SearchQuery) (*SearchResult, error) {
	url := getSearchUrl(query)

	response, err := http.Get(url)
	if (err != nil) {
		return nil, err
	}

	defer response.Body.Close()

	var result SearchResult

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)
	if (err != nil) {
		return nil, err
	}

	return &result, nil
}

func getSearchUrl(query SearchQuery) string {
	var route bytes.Buffer

	// A root is concatenation of flights. Each flight has the following format: DDMMFromAirportToAirport
	// For example, route "2408SYDIEV1509IEVSYD" has two flights, a direct one from SYD to IEV on August 24
	// and a return one from IEV to SYD on September 15.
	for _, flight := range query.Flights {
		r := fmt.Sprintf("%s%s%s", flight.DepartureDate.Format("0201"), flight.FromAirport, flight.ToAirport)
		route.WriteString(r)
	}

	url := fmt.Sprintf("https://secure.onetwotrip.com/_api/searching/startSync/?route=%s&ad=1&cs=E", route.String())
	return url
}

func (p *PriceInfo) TotalPrice() float64 {
	return p.AdultFare + p.AdultTaxes + p.Markup
}

func (t *Trip) DepartureTime() (*time.Time, error) {
	dateTime := t.StartDate + t.StartTime

	time, err := time.Parse("200601021504", dateTime)
	if (err != nil) {
		return nil, err
	}

	return &time, nil
}

func (t *Trip) ArrivalTime() (*time.Time, error) {
	var dateTime string
	if (t.EndDate != "") {
		dateTime = t.EndDate + t.EndTime
	} else
	{
		dateTime = t.StartDate + t.EndTime
	}

	time, err := time.Parse("200601021504", dateTime)
	if (err != nil) {
		return nil, err
	}

	return &time, nil
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	if (b[0] == '"' && b[len(b) - 1] == '"') {
		b = b[1 : len(b) - 1]
	}

	if (len(b) != 4) {
		return errors.New(fmt.Sprintf("Failed to parse duration: '%s'", string(b)))
	}

	hours, err := strconv.ParseInt(string(b[:2]), 10, 0)
	if (err != nil) {
		return errors.New(fmt.Sprintf("Failed to parse hours part of duration '%s': %s", string(b), err))
	}

	minutes, err := strconv.ParseInt(string(b[2:]), 10, 0)
	if (err != nil) {
		return errors.New(fmt.Sprintf("Failed to parse minutes part of duration '%s': %s", string(b), err))
	}

	*d = Duration(time.Duration(hours) * time.Hour + time.Duration(minutes) * time.Minute)
	return nil
}