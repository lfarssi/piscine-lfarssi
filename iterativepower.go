package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	pow := 1
	for i := 1; i <= power; i++ {
		pow *= nb
	}
	return pow
}
