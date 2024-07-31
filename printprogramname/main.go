package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programme := os.Args[0]
	for i := 2; i <= len(programme)-1; i++ {
		z01.PrintRune(rune(programme[i]))
	}
	z01.PrintRune('\n')
}
