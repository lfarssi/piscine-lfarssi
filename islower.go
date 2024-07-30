package piscine

func IsLower(s string) bool {
	for _, i := range s {
		if i < 'a' || i > 'z' {
			return false
		}
	}
	return true
}
