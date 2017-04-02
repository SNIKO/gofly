package elasticsearch

import (
	"net/http"
	"fmt"
)

func (c *Client) TemplateExists(templateName string) (bool, error) {
	path := fmt.Sprintf("_template/%s", templateName)

	response, err := c.head(path, http.StatusNotFound)
	if (err != nil) {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}