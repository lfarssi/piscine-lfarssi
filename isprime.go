package piscine

func IsPrime(nb int) bool {
	var r bool
	if nb <= 1 {
		return false
	} else if nb == 2 {
		r = true
	} else if nb%2 == 0 {
		r = false
	} else {
		for i := 3; i <= Sqrt(nb); i += 2 {
			if nb%i == 0 {
				r = true
			}
		}
	}
	return r
}

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	r := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			r = i
		}
	}
	return r
}
