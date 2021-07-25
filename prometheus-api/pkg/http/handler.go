package http

import (
	"context"
	endpoint "cpumem/pkg/endpoint"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeGetCPUHandler creates the handler logic
func makeGetCPUHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/get-cpu/{pod_name}").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.GetCPUEndpoint, decodeGetCPURequest, encodeGetCPUResponse, options...)))
}

// decodeGetCPURequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetCPURequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name, ok := vars["pod_name"]
	if !ok {
		return nil, errors.New("not a valid pod_name")
	}
	req := endpoint.GetCPURequest{
		PodName: name,
	}
	return req, nil
}

// encodeGetCPUResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetCPUResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetMEMHandler creates the handler logic
func makeGetMEMHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/get-mem/{pod_name}").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.GetMEMEndpoint, decodeGetMEMRequest, encodeGetMEMResponse, options...)))
}

// decodeGetMEMRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetMEMRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name, ok := vars["pod_name"]
	if !ok {
		return nil, errors.New("not a valid pod_name")
	}
	req := endpoint.GetMEMRequest{
		PodName: name,
	}
	return req, nil
}

// encodeGetMEMResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetMEMResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetPodsHandler creates the handler logic
func makeGetPodsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/get-pods/{namespace}").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.GetPodsEndpoint, decodeGetPodsRequest, encodeGetPodsResponse, options...)))
}

// decodeGetPodsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetPodsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	namespace, ok := vars["namespace"]
	if !ok {
		return nil, errors.New("not a valid namespace")
	}
	req := endpoint.GetPodsRequest{
		Namespace: namespace,
	}
	return req, nil
}

// encodeGetPodsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetPodsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
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

// makeGetNamespacesHandler creates the handler logic
func makeGetNamespacesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/get-namespaces").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.GetNamespacesEndpoint, decodeGetNamespacesRequest, encodeGetNamespacesResponse, options...)))
}

// decodeGetNamespacesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetNamespacesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetNamespacesRequest{}
	return req, nil
}

// encodeGetNamespacesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetNamespacesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
