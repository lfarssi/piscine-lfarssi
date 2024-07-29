package piscine

func IterativeFactorial(nb int) int {
	fac := 1
	for i := 1; i <= nb; i++ {
		fac *= i
	}
	return fac
}
