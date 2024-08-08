package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: non divisible"
	} else if n%2 == 0 && n%3 == 0 {
		return "rock and roll"
	} else if n%3 == 0 {
		return "roll"
	} else if n%2 == 0 {
		return "rock"
	}
	return "error: non divisible"
}
