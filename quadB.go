package piscine

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("/")
		for i := 1; i < x-1; i++ {
			fmt.Print("*")
		}
		if x > 1 {
			fmt.Print("\\")
		}
		fmt.Println()
		for j := 1; j < y-1; j++ {
			fmt.Print("*")
			for i := 1; i < x-1; i++ {
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Print("*")
			}
			fmt.Println()
		}

		if y > 1 {
			fmt.Print("\\")
			for i := 1; i < x-1; i++ {
				fmt.Print("*")
			}
			if x > 1 {
				fmt.Print("/")
			}
			fmt.Println()
		}
	}
}
