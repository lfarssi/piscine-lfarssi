package piscine

func Index(s string, toFind string) int {
	var r int
	si := []rune(s)
	t := []rune(toFind)
	for i := range s {
		if i == t[0] {
			r = si[i]
		} else {
			r = -1
		}
	}
	return r
}
