package main

import "github.com/01-edu/z01"

var str string = "x = 42, y = 21"

func main() {
	for _, i := range str {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
