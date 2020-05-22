package kata

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	UppercaseRequest struct {
		S string `json:"s"`
	}

	UppercaseResponse struct {
		V   string `json:"v"`
		Err string `json:"err,omitempty"`
	}

	CountRequest struct {
		S string `json:"s"`
	}

	CountResponse struct {
		V int `json:"v"`
	}
)

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
