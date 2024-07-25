package piscine

func BasicAtoi(s string) int {
	res := 0
	for _, char := range s {
		res = res*10 + int(char-'0')
	}
	return res
}
