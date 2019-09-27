package deck

import (
	"go-card-deck/card"
	"log"
	"math/rand"
	"time"
)

type Deck struct {
	cards []card.Card
}

func (d Deck) pickIndex() int {
	length := len(d.cards)
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(length)
}

func (d Deck) pick() card.Card {
	return d.cards[d.pickIndex()]
}

func (d *Deck) draw() card.Card {
	index := d.pickIndex()
	card := d.cards[index]
	d.remove(card)

	return card
}

func (d *Deck) drawHand(count int) []card.Card {
	if count < 1 {
		log.Print("A hand of cards cannot have less then 1 card")
	}
	var hand []card.Card
	for i := 0; i < count; i++ {
		hand = append(hand, d.draw())
	}
	return hand
}

func (d Deck) shuffle() Deck {
	var shuffled []Card
	for len(d.cards) > 0 {
		shuffled = append(shuffled, d.draw())
	}
	d.cards = shuffled
	return d
}

func (d *Deck) remove(card card.Card) *Deck {
	var index int
	for i := 0; i < len(d.cards); i++ {
		if d.cards[i].equalSuit(card) && d.cards[i].equalRank(card) {
			index = i
		}
	}

	cards := d.cards

	cards[len(cards)-1], cards[index] = cards[index], cards[len(cards)-1]
	d.cards = cards[:len(cards)-1]

	return d
}

func newCard(suit card.Suit, rank card.Rank) card.Card {
	return card.Card{
		rank,
		suit,
	}
}

func newDeck() Deck {
	ranks := [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	suits := [4]Suit{Heart, Diamond, Club, Spade}

	var cards []card.Card
	for i := 0; i < len(ranks); i++ {
		for k := 0; k < len(suits); k++ {
			card := card.Card{ranks[i], suits[k]}
			cards = append(cards, card)
		}
	}
	return Deck{cards}
}
