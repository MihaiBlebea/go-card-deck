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

func GetRankTypes() [13]Rank {
	return [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

func GetSuitTypes() [4]Suit {
	return [4]Suit{Heart, Diamond, Club, Spade}
}

// Card model is the main actor of our package
type Card struct {
	rank Rank
	suit Suit
}

func (c Card) GetSuit() (string, string) {
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

func (c Card) GetValue() int {
	return int(c.rank)
}

func (c *Card) SetValue(value int) {
	c.rank = Rank(value)
}

func (c Card) equalColor(card Card) bool {
	suit1, _ := c.GetSuit()
	suit2, _ := card.GetSuit()
	if suit1 == suit2 {
		return true
	}
	return false
}

func (c Card) EqualSuit(card Card) bool {
	if c.suit == card.suit {
		return true
	}
	return false
}

func (c Card) EqualRank(card Card) bool {
	if c.rank == card.rank {
		return true
	}
	return false
}

func (c Card) IsBiggerThen(card Card) bool {
	if c.rank > card.rank {
		return true
	}
	return false
}

func (c Card) IsSmallerThen(card Card) bool {
	if c.rank < card.rank {
		return true
	}
	return false
}

func (c Card) Show() (string, string, int) {
	suit, color := c.GetSuit()
	return suit, color, c.GetValue()
}

func New(suit Suit, rank Rank) Card {
	return Card{
		rank,
		suit,
	}
}
