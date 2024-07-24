package piscine

import "github.com/01-edu/z01"

func IsdPositive(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main2() {
	IsdPositive(-10)
	IsdPositive(-2)
	IsdPositive(2)
}
