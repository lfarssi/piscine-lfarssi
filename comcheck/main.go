package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	for _, i := range arg {
		if i == "01" || i == "galaxy" || i == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
