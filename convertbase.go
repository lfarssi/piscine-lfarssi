package piscine

import (
	"strings"
)

// ConvertBase converts a number from baseFrom to baseTo
func ConvertBase(nbr, baseFrom, baseTo string) string {
	if nbr == "" {
		return ""
	}

	// Convert from baseFrom to decimal
	decimalValue := 0
	baseFromLen := len(baseFrom)
	for i := 0; i < len(nbr); i++ {
		char := string(nbr[i])
		position := strings.Index(baseFrom, char)
		if position == -1 {
			return "" // Invalid character found
		}
		decimalValue = decimalValue*baseFromLen + position
	}

	// Handle the case where baseTo is decimal
	if baseTo == "0123456789" {
		return fmt.Sprintf("%d", decimalValue)
	}

	// Convert from decimal to baseTo
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	var result []string
	baseToLen := len(baseTo)
	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = append([]string{string(baseTo[remainder])}, result...)
		decimalValue /= baseToLen
	}

	return strings.Join(result, "")
}
