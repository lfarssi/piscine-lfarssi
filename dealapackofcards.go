package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	var player int = 4
	var Card int = 3
	for i := 0; i < player; i++ {
		start := i * Card
		end := start + Card
		card := deck[start:end]
		fmt.Printf("player : %d", i+1)
		for j, v := range card {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d",v)
		}
		fmt.Printf("\n")
	}
}
