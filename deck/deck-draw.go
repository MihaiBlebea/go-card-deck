package deck

import (
	"go-card-deck/card"
	"log"
	"math/rand"
	"time"
)

// PickIndex returns a random valid index from the deck of cards
func (d *Deck) PickIndex() int {
	length := len(d.cards)
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(length)
}

// Pick returns a random valid card from the deck of cards
func (d *Deck) Pick() card.CardType {
	return d.cards[d.PickIndex()]
}

// Draw returns a valid card from the deck. In the same time, the returned card is removed from the deck
func (d *Deck) Draw() card.CardType {
	index := d.PickIndex()
	card := d.cards[index]
	d.remove(card)

	return card
}

// DrawHand returns a number of cards from the deck picked randomly. The returned cards are removed from the deck
func (d *Deck) DrawHand(count int) []card.CardType {
	if count < 1 {
		log.Print("A hand of cards cannot have less then 1 card")
	}
	var hand []card.CardType
	for i := 0; i < count; i++ {
		hand = append(hand, d.Draw())
	}
	return hand
}
