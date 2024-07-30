package piscine

func IsUpper(s string) bool {
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return true
		}
	}
	return false
}
