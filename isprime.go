package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= Sqrtt(nb); i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func Sqrtt(nb int) int {
	if nb <= 0 {
		return 0
	}
	x := nb
	y := (x + nb/x) / 2
	for y < x {
		x = y
		y = (x + nb/x) / 2
	}
	return x
}
