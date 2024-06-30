package validator

import "regexp"

// Fungsi validasi PIN
func IsValidPin(pin string) bool {
	re := regexp.MustCompile(`^\d{6}$`)
	return re.MatchString(pin)
}
