package piscine

func Map(f func(int) bool, a []int) []bool {
	bl := make([]bool, len(a))
	for _, i := range a {
		bl = append(bl, f(i))
	}
	return bl
}
