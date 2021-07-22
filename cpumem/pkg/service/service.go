package service

import (
	"context"
	"cpumem/pkg/io"
	"fmt"
	"github.com/imroc/req"
	"log"
)

var (
	prometheusAPIAddr = "http://178.170.194.224:8080/api/v1/"
)

type Namespace struct {
	Namespace string `json:"namespace"`
}

type Pod struct {
	Pod string `json:"pod"`
}

// CpumemService describes the service.
type CpumemService interface {
	GetCPU(ctx context.Context, podName string) (rs string, err error)
	GetMEM(ctx context.Context, podName string) (rs string, err error)
	GetNamespaces(ctx context.Context) (namespaces []Namespace, err error)
	GetPods(ctx context.Context, namespace string) (pods []Pod, err error)
}

type basicCpumemService struct{}

func (b *basicCpumemService) GetCPU(ctx context.Context, podName string) (rs string, err error) {
	queryStr := fmt.Sprintf("container_cpu_usage_seconds_total{container=\"%s\"}", podName)
	param := req.Param{
		"query": queryStr,
	}
	r, err := req.Get(prometheusAPIAddr+"query", param)
	log.Printf("%+v", r)

	var resp io.PromResponse
	r.ToJSON(&resp)

	return resp.Data.Results[0].Value[1], err
}
func (b *basicCpumemService) GetMEM(ctx context.Context, podName string) (rs string, err error) {
	queryStr := fmt.Sprintf("container_memory_usage_bytes{pod=~\"%s\", container!=\"\", container!=\"POD\"}", podName)
	param := req.Param{
		"query": queryStr,
	}
	r, err := req.Get(prometheusAPIAddr+"query", param)
	log.Printf("%+v", r)

	var resp io.PromResponse
	r.ToJSON(&resp)

	return resp.Data.Results[0].Value[1], err
}

func (b *basicCpumemService) GetNamespaces(ctx context.Context) (namespaces []Namespace, err error) {
	queryStr := fmt.Sprintf("kube_namespace_labels")
	param := req.Param{
		"query": queryStr,
	}
	r, err := req.Get(prometheusAPIAddr+"query", param)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v", r)

	var resp io.PromResponse
	r.ToJSON(&resp)

	namespaces = make([]Namespace, len(resp.Data.Results))
	for i := range namespaces {
		namespaces[i].Namespace = resp.Data.Results[i].Metric.Namespace
	}

	return namespaces, err
}
func (b *basicCpumemService) GetPods(ctx context.Context, namespace string) (pods []Pod, err error) {
	queryStr := fmt.Sprintf("kube_pod_info{namespace=\"%s\"}", namespace)
	param := req.Param{
		"query": queryStr,
	}
	r, err := req.Get(prometheusAPIAddr+"query", param)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v", r)

	var resp io.PromResponse
	r.ToJSON(&resp)

	pods = make([]Pod, len(resp.Data.Results))
	for i := range pods {
		pods[i].Pod = resp.Data.Results[i].Metric.Pod
	}

	return pods, err
}

// NewBasicCpumemService returns a naive, stateless implementation of CpumemService.
func NewBasicCpumemService() CpumemService {
	return &basicCpumemService{}
}

// New returns a CpumemService with all of the expected middleware wired in.
func New(middleware []Middleware) CpumemService {
	var svc CpumemService = NewBasicCpumemService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
