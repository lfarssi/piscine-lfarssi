package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n < 0 || n > len(r) {
		return 0
	}
	return r[n-1]
}
