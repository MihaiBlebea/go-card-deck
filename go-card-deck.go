package main

import (
	"fmt"
	"go-card-deck/deck"
)

func main() {
	deck := deck.New()
	deck.Shuffle()

	hand1 := deck.DrawHand(5)
	hand2 := deck.DrawHand(5)
	onTable := deck.Draw()

	fmt.Println(hand1[0].Show())
	fmt.Println(hand1)
	fmt.Println(hand2)
	fmt.Println(onTable)
}
