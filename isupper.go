package piscine

func IsUpper(s string) bool {
	for i := range s {
		if i >= 'A' && i <= 'Z' {
			return true
		}
	}
	return false
}
