package capsule

type Capsule struct {
	Id     string          `json:"id"`
	Meta   CapsuleMetaData `json:"meta"`
	Conent []byte
}

type CapsuleMetaData struct {
	Type         string `json:"type"` //Supported types: email, text, phone
	From         string `json:"from"`
	To           string `json:"to"`
	DeliveryDate string `json:"deliveryDate"`
}
