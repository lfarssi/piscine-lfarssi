package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	combinations := generateCombinations(n, 0, "")

	for i, comb := range combinations {
		for _, ch := range comb {
			z01.PrintRune(ch)
		}
		if i < len(combinations)-1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func generateCombinations(n int, start int, prefix string) []string {
	if n == 0 {
		return []string{prefix}
	}

	combinations := []string{}
	for i := start; i <= 9; i++ {
		newPrefix := prefix + string('0'+i)
		combinations = append(combinations, generateCombinations(n-1, i+1, newPrefix)...)
	}
	return combinations
}
