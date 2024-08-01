package piscine

func AppendRange(min, max int) []int {
	r := []int{}
	if max <= min {
		return r
	}
	for i := min; i <= max; i++ {
		r = append(r, i)
	}
	return r
}
