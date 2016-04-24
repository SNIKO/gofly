package elasticsearch

import (
	"fmt"
	"encoding/json"
	"strings"
)

type IndexResult struct {
	Shards       ShardsInfo `json:"_shards"`
	index        string     `json:"_index"`
	DocumentType string     `json:"_type"`
	Id           string     `json:"_id"`
	Version      int        `json:"_version"`
	Created      bool
}

type ShardsInfo struct {
	Total      int
	Failed     int
	Successful int
}

func (c *Client) Index(indexName string, documentType string, document interface{}) (*IndexResult, error) {
	path := fmt.Sprintf("%s/%s", indexName, documentType)

	response, err := c.post(path, document)
	if (err != nil) {
		return nil, err
	}

	var result IndexResult

	decoder := json.NewDecoder(strings.NewReader(response.Body))
	err = decoder.Decode(&result)
	if (err != nil) {
		return nil, err
	}

	return &result, nil
}
