package deck

import "deckofcards/card"

// Filter accepts a function that returns bool as a param and returns a slice of cards filtered by that function
func (d *Deck) Filter(f func(indx int, crd card.Card) bool) []card.Card {
	var crds []card.Card
	for index, currentCard := range d.GetCards() {
		if f(index, currentCard) == true {
			crds = append(crds, currentCard)
		}
	}
	return crds
}
