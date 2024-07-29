package piscine

import "math"

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	fac := 1
	for i := 1; i <= nb; i++ {
		if fac > math.MaxInt32/i {
			return 0
		}
		fac *= i
	}

	return fac
}
