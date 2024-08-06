package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		// No arguments: read from standard input
		buffer := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buffer)
			if n > 0 {
				printBytes(buffer[:n])
			}
			if err != nil {
				if err.Error() != "EOF" {
					printError("ERROR: %v\n", err)
				}
				break
			}
		}
	} else {
		// Read from files specified as command-line arguments
		for _, fileName := range os.Args[1:] {
			file, err := os.Open(fileName)
			if err != nil {
				printError("ERROR: open %s: %v\n", fileName, err)
				continue
			}
			buffer := make([]byte, 1024)
			for {
				n, err := file.Read(buffer)
				if n > 0 {
					printBytes(buffer[:n])
				}
				if err != nil {
					if err.Error() != "EOF" {
						printError("ERROR: %v\n", err)
					}
					break
				}
			}
			file.Close()
		}
	}
}

func printBytes(data []byte) {
	for _, b := range data {
		z01.PrintRune(rune(b))
	}
}

func printError(format string, args ...interface{}) {
	for _, r := range format {
		z01.PrintRune(r)
	}
	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			printBytes([]byte(v.Error()))
		default:
			// Handle other types if needed
		}
	}
}
