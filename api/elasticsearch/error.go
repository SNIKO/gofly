package elasticsearch

import (
	"net/http"
	"fmt"
)

type Error struct {
	Details *ErrorDetails `json:"error"`
	Status  int
}

type ErrorDetails struct {
	Reason    string
	ErrorType string          `json:"type"`
	RootCause []*ErrorDetails `json:"root_cause"`
	CausedBy  *ErrorDetails   `json:"caused_by"`
	Status    string
}

func (e Error) Error() string {
	if e.Details != nil && e.Details.Reason != "" {
		return fmt.Sprintf("Error %d (%s): %s [type=%s]", e.Status, http.StatusText(e.Status), e.Details.Reason, e.Details.ErrorType)
	} else {
		return fmt.Sprintf("Error %d (%s)", e.Status, http.StatusText(e.Status))
	}
}
