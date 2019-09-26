package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Suit int

// All the suit types with a value of integer starting from 1
const (
	_ Suit = iota
	Heart
	Spade
	Club
	Diamond
)

type Rank int

// All the rank types with a value of integer starting from 1
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card model is the main actor of our package
type Card struct {
	rank Rank
	suit Suit
}

func (c Card) getSuit() (string, string) {
	var color string
	var suit string
	switch c.suit {
	case Heart:
		color = "red"
		suit = "heart"
	case Diamond:
		color = "black"
		suit = "diamond"
	case Club:
		color = "black"
		suit = "club"
	case Spade:
		color = "red"
		suit = "spade"
	}
	return suit, color
}

func (c Card) getValue() int {
	return int(c.rank)
}

func (c *Card) setValue(value int) {
	c.rank = Rank(value)
}

func (c Card) equalColor(card Card) bool {
	suit1, _ := c.getSuit()
	suit2, _ := card.getSuit()
	if suit1 == suit2 {
		return true
	}
	return false
}

func (c Card) equalSuit(card Card) bool {
	if c.suit == card.suit {
		return true
	}
	return false
}

func (c Card) equalRank(card Card) bool {
	if c.rank == card.rank {
		return true
	}
	return false
}

func (c Card) isBiggerThen(card Card) bool {
	if c.rank > card.rank {
		return true
	}
	return false
}

func (c Card) isSmallerThen(card Card) bool {
	if c.rank < card.rank {
		return true
	}
	return false
}

func (c Card) show() (string, string, int) {
	suit, color := c.getSuit()
	return suit, color, c.getValue()
}

type Deck struct {
	cards []Card
}

func (d Deck) pickIndex() int {
	length := len(d.cards)
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(length)
}

func (d Deck) pick() Card {
	return d.cards[d.pickIndex()]
}

func (d *Deck) draw() Card {
	index := d.pickIndex()
	card := d.cards[index]
	d.remove(card)

	return card
}

func (d *Deck) drawHand(count int) []Card {
	if count < 1 {
		log.Print("A hand of cards cannot have less then 1 card")
	}
	var hand []Card
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

func (d *Deck) remove(card Card) *Deck {
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

func newCard(suit Suit, rank Rank) Card {
	return Card{
		rank,
		suit,
	}
}

func newDeck() Deck {
	ranks := [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	suits := [4]Suit{Heart, Diamond, Club, Spade}

	var cards []Card
	for i := 0; i < len(ranks); i++ {
		for k := 0; k < len(suits); k++ {
			card := Card{ranks[i], suits[k]}
			cards = append(cards, card)
		}
	}
	return Deck{cards}
}

func main() {
	deck := newDeck()
	deck.shuffle()

	hand1 := deck.drawHand(5)
	hand2 := deck.drawHand(5)
	onTable := deck.draw()

	fmt.Println(hand1[0].show())
	fmt.Println(hand1)
	fmt.Println(hand2)
	fmt.Println(onTable)
}
