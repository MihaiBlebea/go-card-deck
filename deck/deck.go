package deck

import (
	"go-card-deck/card"
)

type Deck struct {
	cards []card.CardType
}

// remove removes a card from the deck and returns a pointer to the deck
func (d *Deck) remove(card card.CardType) *Deck {
	var index int
	for i := 0; i < len(d.cards); i++ {
		if d.cards[i].EqualSuit(card) && d.cards[i].EqualRank(card) {
			index = i
		}
	}

	cards := d.cards
	cards[len(cards)-1], cards[index] = cards[index], cards[len(cards)-1]
	d.cards = cards[:len(cards)-1]

	return d
}

// hasCard checks if the deck contains a card. It returns the index of the card and true if card is found, 0 and false if card is not found
func (d *Deck) hasCard(card card.CardType) (int, bool) {
	for index, deckCard := range d.cards {
		if deckCard.EqualRank(card) && deckCard.EqualSuit(card) {
			return index, true
		}
	}
	return 0, false
}

// Add a card to the deck,
func (d Deck) Add(card card.CardType) error {
	_, found := d.hasCard(card)
	if found == true {
		return newDeckError("Card already exists in the deck")
	}
	d.cards = append(d.cards, card)
	return nil
}

// DeckOfCards interface
type DeckOfCards interface {
	Add(card card.CardType) error
	Pick() card.CardType
	PickIndex() int
	Draw() card.CardType
	DrawHand(count int) []card.CardType
	FilterByRank() DeckOfCards
	FilterBySuit() DeckOfCards
	FilterByColor() DeckOfCards
	Shuffle() DeckOfCards
}
