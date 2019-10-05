package deck

import (
	"deckofcards/card"
)

// DrawHand returns a random card from the deck of cards. Card will be substracted from the deck
func (d *Deck) Draw() card.Card {
	index := d.PickIndex()
	defer d.remove(index)
	return d.Cards[index]
}

// DrawHand returns a random specified number of cards from the deck of cards. Cards will be substracted from the deck
func (d *Deck) DrawHand(count int) []card.Card {
	var hand []card.Card
	for i := 0; i < count; i++ {
		crd := d.Draw()
		hand = append(hand, crd)
	}
	return hand
}

// DrawFirstCard returns the last card in the deck of cards. Card will be substracted from the deck
func (d *Deck) DrawLastCard() card.Card {
	index := len(d.Cards) - 1
	defer d.remove(index)
	return d.GetCards()[index]
}

// DrawFirstCard returns the first card in the deck of cards. Card will be substracted from the deck
func (d *Deck) DrawFirstCard() card.Card {
	index := 0
	defer d.remove(index)
	return d.GetCards()[index]
}

// DrawFirstCard returns a random index from the deck of cards. Card will be substracted from the deck
func (d *Deck) DrawIndex(index int) card.Card {
	defer d.remove(index)
	return d.Cards[index]
}

// DrawCard returns a specified card from the deck of cards. Card will be substracted from the deck
func (d *Deck) DrawCard(crd card.Card) card.Card {
	var removeIndex int
	for index, currentCard := range d.GetCards() {
		if crd.EqualRank(currentCard) && crd.EqualSuit(currentCard) {
			removeIndex = index
		}
	}
	defer d.remove(removeIndex)
	return d.Cards[removeIndex]
}
