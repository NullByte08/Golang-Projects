package main

import "fmt"

func main() {
	cards := deck{"Ace", "Clubs"}

	cards = append(cards, "Hearts")

	for _, card := range cards {
		fmt.Println(card)
	}
}
