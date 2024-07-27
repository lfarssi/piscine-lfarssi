package piscine

import "fmt"

func QuadA(x,y int) {
	if x > 0 && y > 0 {
		for range x {
			fmt.println("°---°")
			for range y {
				fmt.println("|   |")
			}
		}
	}
}