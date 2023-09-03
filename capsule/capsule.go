package capsule

type Capsule struct {
	Id     string          `json:"Id"`
	Meta   CapsuleMetaData `json:"Meta"`
	Conent []byte
}

type CapsuleMetaData struct {
	Type         string `json:"Type"` //Supported types: email, text, phone
	From         string `json:"From"`
	To           string `json:"To"`
	DeliveryDate string `json:"DeliveryDate"`
}
