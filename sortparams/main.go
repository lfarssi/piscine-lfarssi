package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, val := range os.Args[1:] {
		runes := []rune(val)
		n := len(runes)
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if runes[j] > runes[j+1] {
					runes[j], runes[j+1] = runes[j+1], runes[j]
				}
			}
		}
		for _, j := range runes {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
