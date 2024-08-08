package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	var word []rune
	i := 0
	n := len(str)

	for i < n {
		if str[i] != ' ' { // Skip spaces
			word = append(word, rune(str[i]))
		}

		if len(word) == 5 { // When we have 5 characters
			result += string(word) + " "
			word = nil // Reset word for the next segment
			i++        // Move past the character we just processed
		}

		i++
	}

	// Remove trailing space if it exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	// Add newline character at the end of the result
	result += "\n"

	return result
}
