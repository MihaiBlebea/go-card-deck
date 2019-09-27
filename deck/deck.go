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

func (d Deck) PickIndex() int {
	length := len(d.cards)
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(length)
}

func (d Deck) Pick() card.Card {
	return d.cards[d.PickIndex()]
}

func (d *Deck) Draw() card.Card {
	index := d.PickIndex()
	card := d.cards[index]
	d.remove(card)

	return card
}

func (d *Deck) DrawHand(count int) []card.Card {
	if count < 1 {
		log.Print("A hand of cards cannot have less then 1 card")
	}
	var hand []card.Card
	for i := 0; i < count; i++ {
		hand = append(hand, d.Draw())
	}
	return hand
}

func (d Deck) Shuffle() Deck {
	var shuffled []card.Card
	for len(d.cards) > 0 {
		shuffled = append(shuffled, d.Draw())
	}
	d.cards = shuffled
	return d
}

func (d *Deck) remove(card card.Card) *Deck {
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

func New() Deck {
	ranks := card.GetRankTypes()
	suits := card.GetSuitTypes()

	var cards []card.Card
	for i := 0; i < len(ranks); i++ {
		for k := 0; k < len(suits); k++ {
			card := card.New(suits[k], ranks[i])
			cards = append(cards, card)
		}
	}
	return Deck{cards}
}
