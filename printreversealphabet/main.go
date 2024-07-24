package main

import "github.com/01-edu/z01"

func main() {
	var r rune
	for r = 'z'; r >= 'a'; r-- {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
