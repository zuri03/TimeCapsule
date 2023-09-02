package capsule

import (
	"net/http"

	"go.uber.org/zap"
)

type CapsuleContoller struct {
	service *CapsuleService
	logger  *zap.SugaredLogger
}

func Controller(service *CapsuleService, logger *zap.SugaredLogger) *CapsuleContoller {
	return &CapsuleContoller{service: service, logger: logger}
}

func (controller *CapsuleContoller) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	value := request.Context().Value(CapsuleKey)
	if value == nil {
		controller.logger.Error("context value is nil on valid request")
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	capsule, ok := value.(Capsule)
	if !ok {
		controller.logger.Error("context value is not type capsule")
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := controller.service.validateCapsule(&capsule); err != nil {
		controller.logger.Info("Invalid request: %s", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := controller.service.uploadCapsule(&capsule); err != nil {
		controller.logger.Fatal("Error uploading capsule: %s", err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
