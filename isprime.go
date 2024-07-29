package piscine

func IsPrime(nb int) bool {
	var r bool
	if nb < 0 {
		r = false
	}
	if nb%2 == 0 {
		r = false
	} else if nb%2 != 0 {
		if nb == 1 {
			r = false
		}
		r = true
	}
	return r
}
