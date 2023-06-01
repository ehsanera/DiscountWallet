package util

import (
	"regexp"
)

func ValidatePhoneNumber(phoneNumber string) bool {
	pattern := `^09\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}
