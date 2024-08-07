package piscine

func JumpOver(str string) string {
	var ch string
	for i := 2; i < len(str); i += 3 {
		ch += string(str[i])
	}
	ch += "\n"
	return ch
}
