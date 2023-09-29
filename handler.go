package bigqueryudf

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// NewHandler function is a wrapper for bigquery functional cloud function that works inside UDF (User Defined Function).
func NewHandler(transformationFunc func(input []byte) (any, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response bqResponse
		var request bqRequest
		if err := decodeRequestBody(r, &request); err != nil {
			response.ErrorMessage = fmt.Sprintf("internal Error: error decoding request body: %s", err.Error())
			handleJSONResponse(w, response, http.StatusBadRequest)
			return
		}

		replies := make([]string, len(request.Calls))
		errGroup := errgroup.Group{}
		for i, call := range request.Calls {
			i := i
			call := call
			errGroup.Go(func() error {
				rawxml, ok := call[0].(string)
				if !ok {
					return fmt.Errorf("error invalid call type: %T. Call must be a string", call)
				}

				transformedData, err := transformationFunc([]byte(rawxml))
				if err != nil {
					return fmt.Errorf("error parsing xml: %w", err)
				}

				transformedDataEncoded, err := json.Marshal(transformedData)
				if err != nil {
					return fmt.Errorf("error marshaling response: %w", err)
				}

				replies[i] = string(transformedDataEncoded)
				return nil
			})
		}

		if err := errGroup.Wait(); err != nil {
			handleJSONResponse(w, response, http.StatusInternalServerError)
			return
		}

		response.Replies = replies
		handleJSONResponse(w, response, http.StatusOK)
	}
}

func decodeRequestBody(r *http.Request, request any) error {
	defer func() {
		_ = r.Body.Close()
	}()
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&request)
	if err != nil {
		return fmt.Errorf("decoding http request body: %w", err)
	}
	return nil
}

func handleJSONResponse(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}
