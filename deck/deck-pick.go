package deck

import "deckofcards/card"

// PickIndex returns a random valid index from the deck of cards. Card is not substracted from the deck
func (d *Deck) PickIndex() int {
	return random(len(d.Cards))
}

// Pick returns a random valid card from the deck of cards. Card is not substracted from the deck
func (d *Deck) Pick() *card.Card {
	return &d.Cards[random(len(d.Cards))]
}
