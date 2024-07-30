package piscine

func IsPrime(nb int) bool {
	var r bool
	if nb <= 1 {
		return false
	} else if nb == 2 {
		r = true
	} /*else if nb%2 == 0 {
		r = false
	} else {
		for i := 3; i <= Sqrt(nb); i += 2 {
			if nb%i == 0 {
				r = true
			}
		}
	}*/
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			r = false
		}
	}
	return r
}
