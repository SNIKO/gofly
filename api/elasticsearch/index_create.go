package elasticsearch

import (
	"encoding/json"
	"strings"
)

type IndexCreationResult struct {
	Acknowledged bool
}

func (c *Client) CreateIndex(index string, indexConfig string) (*IndexCreationResult, error) {
	response, err := c.put(index, indexConfig)
	if (err != nil) {
		return nil, err
	}

	var result IndexCreationResult

	decoder := json.NewDecoder(strings.NewReader(response.Body))
	err = decoder.Decode(&result)
	if (err != nil) {
		return nil, err
	}

	return &result, nil
}