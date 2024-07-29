package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	fac := 1
	for i := 1; i <= nb; i++ {
		if fac > 2147483647/i {
			return 0
		}
		fac *= i
	}

	return fac
}
