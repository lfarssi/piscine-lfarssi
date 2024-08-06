package piscine

func Map(f func(int) bool, a []int) []bool {
	bl := make([]bool, len(a))
	for i,v := range a {
		bl[i] = f(v)
	}
	return bl
}
