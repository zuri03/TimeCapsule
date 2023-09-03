package capsule

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"go.uber.org/zap"
)

const (
	CapsuleKey       = "request-capsule"
	MetaDataKey      = "metadata"
	MAX_CONTENT_SIZE = 32 << 10 //32kb TODO: Adjust this limit
)

type RequestPipeline struct {
	logger    *zap.SugaredLogger
	waitgroup *sync.WaitGroup
}

func Pipeline(waitgroup *sync.WaitGroup, logger *zap.SugaredLogger) *RequestPipeline {
	return &RequestPipeline{waitgroup: waitgroup, logger: logger}
}

func (pipeline *RequestPipeline) ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		pipeline.waitgroup.Add(1)
		defer pipeline.waitgroup.Done()

		contentType := request.Header.Get("Content-Type")
		dataType := strings.Split(contentType, ";")[0]
		if dataType != "multipart/form-data" {
			pipeline.logger.Info("Incorrect content-type got: ", contentType)
			http.Error(writer, "Content type must be multipart/form-data", http.StatusUnsupportedMediaType)
			return
		}

		if request.Method != http.MethodPost {
			pipeline.logger.Info("Incorrect method got: ", request.Method)
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func (pipeline *RequestPipeline) ParseCapsuleFromRequest(controller *CapsuleContoller) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//This must be called first to make request.MultipartForm available
		err := request.ParseMultipartForm(MAX_CONTENT_SIZE)
		if err != nil {
			pipeline.logger.Info(fmt.Sprintf("Error on request.ParseMultipartForm: %s", err))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		metadata := request.FormValue(MetaDataKey)
		if metadata == "" {
			pipeline.logger.Info("Request missing metadata")
			http.Error(writer, "Missing metadata from request", http.StatusBadRequest)
			return
		}

		var capsuleMetaData CapsuleMetaData
		if err := json.Unmarshal([]byte(metadata), &capsuleMetaData); err != nil {
			pipeline.logger.Info(fmt.Sprintf("Error unmarshalling meta json: %s", err))
			http.Error(writer, "Unable to marshal request", http.StatusBadRequest)
			return
		}

		capsule := Capsule{Meta: capsuleMetaData}

		ctx := context.WithValue(request.Context(), CapsuleKey, &capsule)
		request = request.WithContext(ctx)

		controller.ServeHTTP(writer, request)
	})
}
