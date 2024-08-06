package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(message string) {
	for _, c := range message {
		z01.PrintRune(c)
	}
}

func main() {
	if len(os.Args) < 2 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printError("Error reading from stdin: " + err.Error() + "\n")
			os.Exit(1)
		}
		return
	}

	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			printError("ERROR: " + err.Error() + "\n")
			os.Exit(1)
			continue
		}
		defer file.Close()

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			printError("Error reading from file " + fileName + ": " + err.Error() + "\n")
			os.Exit(1)
		}
	}
}
