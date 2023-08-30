package capsuleserver

import (
	"fmt"
	"net/http"
	"sync"
)

type CapsuleRouter struct {
	waitgroup *sync.WaitGroup
}

func (router *CapsuleRouter) serveHTTP(writer http.ResponseWriter, request *http.Request) {
	router.waitgroup.Add(1)
	defer router.waitgroup.Done()

	switch request.Method {
	case http.MethodPost:
		capsule := ServeHTTPPost(writer, request)
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
