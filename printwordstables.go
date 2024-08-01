package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var b string
	for _, i := range a {
		b += i
		for _, j := range b {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
		b = ""
	}
}
