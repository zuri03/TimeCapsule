package capsule

import (
	"fmt"

	"go.uber.org/zap"
)

//Struct to encapsulate dependencies
type CapsuleService struct {
	logger *zap.SugaredLogger
}

//Function to initialize new struct
func Service(logger *zap.SugaredLogger) *CapsuleService {
	return &CapsuleService{logger: logger}
}

func (service *CapsuleService) validateCapsule(capsule *Capsule) error {
	capsuleType := capsule.Meta.Type
	if capsuleType != "email" && capsuleType != "message" {
		service.logger.Info(fmt.Sprintf("Unexpected capsule type got %s", capsuleType))
		return fmt.Errorf("Invalid capsule type")
	}
	return nil
}

//Function for uploading Capsule to db
func (service *CapsuleService) uploadCapsule(capsule *Capsule) error {
	fmt.Printf("Now uploading capsule to %s\n", capsule.Meta.To)
	return nil
}
