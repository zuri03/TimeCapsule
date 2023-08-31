package capsuleserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

const (
	MetaDataKey      = "metadata"
	MAX_CONTENT_SIZE = 32 << 10 //32kb TODO: Adjust this limit
)

func ServeHTTPPost(writer http.ResponseWriter, request *http.Request, logger *zap.SugaredLogger) *Capsule {
	//This must be called first to make request.MultipartForm available
	err := request.ParseMultipartForm(MAX_CONTENT_SIZE)
	if err != nil {
		logger.Error(fmt.Sprintf("Error on request.ParseMultipartForm: %s", err.Error()))
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	metadataFile, header, err := request.FormFile(MetaDataKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Error on request.ParseMultipartForm: %s", err.Error()))
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	if header.Size > MAX_CONTENT_SIZE {
		logger.Warn(fmt.Sprintf("Error on request.ParseMultipartForm: %s", err.Error()))
		//return http error
		return nil
	}

	metadataBytes, err := ioutil.ReadAll(metadataFile)
	if err != nil {
		logger.Error(fmt.Sprintf("Error on request.ParseMultipartForm: %s", err.Error()))
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	var capsuleMetaData CapsuleMetaData
	if err := json.Unmarshal(metadataBytes, &capsuleMetaData); err != nil {
		logger.Error(fmt.Sprintf("Error on request.ParseMultipartForm: %s", err.Error()))
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	capsule := &Capsule{Meta: capsuleMetaData}

	return capsule
}
