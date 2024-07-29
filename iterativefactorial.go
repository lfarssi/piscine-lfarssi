package piscine

const maxInt = 1<<31 - 1

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	fac := 1
	for i := 1; i <= nb; i++ {
		if fac > maxInt/i {
			return 0
		}
		fac *= i
	}

	return fac
}
