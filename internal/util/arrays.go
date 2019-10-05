package util

// returns true if array s contains e, otherwise false
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
