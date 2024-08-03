package piscine

// Split splits the string s by the separator sep and returns a slice of strings.
func Split(s, sep string) []string {
	var result []string

	if len(sep) == 0 {
		return []string{s} // If separator is empty, return the original string in a slice.
	}

	start := 0
	for i := 0; i < len(s); i++ {
		// Check if the current position starts the separator
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			// Add the substring from start to the current index to the result
			result = append(result, s[start:i])
			// Move the start position past the separator
			start = i + len(sep)
			// Skip over the separator
			i += len(sep) - 1
		}
	}

	// Add the last segment after the last separator
	result = append(result, s[start:])

	return result
}
