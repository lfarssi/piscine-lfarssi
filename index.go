package piscine

func Index(s string, toFind string) int {
	var r int
	for i := range s {
		if i == toFind[0] {
			r = s[i]
		} else {
			r = -1
		}
	}
	return r
}
