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

// GetSuit returns the suit and color of the card in literal, string form
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

// GetValue returns an integer value of the card
func (c Card) GetValue() int {
	return int(c.rank)
}

// SetValue sets an integer as the value of the card
func (c Card) SetValue(value int) {
	c.rank = Rank(value)
}

// GetRank returns an integer value of the card
func (c Card) GetRank() int {
	return int(c.rank)
}

// Show returns the suit, color and value of the card
func (c Card) Show() map[string]interface{} {
	suit, color := c.GetSuit()
	result := make(map[string]interface{})
	result["suit"] = suit
	result["color"] = color
	result["rank"] = c.GetValue()
	return result
}

// CardType is the interface for the Card model
type CardType interface {
	GetSuit() (string, string)
	GetValue() int
	SetValue(value int)
	GetRank() int
	Show() map[string]interface{}
	EqualColor(card CardType) bool
	EqualSuit(card CardType) bool
	EqualRank(card CardType) bool
	IsBiggerThen(card CardType) bool
	IsSmallerThen(card CardType) bool
}
