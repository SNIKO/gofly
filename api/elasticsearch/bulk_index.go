package elasticsearch

import (
	"encoding/json"
	"bytes"
	"fmt"
)

type IndexCommand struct {
	Index        string        	`json:"_index,omitempty"`
	DocumentType string        	`json:"_type,omitempty"`
	Id           string 		`json:"_id,omitempty"`
	Content      interface{}	`json:"-"`
}

func (c IndexCommand) GetBody() (string, error) {
	var command bytes.Buffer

	indexCommand, err := json.Marshal(c)
	if (err != nil) {
		return "", err
	}

	header := fmt.Sprintf("{ \"index\" : %s }\n", indexCommand)
	command.WriteString(header)

	reader, err := getBodyReader(c.Content)
	if (err != nil) {
		return "", err
	}

	_, err = command.ReadFrom(reader)
	if (err != nil) {
		return "", err
	}

	return command.String(), nil
}