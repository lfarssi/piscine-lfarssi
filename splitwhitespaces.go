package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	var r string
	for _, i := range s {
		if i == 32 || i == '\t' || i == '\n' {
			if len(r) > 0 {
				arr = append(arr, r)
			}
		} else {
			r += string(i)
		}
	}
	if len(r) > 0 {
		arr = append(arr, r)
	}
	return arr
}
