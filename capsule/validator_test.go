package capsule

import (
	"testing"
)

type validatorTestCase struct {
	name          string
	input         string
	expectedValue bool
}

func TestEmailValidator(t *testing.T) {
	emailTests := []validatorTestCase{
		{
			name:          "Empty string test",
			input:         "",
			expectedValue: false,
		},
		{
			name:          "Whitespace test",
			input:         "    ",
			expectedValue: false,
		},
		{
			name:          "Valid email test",
			input:         "jon@gmail.com",
			expectedValue: true,
		},
		{
			name:          "Valid email with random domain test",
			input:         "jane@yahoo.com",
			expectedValue: true,
		},
		{
			name:          "Missing domain test",
			input:         "jon@.com",
			expectedValue: false,
		},
		{
			name:          "Missing '@' test",
			input:         "jon#email.com",
			expectedValue: false,
		},
		{
			name:          "Missing username est",
			input:         "@gmail.com",
			expectedValue: false,
		},
		{
			name:          "Username is an integer test",
			input:         "9@gmail.com",
			expectedValue: true,
		},
		{
			name:          "Missing '.' est",
			input:         "jon@gmailcom",
			expectedValue: false,
		},
		{
			name:          "Valid email differnt top domain test",
			input:         "jon@gmail.edu",
			expectedValue: true,
		},
	}

	for _, testCase := range emailTests {
		t.Run(testCase.name, func(t *testing.T) {
			actualValue := emailValidator(testCase.input)

			if actualValue != testCase.expectedValue {
				t.Errorf("Expected %t", testCase.expectedValue)
			}
		})
	}
}
