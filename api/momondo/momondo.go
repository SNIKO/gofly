package momondo

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/sniko/gofly/references/airports"
	"errors"
	"time"
	"compress/gzip"
)

func NewRequest(adultCount int, culture string, ticketClass string, segments []Direction) *SearchRequest {
	return &SearchRequest{
		AdultCount: adultCount,
		ChildAges: make([]int, 0),
		Culture: culture,
		TicketClass: ticketClass,
		Segments: segments,
		Mix: "Segments",
		Market: "",
		DirectOnly: false,
		IncludeNearby: false,
	}
}

func StartSearch(request SearchRequest) (*SearchSessionInfo, error) {
	body, err := json.Marshal(request)
	if (err != nil) {
		return nil, err;
	}

	req, err := http.NewRequest("POST", "http://api.momondo.com/api/3.0/FlightSearch", bytes.NewReader(body))
	if (err != nil) {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return nil, err;
	}

	defer response.Body.Close()

	reader, err := gzip.NewReader(response.Body)
	if (err != nil){
		return  nil, err
	}

	defer reader.Close()

	var res SearchSessionInfo
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&res)
	if (err != nil) {
		return nil, err
	}

	return &res, nil
}

func PollSearchResults(searchId string, engineId int) (*SearchResult, error) {
	url := fmt.Sprintf("http://api.momondo.com/api/3.0/FlightSearch/%s/%d/true", searchId, engineId)

	req, err := http.NewRequest("GET", url, nil)
	if (err != nil) {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	response, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return nil, err
	}

	defer response.Body.Close()

	reader, err := gzip.NewReader(response.Body)
	if (err != nil){
		return  nil, err
	}

	defer reader.Close()

	var res SearchResult
	decoder := json.NewDecoder(reader)
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