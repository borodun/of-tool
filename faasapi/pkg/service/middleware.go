package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(FaasapiService) FaasapiService

type loggingMiddleware struct {
	logger log.Logger
	next   FaasapiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a FaasapiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next FaasapiService) FaasapiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ListFunctions(ctx context.Context) (functions []string, err error) {
	defer func() {
		l.logger.Log("method", "ListFunctions", "functions", functions, "err", err)
	}()
	return l.next.ListFunctions(ctx)
}
func (l loggingMiddleware) InvokeFunction(ctx context.Context, functionName string, requestBody string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "InvokeFunction", "functionName", functionName, "requestBody", requestBody, "rs", rs, "err", err)
	}()
	return l.next.InvokeFunction(ctx, functionName, requestBody)
}
