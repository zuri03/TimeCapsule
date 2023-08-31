package capsuleserver

import (
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

type CapsuleRouter struct {
	waitgroup *sync.WaitGroup
	logger    *zap.SugaredLogger
}

func (router *CapsuleRouter) serveHTTP(writer http.ResponseWriter, request *http.Request) {
	router.waitgroup.Add(1)
	defer router.waitgroup.Done()

	router.logger.Info("Method: ", request.Method, "Path: ", request.URL.Path)

	switch request.Method {
	case http.MethodPost:
		capsule := ServeHTTPPost(writer, request, router.logger)
		if capsule == nil {
			return
		}
		return
	default:
		http.Error(writer, fmt.Sprintf("Method %s is not supported on this endpoint", request.Method), http.StatusMethodNotAllowed)
	}
}

func Router(wg *sync.WaitGroup) *CapsuleRouter {
	return &CapsuleRouter{waitgroup: wg}
}
