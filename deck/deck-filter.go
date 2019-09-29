package deck

import (
	"fmt"
	"go-card-deck/card"
	"sort"
)

// FilterByRank filters the deck by card rank. Returns a pointer to the deck
func (d *Deck) FilterByRank() *Deck {
	sort.Ints([]int{1,2,3,4})
	return d
}

// FilterByRank filters the deck by card suit. Returns a pointer to the deck
func (d *Deck) FilterBySuit() *Deck {
	fmt.Println(d.cards)
	return d
}

// FilterByRank filters the deck by card color. Returns a pointer to the deck
func (d *Deck) FilterByColor() *Deck {
	fmt.Println(d.cards)
	return d
}

// Shuffle returns a deck which has all his cards randomly sorted
func (d *Deck) Shuffle() *Deck {
	var shuffled []card.CardType
	for len(d.cards) > 0 {
		shuffled = append(shuffled, d.Draw())
	}
	d.cards = shuffled
	return d
}
