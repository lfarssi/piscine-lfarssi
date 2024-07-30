package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if i == '\' {
			return false
		}
	}
	return true
}
