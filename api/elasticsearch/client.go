package elasticsearch

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"errors"
	"io"
	"bytes"
)

const (
	DefaultHost = "localhost"
	DefaultPort = 9200
)

type Client struct {
	Host   string
	Port   int
	client *http.Client
}

type Response struct {
	StatusCode int
	Body       string
}

func NewClient(host string, port int) *Client {
	return  &Client{host, port, &http.Client{}}
}

func NewDefaultClient() *Client {
	return &Client{DefaultHost, DefaultPort, &http.Client{}}
}

func (c *Client) put(path string, body interface{}) (*Response, error) {
	return c.sendRequest("PUT", path, body)
}

func (c *Client) post(path string, body interface{}) (*Response, error) {
	return c.sendRequest("POST", path, body)
}

func (c *Client) head(path string, ignoreErrors ...int) (*Response, error) {
	return c.sendRequest("HEAD", path, nil, ignoreErrors...)
}

func (c *Client) sendRequest(method string, path string, body interface{}, ignoreErrors ...int) (*Response, error) {
	url := fmt.Sprintf("http://%s:%d/%s", c.Host, c.Port, path)

	bodyReader, err := getBodyReader(body)
	if (err != nil) {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bodyReader)
	if (err != nil) {
		return nil, err
	}

	response, err := c.client.Do(request)
	if (err != nil) {
		return nil, err
	}

	if (response.Body != nil) {
		defer response.Body.Close()
	}

	err = CheckErrorCode(response, ignoreErrors...)
	if (err != nil) {
		return nil, err
	}

	bodyAsString, err := ioutil.ReadAll(response.Body)
	if (err != nil) {
		return nil, err
	}

	return &Response{response.StatusCode, string(bodyAsString)}, nil
}

func getBodyReader(body interface{}) (io.Reader, error) {
	switch b := body.(type) {
	case string:
		return strings.NewReader(b), nil
	default:
		json, err := json.Marshal(body)
		if (err != nil) {
			return nil, err
		}

		return bytes.NewReader(json), nil
	}
}

func CheckErrorCode(response *http.Response, ignoreErrors ...int) error {
	if (response.StatusCode >= 200 && response.StatusCode <= 299) {
		return nil
	}

	for _, code := range ignoreErrors {
		if (code == response.StatusCode) {
			return nil
		}
	}

	if (response.Body == nil) {
		return errors.New(response.Status)
	}

	var errorResult Error

	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&errorResult)
	if (err != nil) {
		return err
	}

	return errorResult
}