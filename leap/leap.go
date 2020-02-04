package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%400 == 0 && year%100 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
