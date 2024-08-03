package piscine

// ConvertBase converts a number from baseFrom to baseTo.
func ConvertBase(nbr, baseFrom, baseTo string) string {
	if len(nbr) == 0 || len(baseFrom) < 2 || len(baseTo) < 2 {
		return ""
	}

	// Helper function to get the value of a digit in the base
	getDigitValue := func(digit byte, base string) int {
		for i := 0; i < len(base); i++ {
			if base[i] == digit {
				return i
			}
		}
		return -1 // Not found
	}

	// Convert from baseFrom to decimal
	decimalValue := 0
	baseFromLen := len(baseFrom)
	for i := 0; i < len(nbr); i++ {
		digitValue := getDigitValue(nbr[i], baseFrom)
		if digitValue == -1 {
			return ""
		}
		decimalValue = decimalValue*baseFromLen + digitValue
	}

	// Special case for zero
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	// Convert from decimal to baseTo
	var result []byte
	baseToLen := len(baseTo)
	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = append(result, baseTo[remainder])
		decimalValue /= baseToLen
	}

	// Reverse the result slice manually
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Convert result slice to string
	return string(result)
}
