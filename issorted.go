package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i <= len(a); i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
