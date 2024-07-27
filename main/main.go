package main

import "fmt"

func main() {
	QuadA(1,3)
}

func QuadA(x,y int) {
	if x > 0 && y > 0 {	
	c := "-" 
	d := " "
	fmt.Printf("°")
	for  range x {
		fmt.Printf("%v",c)
		}
	fmt.Printf("°\n")

		fmt.Printf("|")
		for range x {
			fmt.Printf("%v",d)
		}
		fmt.Printf("|\n")


	fmt.Printf("°")
	for  range x {
		fmt.Printf("%v",c)
		}
	fmt.Printf("°\n")
	}
}
