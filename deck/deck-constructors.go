package deck

import "go-card-deck/card"

// New is a deck constructor. It returns a new sorted (not shuffled) deck of cards
func New() Deck {
	ranks := card.GetRankTypes()
	suits := card.GetSuitTypes()

	var cards []card.CardType
	for i := 0; i < len(ranks); i++ {
		for k := 0; k < len(suits); k++ {
			card := card.New(suits[k], ranks[i])
			cards = append(cards, card)
		}
	}
	return Deck{cards}
}

// NewFromCards is a deck constructor. It returns a new deck filled with the supplied cards
func NewFromCards(cards []card.CardType) Deck {
	return Deck{cards}
}
