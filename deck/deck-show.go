package deck

import "deckofcards/card"

// ShowLastCard returns the last card in the deck. Card will not be substracted from the deck
func (d *Deck) ShowLastCard() *card.Card {
	return &d.GetCards()[len(d.GetCards())-1]
}

// ShowFirstCard returns the first card in the deck. Card will not be substracted from the deck
func (d *Deck) ShowFirstCard() *card.Card {
	return &d.GetCards()[0]
}

// ShowCards returns a list of cards from the deck. Cards will not be substracted from the deck
func (d *Deck) ShowCards() *[]card.Card {
	var allCards []card.Card
	for _, card := range d.GetCards() {
		allCards = append(allCards, card)
	}
	return &allCards
}
