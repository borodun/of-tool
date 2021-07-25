package endpoint

import (
	"context"
	service "faasapi/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// ListFunctionsRequest collects the request parameters for the ListFunctions method.
type ListFunctionsRequest struct{}

// ListFunctionsResponse collects the response parameters for the ListFunctions method.
type ListFunctionsResponse struct {
	Functions []string `json:"functions"`
	Err       error    `json:"err"`
}

// MakeListFunctionsEndpoint returns an endpoint that invokes ListFunctions on the service.
func MakeListFunctionsEndpoint(s service.FaasapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		functions, err := s.ListFunctions(ctx)
		return ListFunctionsResponse{
			Err:       err,
			Functions: functions,
		}, nil
	}
}

// Failed implements Failer.
func (r ListFunctionsResponse) Failed() error {
	return r.Err
}

// InvokeFunctionRequest collects the request parameters for the InvokeFunction method.
type InvokeFunctionRequest struct {
	FunctionName string `json:"function_name"`
	RequestBody  string `json:"request_body"`
}

// InvokeFunctionResponse collects the response parameters for the InvokeFunction method.
type InvokeFunctionResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeInvokeFunctionEndpoint returns an endpoint that invokes InvokeFunction on the service.
func MakeInvokeFunctionEndpoint(s service.FaasapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InvokeFunctionRequest)
		rs, err := s.InvokeFunction(ctx, req.FunctionName, req.RequestBody)
		return InvokeFunctionResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r InvokeFunctionResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// ListFunctions implements Service. Primarily useful in a client.
func (e Endpoints) ListFunctions(ctx context.Context) (functions []string, err error) {
	request := ListFunctionsRequest{}
	response, err := e.ListFunctionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListFunctionsResponse).Functions, response.(ListFunctionsResponse).Err
}

// InvokeFunction implements Service. Primarily useful in a client.
func (e Endpoints) InvokeFunction(ctx context.Context, functionName string, requestBody string) (rs string, err error) {
	request := InvokeFunctionRequest{
		FunctionName: functionName,
		RequestBody:  requestBody,
	}
	response, err := e.InvokeFunctionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(InvokeFunctionResponse).Rs, response.(InvokeFunctionResponse).Err
}
