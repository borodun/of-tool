package service

import (
	"context"
	"encoding/json"
	"faasapi/pkg/io"
	"fmt"
	"github.com/imroc/req"
	"log"
)

const OpenFaaSAPI = "http://178.170.194.224:31112"

// FaasapiService describes the service.
type FaasapiService interface {
	ListFunctions(ctx context.Context) (functions []string, err error)
	InvokeFunction(ctx context.Context, functionName string, requestBody string) (rs string, err error)
}

type basicFaasapiService struct{}

func (b *basicFaasapiService) ListFunctions(ctx context.Context) (functions []string, err error) {
	requestStr := fmt.Sprintf("/system/functions")
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic YWRtaW46MzRvMklIcVRMMUVzbE4zMkY5MTB5UDdrQw==",
	}

	r, err := req.Get(OpenFaaSAPI+requestStr, authHeader)
	log.Printf("%+v", r)

	resp := r.String()

	var decodedFuncs []io.FaaSFunction
	json.Unmarshal([]byte(resp), &decodedFuncs)

	var funcs []string
	for _, el := range decodedFuncs {
		funcs = append(funcs, el.Name)
	}

	return funcs, err
}
func (b *basicFaasapiService) InvokeFunction(ctx context.Context, functionName string, requestBody string) (rs string, err error) {
	requestStr := fmt.Sprintf("/function/%s", functionName)
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic YWRtaW46MzRvMklIcVRMMUVzbE4zMkY5MTB5UDdrQw==",
	}

	r, err := req.Get(OpenFaaSAPI+requestStr, authHeader, requestBody)
	log.Printf("%+v", r)

	resp := r.String()

	return resp, err
}

// NewBasicFaasapiService returns a naive, stateless implementation of FaasapiService.
func NewBasicFaasapiService() FaasapiService {
	return &basicFaasapiService{}
}

// New returns a FaasapiService with all of the expected middleware wired in.
func New(middleware []Middleware) FaasapiService {
	var svc FaasapiService = NewBasicFaasapiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
