package piscine

func ActiveBits(n int) int {
	c := 0
	for n != 0 {
		c += n & 1
		n >>= 1
	}
	return c
}
