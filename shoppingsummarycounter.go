package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var ch string
	list := make(map[string]int)
	for _, i := range str {
		if i == 32 {
			list[ch] += 1
			ch = ""
		} else if i != 32 {
			ch += string(rune(i))
		}
	}
	list[ch] += 1
	return list
}
