package endpoint

import (
	"context"
	service "cpumem/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetCPURequest collects the request parameters for the GetCPU method.
type GetCPURequest struct {
	PodName string `json:"pod_name"`
}

// GetCPUResponse collects the response parameters for the GetCPU method.
type GetCPUResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetCPUEndpoint returns an endpoint that invokes GetCPU on the service.
func MakeGetCPUEndpoint(s service.CpumemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCPURequest)
		rs, err := s.GetCPU(ctx, req.PodName)
		return GetCPUResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetCPUResponse) Failed() error {
	return r.Err
}

// GetMEMRequest collects the request parameters for the GetMEM method.
type GetMEMRequest struct {
	PodName string `json:"pod_name"`
}

// GetMEMResponse collects the response parameters for the GetMEM method.
type GetMEMResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetMEMEndpoint returns an endpoint that invokes GetMEM on the service.
func MakeGetMEMEndpoint(s service.CpumemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMEMRequest)
		rs, err := s.GetMEM(ctx, req.PodName)
		return GetMEMResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetMEMResponse) Failed() error {
	return r.Err
}

// GetNamespacesRequest collects the request parameters for the GetNamespaces method.
type GetNamespacesRequest struct{}

// GetNamespacesResponse collects the response parameters for the GetNamespaces method.
type GetNamespacesResponse struct {
	Namespaces []service.Namespace `json:"namespaces"`
	Err        error               `json:"err"`
}

// MakeGetNamespacesEndpoint returns an endpoint that invokes GetNamespaces on the service.
func MakeGetNamespacesEndpoint(s service.CpumemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		namespaces, err := s.GetNamespaces(ctx)
		return GetNamespacesResponse{
			Err:        err,
			Namespaces: namespaces,
		}, nil
	}
}

// Failed implements Failer.
func (r GetNamespacesResponse) Failed() error {
	return r.Err
}

// GetPodsRequest collects the request parameters for the GetPods method.
type GetPodsRequest struct {
	Namespace string `json:"namespace"`
}

// GetPodsResponse collects the response parameters for the GetPods method.
type GetPodsResponse struct {
	Pods []service.Pod `json:"pods"`
	Err  error         `json:"err"`
}

// MakeGetPodsEndpoint returns an endpoint that invokes GetPods on the service.
func MakeGetPodsEndpoint(s service.CpumemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPodsRequest)
		pods, err := s.GetPods(ctx, req.Namespace)
		return GetPodsResponse{
			Err:  err,
			Pods: pods,
		}, nil
	}
}

// Failed implements Failer.
func (r GetPodsResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetCPU implements Service. Primarily useful in a client.
func (e Endpoints) GetCPU(ctx context.Context, podName string) (rs string, err error) {
	request := GetCPURequest{PodName: podName}
	response, err := e.GetCPUEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCPUResponse).Rs, response.(GetCPUResponse).Err
}

// GetMEM implements Service. Primarily useful in a client.
func (e Endpoints) GetMEM(ctx context.Context, podName string) (rs string, err error) {
	request := GetMEMRequest{PodName: podName}
	response, err := e.GetMEMEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetMEMResponse).Rs, response.(GetMEMResponse).Err
}

// GetNamespaces implements Service. Primarily useful in a client.
func (e Endpoints) GetNamespaces(ctx context.Context) (namespaces []service.Namespace, err error) {
	request := GetNamespacesRequest{}
	response, err := e.GetNamespacesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetNamespacesResponse).Namespaces, response.(GetNamespacesResponse).Err
}

// GetPods implements Service. Primarily useful in a client.
func (e Endpoints) GetPods(ctx context.Context, namespace string) (pods []service.Pod, err error) {
	request := GetPodsRequest{Namespace: namespace}
	response, err := e.GetPodsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetPodsResponse).Pods, response.(GetPodsResponse).Err
}
