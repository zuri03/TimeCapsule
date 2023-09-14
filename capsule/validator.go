package capsule

import (
	"regexp"
)

var (
	emailExpression = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	phoneExpression = regexp.MustCompile("^(\\+\\d{1,2}\\s?)?\\(?\\d{3}\\)?[\\s.-]?\\d{3}[\\s.-]?\\d{4}$")
)

//Ensures that all of the value passed into this function is a valid email address
//Perhaps regex can be used instead if it is faster
func emailValidator(emails ...string) bool {
	for _, emailAddress := range emails {
		if !emailExpression.MatchString(emailAddress) {
			return false
		}
	}

	return true
}

//Ensures that all of the value passed in is a valid phone number
func phoneNumberValidator(phoneNumbers ...string) bool {
	for _, phoneNumber := range phoneNumbers {
		if !phoneExpression.MatchString(phoneNumber) {
			return false
		}
	}

	return true
}
