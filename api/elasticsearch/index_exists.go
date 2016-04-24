package elasticsearch

import "net/http"

func (c *Client) IndexExists(indexName string) (bool, error) {
	response, err := c.head(indexName, http.StatusNotFound)
	if (err != nil) {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}