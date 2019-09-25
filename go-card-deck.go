package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var colors = [2]string{"red", "black"}
var signs = [4]string{"hearts", "spades", "diamonds", "clubs"}

type Card struct {
	number Number
	color  Color
	sign   Sign
}

func (c Card) equalColor(card Card) bool {
	if c.color.value == card.color.value {
		return true
	}
	return false
}

func (c Card) isBiggerThen(card Card) bool {
	if c.number.value > card.number.value {
		return true
	}
	return false
}

func (c Card) isSmallerThen(card Card) bool {
	if c.number.value < card.number.value {
		return true
	}
	return false
}

func (c Card) isEqualWith(card Card) bool {
	if c.number.value == card.number.value {
		return true
	}
	return false
}

type Number struct {
	value int
}

type Color struct {
	value string
}

type Sign struct {
	value string
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
	cards := d.cards
	cards[len(cards)-1], cards[index] = cards[index], cards[len(cards)-1]
	d.cards = cards[:len(cards)-1]

	return cards[len(cards)-1]
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

func newNumber(value int) Number {
	valid := false
	for i := 1; i <= 14; i++ {
		if value == i {
			valid = true
		}
	}

	if valid != true {
		log.Print("Number value is not valid")
	}
	return Number{value}
}

func newColor(value string) Color {
	valid := false
	for i := 0; i < len(colors); i++ {
		if colors[i] == value {
			valid = true
		}
	}
	if valid != true {
		log.Panic("Color is not valid")
	}
	return Color{value}
}

func newSign(value string) Sign {
	valid := false
	for i := 0; i < 4; i++ {
		if signs[i] == value {
			valid = true
		}
	}

	if valid != true {
		log.Panic("Sign is not valid")
	}
	return Sign{value}
}

func newCard(nr int, colr string, sgn string) Card {
	color := newColor(colr)
	sign := newSign(sgn)
	number := newNumber(nr)

	if color.value == "red" {
		if sign.value != "hearts" && sign.value != "diamonds" {
			log.Panic("Color and sign are not valid")
		}
	}

	if color.value == "black" {
		if sign.value != "spades" && sign.value != "clubs" {
			log.Panic("Color and sign are not valid")
		}
	}
	return Card{
		number,
		color,
		sign,
	}
}

func newDeck() Deck {
	var cards []Card
	for i := 1; i <= 14; i++ {
		cards = append(cards, newCard(i, "red", "hearts"))
		cards = append(cards, newCard(i, "red", "diamonds"))
		cards = append(cards, newCard(i, "black", "spades"))
		cards = append(cards, newCard(i, "black", "clubs"))
	}
	return Deck{cards}
}

func main() {
	deck := newDeck()
	deck.shuffle()

	card1 := deck.draw()
	card2 := deck.draw()

	result := card1.isBiggerThen(card2)
	fmt.Println(card1)
	fmt.Println(card2)
	fmt.Println(result)
}
