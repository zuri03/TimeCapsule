package capsule

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

		metadataFile, header, err := request.FormFile(MetaDataKey)
		if err != nil {
			pipeline.logger.Error(fmt.Sprintf("Error fetching form file from request: %s", err))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if header.Size > MAX_CONTENT_SIZE {
			pipeline.logger.Info(fmt.Sprintf("Metadata was %d bytes, request rejected", header.Size))
			http.Error(writer, "Content too large", http.StatusRequestEntityTooLarge)
			return
		}

		metadataBytes, err := ioutil.ReadAll(metadataFile)
		if err != nil {
			pipeline.logger.Error(fmt.Sprintf("Error reading bytes from metadata file: %s", err))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var capsuleMetaData CapsuleMetaData
		if err := json.Unmarshal(metadataBytes, &capsuleMetaData); err != nil {
			pipeline.logger.Error(fmt.Sprintf("Error unmarshalling meta json: %s", err))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		capsule := &Capsule{Meta: capsuleMetaData}

		ctx := context.WithValue(request.Context(), CapsuleKey, capsule)
		request = request.WithContext(ctx)

		controller.ServeHTTP(writer, request)
	})
}
