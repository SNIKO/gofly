package elasticsearch

import (
	"bytes"
	"fmt"
	"encoding/json"
	"strings"
)

type BulkCommand interface {
	GetBody() (string, error)
}

type BulkResponse struct {
	Took   int 				`json:"took"`
	Errors bool 				`json:"errors"`
	Items  []map[string]BulkResponseItem 	`json:"items"`
}

type BulkResponseItem struct {
	Index        string 		`json:"_index"`
	DocumentType string 		`json:"_type"`
	Id           string 		`json:"_id"`
	Version      int 		`json:"_version"`
	Status       int 		`json:"status"`
	Error        *ErrorDetails 	`json:"error"`
}

func (c *Client) Bulk(commands []BulkCommand) (*BulkResponse, error) {
	var requestBody bytes.Buffer

	for _, command := range commands {
		body, err := command.GetBody()
		if (err != nil) {
			return nil, err
		}

		_, err = requestBody.WriteString(fmt.Sprintf("%s\n", body))
		if (err != nil) {
			return nil, err
		}
	}

	response, err := c.post("_bulk", requestBody.String())
	if (err != nil) {
		return nil, err
	}

	var result BulkResponse

	decoder := json.NewDecoder(strings.NewReader(response.Body))
	err = decoder.Decode(&result)
	if (err != nil) {
		return nil, err
	}

	return &result, nil
}