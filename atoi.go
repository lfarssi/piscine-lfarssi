package piscine

func Atoi(s string) int {
	res := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		if char == '+' || char == '-' {
			return 0
		}
		res = res*10 + int(char-'0')
	}
	return res
}
