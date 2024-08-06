package main

import (
	"os"
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	r := true
	if nbr%2 == 1 {
		return true
	} else if nbr % 2 == 0{
		return false
	}
	return r
}

func main() {
	if isEven(len(os.Args)) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
