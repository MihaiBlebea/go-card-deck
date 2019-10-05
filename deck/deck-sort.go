package deck

import (
	"deckofcards/card"
	"sort"
)

type byRank []card.Card

func (r byRank) Len() int {
	return len(r)
}

func (r byRank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRank) Less(i, j int) bool {
	return r[i].GetRank() < r[j].GetRank()
}

type bySuit []card.Card

func (r bySuit) Len() int {
	return len(r)
}

func (r bySuit) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r bySuit) Less(i, j int) bool {
	return r[i].GetSuit() < r[j].GetSuit()
}

// SortByRank sorts the deck by card rank. Returns a deck
func (d *Deck) SortByRank() *Deck {
	sort.Sort(byRank(d.GetCards()))
	return d
}

// SortBySuit sorts the deck by card suit. Returns a deck
func (d *Deck) SortBySuit() *Deck {
	sort.Sort(bySuit(d.GetCards()))
	return d
}

// SortByColor sorts the deck by color (or suit, is the same). Returns a deck
func (d *Deck) SortByColor() *Deck {
	sort.Sort(bySuit(d.GetCards()))
	return d
}

// Shuffle returns a deck which has all his cards randomly sorted. Returns a deck
func (d *Deck) Shuffle() *Deck {
	var shuffled []card.Card
	for len(d.Cards) > 0 {
		index := d.PickIndex()
		shuffled = append(shuffled, d.Cards[index])
		d.remove(index)
	}
	d.Cards = shuffled
	return d
}
