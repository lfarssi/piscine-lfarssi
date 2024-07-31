package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	executablePath := os.Args[0]

	var executableName string
	for i := len(executablePath) - 1; i >= 0; i-- {
		if executablePath[i] == '/' {
			executableName = executablePath[i+1:]
			break
		}
	}
	if executableName == "" {
		executableName = executablePath
	}

	runes := []rune(executableName)
	for i := range runes {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
