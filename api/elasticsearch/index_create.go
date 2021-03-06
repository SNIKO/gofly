package elasticsearch

import (
	"encoding/json"
	"strings"
)

func (c *Client) CreateIndex(index string, indexConfig string) (*Acknowledge, error) {
	response, err := c.put(index, indexConfig)
	if (err != nil) {
		return nil, err
	}

	var result Acknowledge

	decoder := json.NewDecoder(strings.NewReader(response.Body))
	err = decoder.Decode(&result)
	if (err != nil) {
		return nil, err
	}

	return &result, nil
}