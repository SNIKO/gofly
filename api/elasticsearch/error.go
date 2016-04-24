package elasticsearch

type Error struct {
	Details ErrorDetails `json:"error"`
	Status  int
}

type ErrorDetails struct {
	Reason    string
	ErrorType string          `json:"type"`
	RootCause []*ErrorDetails `json:"root_cause"`
	CausedBy  *ErrorDetails   `json:"coused_by"`
	Status    string
}

func (e Error) Error() string {
	return e.Details.Reason
}
