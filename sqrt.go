package piscine

func Sqrt(nb int) int {
	r := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			r = i
		}
	}
	if nb%2 != 0 {
		r = 0
	}
	return r
}
