package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("Too many arguments\n")
		return
	} else if len(os.Args) < 2 {
		fmt.Printf("File name missing\n")
		return
	}
	name := os.Args[1]
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Err reading files %v\n ", err)
		return
	}
	fmt.Printf("%s", file)
}
