package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "faasapi/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

// makeListFunctionsHandler creates the handler logic
func makeListFunctionsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/list-functions").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.ListFunctionsEndpoint, decodeListFunctionsRequest, encodeListFunctionsResponse, options...)))
}

// decodeListFunctionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListFunctionsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListFunctionsRequest{}
	return req, nil
}

// encodeListFunctionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListFunctionsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeInvokeFunctionHandler creates the handler logic
func makeInvokeFunctionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/invoke-function/{function}/{requestBody}").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.InvokeFunctionEndpoint, decodeInvokeFunctionRequest, encodeInvokeFunctionResponse, options...)))
}

// decodeInvokeFunctionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeInvokeFunctionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	function, ok := vars["function"]
	if !ok {
		return nil, errors.New("not a valid function")
	}

	requestBody, ok2 := vars["requestBody"]
	if !ok2 {
		return nil, errors.New("not a valid request body")
	}

	req := endpoint.InvokeFunctionRequest{
		FunctionName: function,
		RequestBody:  requestBody,
	}
	return req, nil
}

// encodeInvokeFunctionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeInvokeFunctionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
