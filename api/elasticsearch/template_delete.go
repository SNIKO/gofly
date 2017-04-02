package elasticsearch

import (
	"encoding/json"
	"strings"
	"fmt"
)

func (c *Client) DeleteTemplate(templateName string) (*Acknowledge, error) {
	path := fmt.Sprintf("_template/%s", templateName)

	response, err := c.delete(path)
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