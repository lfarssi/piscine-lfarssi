package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var s string
	var r []rune
	for _, i := range str {
		if i != 32 {
			r = append(r, i)
		}
		if len(r) == 5 {
			s += string(r) + " "
			r = nil
			if len(str) > 0 {
				str = str[1:]
			}
		}
	}
	if len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	s += "\n"
	return s
}
