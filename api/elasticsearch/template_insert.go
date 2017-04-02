package elasticsearch

import (
	"encoding/json"
	"strings"
	"fmt"
)

func (c *Client) CreateTemplate(name string, template string) (*Acknowledge, error) {
	path := fmt.Sprintf("_template/%s", name)

	response, err := c.put(path, template)
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