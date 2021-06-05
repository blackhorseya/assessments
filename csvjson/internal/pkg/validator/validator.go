package validator

import "regexp"

var r = regexp.MustCompile("\\d{3}-\\d{3}-\\d{4}")

// ValidPhone serve caller to valid phone number format is XXX-XXX-XXXX
func ValidPhone(number string) bool {
	return r.MatchString(number)
}
