package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1]
	runes := []rune(params)
	for i := range runes {
		z01.PrintRune(runes[i])
		z01.PrintRune('\n')
	}
}
