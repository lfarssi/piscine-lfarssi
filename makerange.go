package piscine

func MakeRange(min, max int) []int {
	r := make([]int, max-min)
	if max <= min {
		return nil
	}
	for i := min; i < max; i++ {
		r[i-min] += i
	}
	return r
}
