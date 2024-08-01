package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
		}
	r := make([]int, max-min)
	for i := min; i < max; i++ {
		r[i-min] = i
	}
	return r
}
