package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 12 {
		return 0
	}
	return nb * IterativeFactorial(nb-1)
}
