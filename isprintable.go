package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for _, i := range r {
		if i == '\' {
			return false
		}
	}
	return true
}
