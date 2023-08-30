package capsuleserver

import (
	"net/http"
)

//Struct to encapsulate dependencies
type CapsuleService struct {
}

//Function to initialize new struct
func New() *CapsuleService {
	return &CapsuleService{}
}

//Reading and parsing request into Capsule struct logic needs to be moved to seperate function
//Function for uploading Capsule to db
func (service *CapsuleService) uploadCapsule(writer http.ResponseWriter, request *http.Request) {

}
