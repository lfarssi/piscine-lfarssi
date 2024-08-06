package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// No arguments: read from standard input
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Read from files specified as command-line arguments
		for _, fileName := range os.Args[1:] {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("ERROR: open %s: %v\n", fileName, err)
				continue
			}
			if _, err := io.Copy(os.Stdout, file); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
			file.Close()
		}
	}
}
