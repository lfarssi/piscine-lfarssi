package piscine

func Unmatch(a []int) int {
	for i := range a {
		find := false
		for j := range a {
			if a[i] == a[j] && i != j {
				find = true
				a[i] = -1
				a[j] = -1
			}
		}
		if !find {
			return a[i]
		}
	}
	return -1
}
