package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result string
	var word []rune
	i := 0
	for i < len(str) {
		if str[i] != ' ' {
			word = append(word, rune(str[i]))
		}

		if len(word) == 5 {
			result += string(word) + " "
			word = nil
			i++
		}

		i++
	}
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result += "\n"
	return result
}
