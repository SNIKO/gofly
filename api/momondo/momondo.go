package momondo

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
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

	return &res, nil
}