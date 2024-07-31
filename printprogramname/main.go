package main

import (
	"github.com/01-edu/z01"

	"os"
	
	"path/filepath"
)

func main() {
	name := os.Args[0]
	n := filepath.Base(name)
	for _, val := range n {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
