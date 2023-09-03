package capsule

import (
	"net/mail"
	"regexp"
	"strings"
)

//Ensures that all of the value passed into this function is a valid email address
//Perhaps regex can be used instead if it is faster
func emailValidator(emails ...string) bool {
	list := strings.Join(emails, ",")
	_, err := mail.ParseAddressList(list)
	return err != nil
}

//Ensures that all of the value passed in is a valid phone number
func phoneNumberValidator(phoneNumbers ...string) bool {
	expression, _ := regexp.Compile("^(\\+\\d{1,2}\\s?)?\\(?\\d{3}\\)?[\\s.-]?\\d{3}[\\s.-]?\\d{4}$")
	for _, phoneNumber := range phoneNumbers {
		if !expression.MatchString(phoneNumber) {
			return false
		}
	}

	return true
}
