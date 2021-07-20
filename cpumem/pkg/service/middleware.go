package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(CpumemService) CpumemService

type loggingMiddleware struct {
	logger log.Logger
	next   CpumemService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CpumemService) CpumemService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetCPU(ctx context.Context, podName string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "GetCPU", "podName", podName, "rs", rs, "err", err)
	}()
	return l.next.GetCPU(ctx, podName)
}
func (l loggingMiddleware) GetMEM(ctx context.Context, podName string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "GetMEM", "podName", podName, "rs", rs, "err", err)
	}()
	return l.next.GetMEM(ctx, podName)
}
func (l loggingMiddleware) GetNamespaces(ctx context.Context) (namespaces []Namespace, err error) {
	defer func() {
		l.logger.Log("method", "GetNamespaces", "namespaces", namespaces[0].Namespace, "err", err)
	}()
	return l.next.GetNamespaces(ctx)
}
func (l loggingMiddleware) GetPods(ctx context.Context, namespace string) (pods []Pod, err error) {
	defer func() {
		l.logger.Log("method", "GetPods", "namespace", namespace, "pods", pods, "err", err)
	}()
	return l.next.GetPods(ctx, namespace)
}
