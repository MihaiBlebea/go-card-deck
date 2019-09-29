package main

import (
	"fmt"
	// "go-card-deck/card"
	"go-card-deck/deck"
)

func main() {
	pack := deck.New()
	pack.Shuffle()

	player1 := deck.NewFromCards(pack.DrawHand(5))
	player2 := deck.NewFromCards(pack.DrawHand(5))

	table := deck.NewFromCards(pack.DrawHand(1))

	fmt.Println(player1.ShowCards())
	fmt.Println(player2.ShowCards())

	fmt.Println(table.ShowLastCard())


	table.Add(player1.DrawLastCard())

	fmt.Println(player1.ShowCards())
}	
