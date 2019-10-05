package deck

import (
	"deckofcards/card"
	"math/rand"
	"time"
)

func random(max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(max)
}

// remove removes a card from the deck and returns a pointer to the deck
func (d *Deck) remove(index int) *Deck {
	cards := d.GetCards()
	cards[len(cards)-1], cards[index] = cards[index], cards[len(cards)-1]
	d.Cards = cards[:len(cards)-1]
	return d
}

func (d *Deck) hasCard(card card.Card) (int, bool) {
	for index, deckCard := range d.GetCards() {
		if deckCard.EqualRank(card) && deckCard.EqualSuit(card) {
			return index, true
		}
	}
	return 0, false
}
