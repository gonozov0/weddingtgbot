package phone_utils

import (
	"strings"
	"unicode"
)

func NormalizePhoneNumber(phoneNumber string) string {
	var normalized strings.Builder
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			normalized.WriteRune(r)
		}
	}
	return normalized.String()
}

func ArePhoneNumbersEqual(phone1, phone2 string) bool {
	return NormalizePhoneNumber(phone1) == NormalizePhoneNumber(phone2)
}
