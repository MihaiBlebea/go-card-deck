package deck

import "deckofcards/card"

// Add a card to the deck,
func (d *Deck) Add(card card.Card) error {
	_, found := d.hasCard(card)
	if found == true {
		return newDeckError("Card already exists in the deck")
	}
	d.Cards = append(d.Cards, card)
	return nil
}
