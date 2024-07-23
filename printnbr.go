package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			printDigits(-(n / 10))
			z01.PrintRune('8')
			return
		}
		n = -n
	}

	printDigits(n)
}

func printDigits(n int) {
	if n == 0 {
		return
	}

	printDigits(n / 10)
	z01.PrintRune(rune('0' + n%10))
}
