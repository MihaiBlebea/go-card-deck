package deck

import (
	"deckofcards/card"
)

type Deck struct {
	Cards []card.Card
}

func (d *Deck) GetCards() []card.Card {
	return d.Cards
}
