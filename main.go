package main

import (
	"fmt"
	"log"
	"go-card-deck/card"
	"go-card-deck/deck"
)

func main() {
	pack := deck.New()
	pack.Shuffle()

	err := pack.Add(card.New(1, 4))
	if err != nil {
		log.Panic(err)
	}
	printHand("Mihai", pack.DrawHand(10) )

}

func printHand(name string, hand []card.CardType) {

	fmt.Println("Show hand for " + name)
	for _, value := range hand {
		fmt.Println(value.Show())
	}
	fmt.Println("---------------")
}
