package deck

type CardJson map[string]interface{}

// LastCard returns the last card in the deck
func (d *Deck) ShowLastCard() CardJson {
	return d.cards[len(d.cards) - 1].Show()
}

// ShowCards returns an array of cards in json format
func (d *Deck) ShowCards() []CardJson {
	var allCards []CardJson
	for _, card := range d.cards {
		allCards = append(allCards, card.Show())
	}
	return allCards
}