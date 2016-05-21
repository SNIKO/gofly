package momondo

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/sniko/gofly/references/airports"
	"errors"
	"time"
)

func NewRequest(adultCount int, culture string, ticketClass string, segments []Direction) *SearchRequest {
	return &SearchRequest {
		adultCount,
		"mobileapp",
		make([]int, 0),
		culture,
		true,
		ticketClass,
		segments,
	}
}

func StartSearch(request SearchRequest) (*SearchSessionInfo, error) {
	body, err := json.Marshal(request)
	if (err != nil) {
		return nil, err;
	}

	response, err := http.Post("http://api.momondo.com/api/3.0/FlightSearch", "application/json", bytes.NewReader(body))
	if (err != nil) {
		return nil, err;
	}

	defer response.Body.Close()

	var res SearchSessionInfo
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&res)
	if (err != nil) {
		return nil, err
	}

	return &res, nil
}

func PollSearchResults(searchId string, engineId int) (*SearchResult, error) {
	url := fmt.Sprintf("http://api.momondo.com/api/3.0/FlightSearch/%s/%d/true", searchId, engineId)
	response, err := http.Get(url)
	if (err != nil) {
		return nil, err
	}

	defer response.Body.Close()

	var res SearchResult
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&res)
	if (err != nil) {
		return nil, err
	}

	err = adjustTimeZones(&res)
	if (err != nil) {
		return nil, err
	}

	return &res, nil
}

func adjustTimeZones(result *SearchResult) error {
	if (result.Legs == nil) {
		return nil
	}

	for i, _ := range *result.Legs {
		leg := &(*result.Legs)[i]

		if (len(leg.Key) != 24) {
			return errors.New(fmt.Sprintf("Unknown key format. Can't determine airports from key '%s' and thus cannot adjust time zones.", leg.Key))
		}

		// Trying to determine airports from key as at this stage we don't have access to complete reference files received from poll results.
		// The key format is: OS0025VIE01162320BKK1520
		origin := leg.Key[6:9]
		destination := leg.Key[17:20]

		originLocation, err := airports.GetLocation(origin)
		if (err != nil) {
			return errors.New(fmt.Sprintf("Cannot find the time zone for origin '%s' extracted from leg key '%s': %s", origin, leg.Key, err))
		}

		destinationLocation, err := airports.GetLocation(destination)
		if (err != nil) {
			return errors.New(fmt.Sprintf("Cannot find the time zone for destination'%s' extracted from leg key '%s': %s", destination, leg.Key, err))
		}

		t := time.Time(leg.DepartureDateRaw)
		leg.DepartureDate = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), originLocation)

		t = time.Time(leg.ArrivalDateRaw)
		leg.ArrivalDate  = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), destinationLocation)
	}

	return nil
}