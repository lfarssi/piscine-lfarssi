package piscine

func StrRev(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	rev := string(runes)
	return rev
}
