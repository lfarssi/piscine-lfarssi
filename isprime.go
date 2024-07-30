package piscine

func IsPrime(nb int) bool {
	var r bool
	if nb < 2 {
		r = false
	} else if nb == 1 {
		return false
	} else if nb == 2 {
		r = true
	} else if nb%2 != 0 {
		r = true
	}
	return r
}
