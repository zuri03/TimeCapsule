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
	if capsuleType != "email" && capsuleType != "text" {
		service.logger.Info(fmt.Sprintf("Unexpected capsule type got %s", capsuleType))
		return fmt.Errorf("Invalid capsule type")
	}

	var validator func(value ...string) bool
	switch capsuleType {
	case "email":
		validator = emailValidator
		break
	case "text":
		validator = phoneNumberValidator
		break
	}

	if !validator(capsule.Meta.From, capsule.Meta.To) {
		service.logger.Info(fmt.Sprintf("Request had invalid To or From formats, To: %s, From: %s", capsule.Meta.To, capsule.Meta.From))
		return fmt.Errorf("Invalid To or From fields")
	}

	return nil
}

//Function for uploading Capsule to db
func (service *CapsuleService) uploadCapsule(capsule *Capsule) error {
	fmt.Printf("Now uploading capsule to %s\n", capsule.Meta.To)
	return nil
}
