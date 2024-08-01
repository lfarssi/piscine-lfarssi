package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, i := range os.Args[1:] {
		runes := []rune(i)
		for j := range runes {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
