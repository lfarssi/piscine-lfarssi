package piscine

func IsPrime(nb int) bool {
	var r bool

	if nb <= 1 {
		return false
	} else if nb == 2 {
		r = true
	} else {
		for i := 3 ; i < nb ; i++ {
			if nb%i == 0 {
				r = false
			} else {
				r = true
			}
		}
	}
	return r
}
