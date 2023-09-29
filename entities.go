package bigqueryudf

// bqRequest is a struct that represents a request body from BigQuery.
type bqRequest struct {
	RequestID          string            `json:"requestId"`
	Caller             string            `json:"caller"`
	SessionUser        string            `json:"sessionUser"`
	UserDefinedContext map[string]string `json:"userDefinedContext"`
	Calls              [][]interface{}   `json:"calls"`
}

// bqResponse is a struct that represents a response body from BigQuery.
type bqResponse struct {
	Replies      []string `json:"replies,omitempty"`
	ErrorMessage string   `json:"errorMessage,omitempty"`
}
