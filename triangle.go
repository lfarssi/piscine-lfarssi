package piscine 

import "fmt"

func Triangle(a,b int) {

	for i := 0 ; i< a ; i++ {
		for j := i ; j < b ; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}