package piscine

func Join(strs []string, sep string) string {
	r := strs[0]
	for _, i := range strs[1:] {
		r += sep + i
	}
	return r
}
