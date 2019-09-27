package card

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
